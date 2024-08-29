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

type ErrorsTestSuite struct {
	suite.Suite
}

func TestErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorsTestSuite))
}

func (suite *ErrorsTestSuite) TestMapToUserFacing_NilError() {
	err := MapToUserFacing(nil, nil)
	assert.Nil(suite.T(), err)
}

func (suite *ErrorsTestSuite) TestMapToUserFacing_GenericOpenAPIError() {
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

func (suite *ErrorsTestSuite) TestMapToUserFacing_UnknownError() {
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
