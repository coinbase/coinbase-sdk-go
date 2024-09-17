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


type WalletsAPI interface {

	/*
	CreateWallet Create a new wallet

	Create a new wallet scoped to the user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWalletRequest
	*/
	CreateWallet(ctx context.Context) ApiCreateWalletRequest

	// CreateWalletExecute executes the request
	//  @return Wallet
	CreateWalletExecute(r ApiCreateWalletRequest) (*Wallet, *http.Response, error)

	/*
	GetWallet Get wallet by ID

	Get wallet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to fetch
	@return ApiGetWalletRequest
	*/
	GetWallet(ctx context.Context, walletId string) ApiGetWalletRequest

	// GetWalletExecute executes the request
	//  @return Wallet
	GetWalletExecute(r ApiGetWalletRequest) (*Wallet, *http.Response, error)

	/*
	GetWalletBalance Get the balance of an asset in the wallet

	Get the aggregated balance of an asset across all of the addresses in the wallet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to fetch the balance for
	@param assetId The symbol of the asset to fetch the balance for
	@return ApiGetWalletBalanceRequest
	*/
	GetWalletBalance(ctx context.Context, walletId string, assetId string) ApiGetWalletBalanceRequest

	// GetWalletBalanceExecute executes the request
	//  @return Balance
	GetWalletBalanceExecute(r ApiGetWalletBalanceRequest) (*Balance, *http.Response, error)

	/*
	ListWalletBalances List wallet balances

	List the balances of all of the addresses in the wallet aggregated by asset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to fetch the balances for
	@return ApiListWalletBalancesRequest
	*/
	ListWalletBalances(ctx context.Context, walletId string) ApiListWalletBalancesRequest

	// ListWalletBalancesExecute executes the request
	//  @return AddressBalanceList
	ListWalletBalancesExecute(r ApiListWalletBalancesRequest) (*AddressBalanceList, *http.Response, error)

	/*
	ListWallets List wallets

	List wallets belonging to the user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListWalletsRequest
	*/
	ListWallets(ctx context.Context) ApiListWalletsRequest

	// ListWalletsExecute executes the request
	//  @return WalletList
	ListWalletsExecute(r ApiListWalletsRequest) (*WalletList, *http.Response, error)
}

// WalletsAPIService WalletsAPI service
type WalletsAPIService service

type ApiCreateWalletRequest struct {
	ctx context.Context
	ApiService WalletsAPI
	createWalletRequest *CreateWalletRequest
}

func (r ApiCreateWalletRequest) CreateWalletRequest(createWalletRequest CreateWalletRequest) ApiCreateWalletRequest {
	r.createWalletRequest = &createWalletRequest
	return r
}

func (r ApiCreateWalletRequest) Execute() (*Wallet, *http.Response, error) {
	return r.ApiService.CreateWalletExecute(r)
}

/*
CreateWallet Create a new wallet

Create a new wallet scoped to the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWalletRequest
*/
func (a *WalletsAPIService) CreateWallet(ctx context.Context) ApiCreateWalletRequest {
	return ApiCreateWalletRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Wallet
func (a *WalletsAPIService) CreateWalletExecute(r ApiCreateWalletRequest) (*Wallet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Wallet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsAPIService.CreateWallet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createWalletRequest
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

type ApiGetWalletRequest struct {
	ctx context.Context
	ApiService WalletsAPI
	walletId string
}

func (r ApiGetWalletRequest) Execute() (*Wallet, *http.Response, error) {
	return r.ApiService.GetWalletExecute(r)
}

/*
GetWallet Get wallet by ID

Get wallet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to fetch
 @return ApiGetWalletRequest
*/
func (a *WalletsAPIService) GetWallet(ctx context.Context, walletId string) ApiGetWalletRequest {
	return ApiGetWalletRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return Wallet
func (a *WalletsAPIService) GetWalletExecute(r ApiGetWalletRequest) (*Wallet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Wallet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsAPIService.GetWallet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

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

type ApiGetWalletBalanceRequest struct {
	ctx context.Context
	ApiService WalletsAPI
	walletId string
	assetId string
}

func (r ApiGetWalletBalanceRequest) Execute() (*Balance, *http.Response, error) {
	return r.ApiService.GetWalletBalanceExecute(r)
}

/*
GetWalletBalance Get the balance of an asset in the wallet

Get the aggregated balance of an asset across all of the addresses in the wallet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to fetch the balance for
 @param assetId The symbol of the asset to fetch the balance for
 @return ApiGetWalletBalanceRequest
*/
func (a *WalletsAPIService) GetWalletBalance(ctx context.Context, walletId string, assetId string) ApiGetWalletBalanceRequest {
	return ApiGetWalletBalanceRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		assetId: assetId,
	}
}

// Execute executes the request
//  @return Balance
func (a *WalletsAPIService) GetWalletBalanceExecute(r ApiGetWalletBalanceRequest) (*Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsAPIService.GetWalletBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/balances/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
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

type ApiListWalletBalancesRequest struct {
	ctx context.Context
	ApiService WalletsAPI
	walletId string
}

func (r ApiListWalletBalancesRequest) Execute() (*AddressBalanceList, *http.Response, error) {
	return r.ApiService.ListWalletBalancesExecute(r)
}

/*
ListWalletBalances List wallet balances

List the balances of all of the addresses in the wallet aggregated by asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to fetch the balances for
 @return ApiListWalletBalancesRequest
*/
func (a *WalletsAPIService) ListWalletBalances(ctx context.Context, walletId string) ApiListWalletBalancesRequest {
	return ApiListWalletBalancesRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return AddressBalanceList
func (a *WalletsAPIService) ListWalletBalancesExecute(r ApiListWalletBalancesRequest) (*AddressBalanceList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddressBalanceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsAPIService.ListWalletBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

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

type ApiListWalletsRequest struct {
	ctx context.Context
	ApiService WalletsAPI
	limit *int32
	page *string
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.
func (r ApiListWalletsRequest) Limit(limit int32) ApiListWalletsRequest {
	r.limit = &limit
	return r
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListWalletsRequest) Page(page string) ApiListWalletsRequest {
	r.page = &page
	return r
}

func (r ApiListWalletsRequest) Execute() (*WalletList, *http.Response, error) {
	return r.ApiService.ListWalletsExecute(r)
}

/*
ListWallets List wallets

List wallets belonging to the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWalletsRequest
*/
func (a *WalletsAPIService) ListWallets(ctx context.Context) ApiListWalletsRequest {
	return ApiListWalletsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WalletList
func (a *WalletsAPIService) ListWalletsExecute(r ApiListWalletsRequest) (*WalletList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WalletList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsAPIService.ListWallets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets"

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
