package errors

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"unsafe"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MapErrorsTestSuite struct {
	suite.Suite
}

func TestMapErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(MapErrorsTestSuite))
}

func (suite *MapErrorsTestSuite) TestMapToUserFacing_NilError() {
	err := MapToUserFacing(nil, nil)
	assert.Nil(suite.T(), err)
}

func (suite *MapErrorsTestSuite) TestMapToUserFacing_GenericOpenAPIError() {
	body := []byte(`{"code": "test_code", "message": "test_message"}`)
	openAPIError := createGenericOpenAPIError(body)
	resp := &http.Response{StatusCode: http.StatusBadRequest}

	err := MapToUserFacing(&openAPIError, resp)
	var apiErr *APIError
	ok := errors.As(err, &apiErr)

	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "test_code", apiErr.Code)
	assert.Equal(suite.T(), "test_message", apiErr.Message)
	assert.Equal(suite.T(), http.StatusBadRequest, apiErr.HttpStatusCode)
}

func (suite *MapErrorsTestSuite) TestMapToUserFacing_UnknownError() {
	unknownErr := errors.New("unknown error")
	resp := &http.Response{StatusCode: http.StatusInternalServerError}

	err := MapToUserFacing(unknownErr, resp)
	var apiErr *APIError
	ok := errors.As(err, &apiErr)

	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), "unknown", apiErr.Code)
	assert.Equal(suite.T(), "unknown error", apiErr.Message)
	assert.Equal(suite.T(), http.StatusInternalServerError, apiErr.HttpStatusCode)
}

// createGenericOpenAPIError creates a GenericOpenAPIError with the given body.
// We can do this client.GenericOpenAPIError{body: body} as the body field is unexported.
// This is a workaround to use reflection to set the body field.
func createGenericOpenAPIError(body []byte) client.GenericOpenAPIError {
	openAPIError := client.GenericOpenAPIError{}
	bodyField := reflect.ValueOf(&openAPIError).Elem().FieldByName("body")
	reflect.NewAt(bodyField.Type(), unsafe.Pointer(bodyField.UnsafeAddr())).Elem().SetBytes(body)
	return openAPIError
}

type APIErrorTestSuite struct {
	suite.Suite
}

func TestAPIErrorTestSuite(t *testing.T) {
	suite.Run(t, new(APIErrorTestSuite))
}

func (suite *APIErrorTestSuite) TestError() {
	tests := []struct {
		name     string
		apiError APIError
		expected string
	}{
		{
			name: "All fields set",
			apiError: APIError{
				HttpStatusCode: 400,
				Code:           "test_code",
				Message:        "test_message",
				CorrelationId:  "test_correlation_id",
			},
			expected: "APIError{HttpStatusCode: 400, Code: test_code, Message: test_message, CorrelationId: test_correlation_id}",
		},
		{
			name: "Only HttpStatusCode set",
			apiError: APIError{
				HttpStatusCode: 400,
			},
			expected: "APIError{HttpStatusCode: 400}",
		},
		{
			name: "Only Code set",
			apiError: APIError{
				Code: "test_code",
			},
			expected: "APIError{Code: test_code}",
		},
		{
			name: "Only Message set",
			apiError: APIError{
				Message: "test_message",
			},
			expected: "APIError{Message: test_message}",
		},
		{
			name: "Only CorrelationId set",
			apiError: APIError{
				CorrelationId: "test_correlation_id",
			},
			expected: "APIError{CorrelationId: test_correlation_id}",
		},
		{
			name:     "No fields set",
			apiError: APIError{},
			expected: "APIError{}",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			assert.Equal(suite.T(), tt.expected, tt.apiError.Error())
		})
	}
}
