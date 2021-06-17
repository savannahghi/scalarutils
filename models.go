package go_utils

import (
	"fmt"
	"time"
)

// Upload represents a file uploaded to cloud storage
type Upload struct {
	ID          string    `json:"id" firestore:"id"`
	URL         string    `json:"url" firestore:"url"`
	Size        int       `json:"size" firestore:"size"`
	Hash        string    `json:"hash" firestore:"hash"`
	Creation    time.Time `json:"creation" firestore:"creation"`
	Title       string    `json:"title" firestore:"title"`
	ContentType string    `json:"contentType" firestore:"contentType"`
	Language    string    `json:"language" firestore:"language"`
	Base64data  string    `json:"base64data" firestore:"base64data"`
}

// IsEntity marks upload as an apollo federation entity
func (u Upload) IsEntity() {}

// IsNode marks upload as a relay node
func (u Upload) IsNode() {}

// SetID marks upload as a relay node
func (u Upload) SetID(id string) {
	u.ID = id
}

// GetID marks upload as a relay node
func (u Upload) GetID() ID {
	return IDValue(u.ID)
}

// UploadInput is used to send data for new uploads
type UploadInput struct {
	Title       string `json:"title"`
	ContentType string `json:"contentType"`
	Language    string `json:"language"`
	Base64data  string `json:"base64data"`
	Filename    string `json:"filename"`
}

// CustomError represents a custom error struct
// Reference https://blog.golang.org/error-handling-and-go
type CustomError struct {
	Err     error  `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
