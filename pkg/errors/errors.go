package errors

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// APIError is a custom error type for API errors.
type APIError struct {
	HttpStatusCode int
	Code           string // A short string representing the reported error. Can be used to handle errors programmatically.
	Message        string // A human-readable message providing more details about the error.
	CorrelationId  string // A correlation ID that can be used to help debug the error.
}

func (e *APIError) Error() string {
	v := reflect.ValueOf(*e)
	typeOfE := v.Type()
	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		if !fieldValue.IsZero() {
			field := fmt.Sprintf("%s: %v", typeOfE.Field(i).Name, fieldValue.Interface())
			fields = append(fields, field)
		}
	}

	return fmt.Sprintf("APIError{%s}", strings.Join(fields, ", "))
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
		if unmarshalErr := clientError.UnmarshalJSON(openAPIError.Body()); unmarshalErr != nil {
			return &APIError{
				Code:    "unknown",
				Message: err.Error(), // relay back the original error message as is.
			}
		}

		apiError.Message = clientError.Message
		apiError.Code = clientError.Code
		if clientError.CorrelationId != nil {
			apiError.CorrelationId = *clientError.CorrelationId
		}

		return apiError
	}

	apiError.Code = "unknown"
	apiError.Message = err.Error()
	return apiError
}
