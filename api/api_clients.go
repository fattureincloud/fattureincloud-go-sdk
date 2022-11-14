/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.22
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)


// ClientsApiService ClientsApi service
type ClientsApiService service

type ApiCreateClientRequest struct {
	ctx context.Context
	ApiService *ClientsApiService
	companyId int32
	createClientRequest *CreateClientRequest
}

// The client to create
func (r ApiCreateClientRequest) CreateClientRequest(createClientRequest CreateClientRequest) ApiCreateClientRequest {
	r.createClientRequest = &createClientRequest
	return r
}

func (r ApiCreateClientRequest) Execute() (*CreateClientResponse, *http.Response, error) {
	return r.ApiService.CreateClientExecute(r)
}

/*
CreateClient Create Client

Creates a new client.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiCreateClientRequest
*/
func (a *ClientsApiService) CreateClient(ctx context.Context, companyId int32) ApiCreateClientRequest {
	return ApiCreateClientRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateClientResponse
func (a *ClientsApiService) CreateClientExecute(r ApiCreateClientRequest) (*CreateClientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsApiService.CreateClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/clients"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

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
	localVarPostBody = r.createClientRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
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

type ApiDeleteClientRequest struct {
	ctx context.Context
	ApiService *ClientsApiService
	companyId int32
	clientId int32
}

func (r ApiDeleteClientRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClientExecute(r)
}

/*
DeleteClient Delete Client

Deletes the specified client.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param clientId The ID of the client.
 @return ApiDeleteClientRequest
*/
func (a *ClientsApiService) DeleteClient(ctx context.Context, companyId int32, clientId int32) ApiDeleteClientRequest {
	return ApiDeleteClientRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		clientId: clientId,
	}
}

// Execute executes the request
func (a *ClientsApiService) DeleteClientExecute(r ApiDeleteClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsApiService.DeleteClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/clients/{client_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client_id"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetClientRequest struct {
	ctx context.Context
	ApiService *ClientsApiService
	companyId int32
	clientId int32
	fields *string
	fieldset *string
}

// List of comma-separated fields.
func (r ApiGetClientRequest) Fields(fields string) ApiGetClientRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiGetClientRequest) Fieldset(fieldset string) ApiGetClientRequest {
	r.fieldset = &fieldset
	return r
}

func (r ApiGetClientRequest) Execute() (*GetClientResponse, *http.Response, error) {
	return r.ApiService.GetClientExecute(r)
}

/*
GetClient Get Client

Gets the specified client.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param clientId The ID of the client.
 @return ApiGetClientRequest
*/
func (a *ClientsApiService) GetClient(ctx context.Context, companyId int32, clientId int32) ApiGetClientRequest {
	return ApiGetClientRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return GetClientResponse
func (a *ClientsApiService) GetClientExecute(r ApiGetClientRequest) (*GetClientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsApiService.GetClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/clients/{client_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client_id"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.fieldset != nil {
		localVarQueryParams.Add("fieldset", parameterToString(*r.fieldset, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
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

type ApiListClientsRequest struct {
	ctx context.Context
	ApiService *ClientsApiService
	companyId int32
	fields *string
	fieldset *string
	sort *string
	page *int32
	perPage *int32
	q *string
}

// List of comma-separated fields.
func (r ApiListClientsRequest) Fields(fields string) ApiListClientsRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiListClientsRequest) Fieldset(fieldset string) ApiListClientsRequest {
	r.fieldset = &fieldset
	return r
}

// List of comma-separated fields for result sorting (minus for desc sorting).
func (r ApiListClientsRequest) Sort(sort string) ApiListClientsRequest {
	r.sort = &sort
	return r
}

// The page to retrieve.
func (r ApiListClientsRequest) Page(page int32) ApiListClientsRequest {
	r.page = &page
	return r
}

// The size of the page.
func (r ApiListClientsRequest) PerPage(perPage int32) ApiListClientsRequest {
	r.perPage = &perPage
	return r
}

// Query for filtering the results.
func (r ApiListClientsRequest) Q(q string) ApiListClientsRequest {
	r.q = &q
	return r
}

func (r ApiListClientsRequest) Execute() (*ListClientsResponse, *http.Response, error) {
	return r.ApiService.ListClientsExecute(r)
}

/*
ListClients List Clients

Lists the clients.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiListClientsRequest
*/
func (a *ClientsApiService) ListClients(ctx context.Context, companyId int32) ApiListClientsRequest {
	return ApiListClientsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListClientsResponse
func (a *ClientsApiService) ListClientsExecute(r ApiListClientsRequest) (*ListClientsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListClientsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsApiService.ListClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/clients"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.fieldset != nil {
		localVarQueryParams.Add("fieldset", parameterToString(*r.fieldset, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
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

type ApiModifyClientRequest struct {
	ctx context.Context
	ApiService *ClientsApiService
	companyId int32
	clientId int32
	modifyClientRequest *ModifyClientRequest
}

// The modified Client. First level parameters are managed in delta mode.
func (r ApiModifyClientRequest) ModifyClientRequest(modifyClientRequest ModifyClientRequest) ApiModifyClientRequest {
	r.modifyClientRequest = &modifyClientRequest
	return r
}

func (r ApiModifyClientRequest) Execute() (*ModifyClientResponse, *http.Response, error) {
	return r.ApiService.ModifyClientExecute(r)
}

/*
ModifyClient Modify Client

Modifies the specified client.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param clientId The ID of the client.
 @return ApiModifyClientRequest
*/
func (a *ClientsApiService) ModifyClient(ctx context.Context, companyId int32, clientId int32) ApiModifyClientRequest {
	return ApiModifyClientRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ModifyClientResponse
func (a *ClientsApiService) ModifyClientExecute(r ApiModifyClientRequest) (*ModifyClientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModifyClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsApiService.ModifyClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/clients/{client_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client_id"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

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
	localVarPostBody = r.modifyClientRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
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
