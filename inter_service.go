package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GetJWTKey returns a byte slice of the JWT secret key
func GetJWTKey() []byte {
	key := MustGetEnvVar(JWTSecretKey)
	return []byte(key)
}

// Claims a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide public claims
// Provides way for adding private claims
type Claims struct {
	jwt.StandardClaims
}

// GetServiceEnvirionmentSuffix get the env suffix where the app is running
// e.g testing, staging, prod, local
func GetServiceEnvirionmentSuffix() string {
	environment := MustGetEnvVar(ServiceEnvironmentSuffix)

	return environment
}

// Service used to keep record of a service and the REST paths it has
type Service struct {
	Name  string
	Paths map[string]string
}

// InterServiceClient defines a client for use in interservice communication
type InterServiceClient struct {
	// services offering a rest api
	Mailgun Service

	// service is the name of service initializing the client
	service     string
	environment string
	apiScheme   string
	domain      string

	httpClient  *http.Client
	accessToken string
}

// NewInterserviceClient ...
func NewInterserviceClient(service string) (*InterServiceClient, error) {

	env := GetServiceEnvirionmentSuffix()

	mailgun := Service{
		Name: "mailgun",
		Paths: map[string]string{
			"sendEmail": "communication/send_email",
		},
	}

	return &InterServiceClient{
		service:     service,
		Mailgun:     mailgun,
		environment: env,
		apiScheme:   "https",
		domain:      "healthcloud.co.ke",
		httpClient: &http.Client{
			Timeout: time.Duration(1 * time.Minute),
		},
	}, nil
}

// CreateAuthToken returns a signed JWT for use in authentication.
func (c InterServiceClient) CreateAuthToken() (string, error) {

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    c.GenerateBaseURL(c.service),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJWTKey())
	if err != nil {
		return "", fmt.Errorf("failed to create token with err: %v", err)
	}

	c.accessToken = tokenString

	return tokenString, nil
}

// GenerateBaseURL generates a URL depending on the environment
func (c InterServiceClient) GenerateBaseURL(service string) string {

	var address string

	if c.environment == "local" {

		port := MustGetEnvVar("PORT")
		address = "http://localhost:" + port
	} else {

		subdomain := fmt.Sprintf("%v-%v", service, c.environment)
		address = fmt.Sprintf("%v://%v.%v", c.apiScheme, subdomain, c.domain)
	}

	return address
}

// GenerateRequestURL generate a url with path for requested resource.
func (c InterServiceClient) GenerateRequestURL(service string, path string) string {

	address := c.GenerateBaseURL(service)

	return fmt.Sprintf("%v/%v", address, path)
}

// MakeRequest performs an inter service http request and returns a response
func (c InterServiceClient) MakeRequest(method string, url string, body interface{}) (*http.Response, error) {

	token, tknErr := c.CreateAuthToken()
	if tknErr != nil {
		return nil, tknErr
	}

	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(encoded)

	req, reqErr := http.NewRequest(method, url, payload)
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, respErr := c.httpClient.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode > 201 {
		return nil, fmt.Errorf("bad response got: %v", resp.StatusCode)
	}

	return resp, nil
}
