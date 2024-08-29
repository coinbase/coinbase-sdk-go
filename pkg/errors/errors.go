package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// APIError is a custom error type for API errors.
type APIError struct {
	HttpStatusCode int
	// A short string representing the reported error. Can be used to handle errors programmatically.
	Code string
	// A human-readable message providing more details about the error.
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError{ httpStatusCode: %d, apiCode: %s, apiMessage: %s }", e.HttpStatusCode, e.Code, e.Message)
}

// MapToUserFacing maps the error to a custom user facing error type.
func MapToUserFacing(err error, resp *http.Response) error {
	if err == nil {
		return nil
	}

	apiError := &APIError{}

	if resp != nil {
		apiError.HttpStatusCode = resp.StatusCode
	}

	var openAPIError *client.GenericOpenAPIError
	if errors.As(err, &openAPIError) {
		var clientError client.Error
		if err := clientError.UnmarshalJSON(openAPIError.Body()); err != nil {
			return &APIError{
				Code:    "unknown",
				Message: err.Error(),
			}
		}

		apiError.Message = clientError.Message
		apiError.Code = clientError.Code
		return apiError
	}

	apiError.Code = "unknown"
	apiError.Message = err.Error()
	return apiError
}
