/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type ValidatorsAPI interface {

	/*
	GetValidator Get a validator belonging to the CDP project

	Get a validator belonging to the user for a given network, asset and id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the blockchain network.
	@param assetId The symbol of the asset to get the validator for.
	@param validatorId The unique id of the validator to fetch details for.
	@return ApiGetValidatorRequest
	*/
	GetValidator(ctx context.Context, networkId string, assetId string, validatorId string) ApiGetValidatorRequest

	// GetValidatorExecute executes the request
	//  @return Validator
	GetValidatorExecute(r ApiGetValidatorRequest) (*Validator, *http.Response, error)

	/*
	ListValidators List validators belonging to the CDP project

	List validators belonging to the user for a given network and asset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the blockchain network.
	@param assetId The symbol of the asset to get the validators for.
	@return ApiListValidatorsRequest
	*/
	ListValidators(ctx context.Context, networkId string, assetId string) ApiListValidatorsRequest

	// ListValidatorsExecute executes the request
	//  @return ValidatorList
	ListValidatorsExecute(r ApiListValidatorsRequest) (*ValidatorList, *http.Response, error)
}

// ValidatorsAPIService ValidatorsAPI service
type ValidatorsAPIService service

type ApiGetValidatorRequest struct {
	ctx context.Context
	ApiService ValidatorsAPI
	networkId string
	assetId string
	validatorId string
}

func (r ApiGetValidatorRequest) Execute() (*Validator, *http.Response, error) {
	return r.ApiService.GetValidatorExecute(r)
}

/*
GetValidator Get a validator belonging to the CDP project

Get a validator belonging to the user for a given network, asset and id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the blockchain network.
 @param assetId The symbol of the asset to get the validator for.
 @param validatorId The unique id of the validator to fetch details for.
 @return ApiGetValidatorRequest
*/
func (a *ValidatorsAPIService) GetValidator(ctx context.Context, networkId string, assetId string, validatorId string) ApiGetValidatorRequest {
	return ApiGetValidatorRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		assetId: assetId,
		validatorId: validatorId,
	}
}

// Execute executes the request
//  @return Validator
func (a *ValidatorsAPIService) GetValidatorExecute(r ApiGetValidatorRequest) (*Validator, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Validator
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ValidatorsAPIService.GetValidator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/assets/{asset_id}/validators/{validator_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"validator_id"+"}", url.PathEscape(parameterValueToString(r.validatorId, "validatorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListValidatorsRequest struct {
	ctx context.Context
	ApiService ValidatorsAPI
	networkId string
	assetId string
	status *ValidatorStatus
	limit *int32
	page *string
}

// A filter to list validators based on a status.
func (r ApiListValidatorsRequest) Status(status ValidatorStatus) ApiListValidatorsRequest {
	r.status = &status
	return r
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 50.
func (r ApiListValidatorsRequest) Limit(limit int32) ApiListValidatorsRequest {
	r.limit = &limit
	return r
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListValidatorsRequest) Page(page string) ApiListValidatorsRequest {
	r.page = &page
	return r
}

func (r ApiListValidatorsRequest) Execute() (*ValidatorList, *http.Response, error) {
	return r.ApiService.ListValidatorsExecute(r)
}

/*
ListValidators List validators belonging to the CDP project

List validators belonging to the user for a given network and asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the blockchain network.
 @param assetId The symbol of the asset to get the validators for.
 @return ApiListValidatorsRequest
*/
func (a *ValidatorsAPIService) ListValidators(ctx context.Context, networkId string, assetId string) ApiListValidatorsRequest {
	return ApiListValidatorsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		assetId: assetId,
	}
}

// Execute executes the request
//  @return ValidatorList
func (a *ValidatorsAPIService) ListValidatorsExecute(r ApiListValidatorsRequest) (*ValidatorList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ValidatorList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ValidatorsAPIService.ListValidators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/assets/{asset_id}/validators"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
