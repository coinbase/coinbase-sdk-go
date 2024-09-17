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


type ExternalAddressesAPI interface {

	/*
	GetExternalAddressBalance Get the balance of an asset in an external address

	Get the balance of an asset in an external address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the blockchain network
	@param addressId The ID of the address to fetch the balance for
	@param assetId The ID of the asset to fetch the balance for
	@return ApiGetExternalAddressBalanceRequest
	*/
	GetExternalAddressBalance(ctx context.Context, networkId string, addressId string, assetId string) ApiGetExternalAddressBalanceRequest

	// GetExternalAddressBalanceExecute executes the request
	//  @return Balance
	GetExternalAddressBalanceExecute(r ApiGetExternalAddressBalanceRequest) (*Balance, *http.Response, error)

	/*
	ListAddressTransactions List transactions for an address.

	List all transactions that interact with the address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the blockchain network
	@param addressId The ID of the address to fetch the transactions for.
	@return ApiListAddressTransactionsRequest
	*/
	ListAddressTransactions(ctx context.Context, networkId string, addressId string) ApiListAddressTransactionsRequest

	// ListAddressTransactionsExecute executes the request
	//  @return AddressTransactionList
	ListAddressTransactionsExecute(r ApiListAddressTransactionsRequest) (*AddressTransactionList, *http.Response, error)

	/*
	ListExternalAddressBalances Get the balances of an external address

	List all of the balances of an external address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the blockchain network
	@param addressId The ID of the address to fetch the balance for
	@return ApiListExternalAddressBalancesRequest
	*/
	ListExternalAddressBalances(ctx context.Context, networkId string, addressId string) ApiListExternalAddressBalancesRequest

	// ListExternalAddressBalancesExecute executes the request
	//  @return AddressBalanceList
	ListExternalAddressBalancesExecute(r ApiListExternalAddressBalancesRequest) (*AddressBalanceList, *http.Response, error)

	/*
	RequestExternalFaucetFunds Request faucet funds for external address.

	Request faucet funds to be sent to external address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param networkId The ID of the wallet the address belongs to.
	@param addressId The onchain address of the address that is being fetched.
	@return ApiRequestExternalFaucetFundsRequest
	*/
	RequestExternalFaucetFunds(ctx context.Context, networkId string, addressId string) ApiRequestExternalFaucetFundsRequest

	// RequestExternalFaucetFundsExecute executes the request
	//  @return FaucetTransaction
	RequestExternalFaucetFundsExecute(r ApiRequestExternalFaucetFundsRequest) (*FaucetTransaction, *http.Response, error)
}

// ExternalAddressesAPIService ExternalAddressesAPI service
type ExternalAddressesAPIService service

type ApiGetExternalAddressBalanceRequest struct {
	ctx context.Context
	ApiService ExternalAddressesAPI
	networkId string
	addressId string
	assetId string
}

func (r ApiGetExternalAddressBalanceRequest) Execute() (*Balance, *http.Response, error) {
	return r.ApiService.GetExternalAddressBalanceExecute(r)
}

/*
GetExternalAddressBalance Get the balance of an asset in an external address

Get the balance of an asset in an external address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the blockchain network
 @param addressId The ID of the address to fetch the balance for
 @param assetId The ID of the asset to fetch the balance for
 @return ApiGetExternalAddressBalanceRequest
*/
func (a *ExternalAddressesAPIService) GetExternalAddressBalance(ctx context.Context, networkId string, addressId string, assetId string) ApiGetExternalAddressBalanceRequest {
	return ApiGetExternalAddressBalanceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		addressId: addressId,
		assetId: assetId,
	}
}

// Execute executes the request
//  @return Balance
func (a *ExternalAddressesAPIService) GetExternalAddressBalanceExecute(r ApiGetExternalAddressBalanceRequest) (*Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAddressesAPIService.GetExternalAddressBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/addresses/{address_id}/balances/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)

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

type ApiListAddressTransactionsRequest struct {
	ctx context.Context
	ApiService ExternalAddressesAPI
	networkId string
	addressId string
	limit *int32
	page *string
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.
func (r ApiListAddressTransactionsRequest) Limit(limit int32) ApiListAddressTransactionsRequest {
	r.limit = &limit
	return r
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListAddressTransactionsRequest) Page(page string) ApiListAddressTransactionsRequest {
	r.page = &page
	return r
}

func (r ApiListAddressTransactionsRequest) Execute() (*AddressTransactionList, *http.Response, error) {
	return r.ApiService.ListAddressTransactionsExecute(r)
}

/*
ListAddressTransactions List transactions for an address.

List all transactions that interact with the address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the blockchain network
 @param addressId The ID of the address to fetch the transactions for.
 @return ApiListAddressTransactionsRequest
*/
func (a *ExternalAddressesAPIService) ListAddressTransactions(ctx context.Context, networkId string, addressId string) ApiListAddressTransactionsRequest {
	return ApiListAddressTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return AddressTransactionList
func (a *ExternalAddressesAPIService) ListAddressTransactionsExecute(r ApiListAddressTransactionsRequest) (*AddressTransactionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddressTransactionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAddressesAPIService.ListAddressTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/addresses/{address_id}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListExternalAddressBalancesRequest struct {
	ctx context.Context
	ApiService ExternalAddressesAPI
	networkId string
	addressId string
	page *string
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListExternalAddressBalancesRequest) Page(page string) ApiListExternalAddressBalancesRequest {
	r.page = &page
	return r
}

func (r ApiListExternalAddressBalancesRequest) Execute() (*AddressBalanceList, *http.Response, error) {
	return r.ApiService.ListExternalAddressBalancesExecute(r)
}

/*
ListExternalAddressBalances Get the balances of an external address

List all of the balances of an external address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the blockchain network
 @param addressId The ID of the address to fetch the balance for
 @return ApiListExternalAddressBalancesRequest
*/
func (a *ExternalAddressesAPIService) ListExternalAddressBalances(ctx context.Context, networkId string, addressId string) ApiListExternalAddressBalancesRequest {
	return ApiListExternalAddressBalancesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return AddressBalanceList
func (a *ExternalAddressesAPIService) ListExternalAddressBalancesExecute(r ApiListExternalAddressBalancesRequest) (*AddressBalanceList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddressBalanceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAddressesAPIService.ListExternalAddressBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/addresses/{address_id}/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiRequestExternalFaucetFundsRequest struct {
	ctx context.Context
	ApiService ExternalAddressesAPI
	networkId string
	addressId string
	assetId *string
}

// The ID of the asset to transfer from the faucet.
func (r ApiRequestExternalFaucetFundsRequest) AssetId(assetId string) ApiRequestExternalFaucetFundsRequest {
	r.assetId = &assetId
	return r
}

func (r ApiRequestExternalFaucetFundsRequest) Execute() (*FaucetTransaction, *http.Response, error) {
	return r.ApiService.RequestExternalFaucetFundsExecute(r)
}

/*
RequestExternalFaucetFunds Request faucet funds for external address.

Request faucet funds to be sent to external address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The ID of the wallet the address belongs to.
 @param addressId The onchain address of the address that is being fetched.
 @return ApiRequestExternalFaucetFundsRequest
*/
func (a *ExternalAddressesAPIService) RequestExternalFaucetFunds(ctx context.Context, networkId string, addressId string) ApiRequestExternalFaucetFundsRequest {
	return ApiRequestExternalFaucetFundsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return FaucetTransaction
func (a *ExternalAddressesAPIService) RequestExternalFaucetFundsExecute(r ApiRequestExternalFaucetFundsRequest) (*FaucetTransaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FaucetTransaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAddressesAPIService.RequestExternalFaucetFunds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/networks/{network_id}/addresses/{address_id}/faucet"
	localVarPath = strings.Replace(localVarPath, "{"+"network_id"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId, "form", "")
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
