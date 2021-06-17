package go_utils

import "time"

/* #nosec */
const (
	DateLayout = "2006-01-02"

	DateTimeFormatLayout = "2006-01-02T15:04:05+03:00"

	// Sep is a separator, used to create "opaque" IDs
	Sep = "|"

	// GoogleProjectNumberEnvVarName is a numeric project number that
	GoogleProjectNumberEnvVarName = "GOOGLE_PROJECT_NUMBER"

	// DefaultRESTAPIPageSize is the page size to use when calling Slade REST API services if the
	// client does not specify a page size
	DefaultRESTAPIPageSize = 100

	// GoogleCloudProjectIDEnvVarName is used to determine the ID of the GCP project e.g for setting up StackDriver client
	GoogleCloudProjectIDEnvVarName = "GOOGLE_CLOUD_PROJECT"

	// DebugEnvVarName is used to determine if we should print extended tracing / logging (debugging aids)
	// to the console
	DebugEnvVarName = "DEBUG"

	// TestsEnvVarName is used to determine if we are running in a test environment
	IsRunningTestsEnvVarName = "IS_RUNNING_TESTS"
)

var (

	// TimeLocation default timezone
	TimeLocation, _ = time.LoadLocation("Africa/Nairobi")
)
