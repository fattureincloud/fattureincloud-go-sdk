/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.30
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)


// SuppliersAPIService SuppliersAPI service
type SuppliersAPIService service

type ApiCreateSupplierRequest struct {
	ctx context.Context
	ApiService *SuppliersAPIService
	companyId int32
	createSupplierRequest *CreateSupplierRequest
}

// The supplier to create
func (r ApiCreateSupplierRequest) CreateSupplierRequest(createSupplierRequest CreateSupplierRequest) ApiCreateSupplierRequest {
	r.createSupplierRequest = &createSupplierRequest
	return r
}

func (r ApiCreateSupplierRequest) Execute() (*CreateSupplierResponse, *http.Response, error) {
	return r.ApiService.CreateSupplierExecute(r)
}

/*
CreateSupplier Create Supplier

Creates a new supplier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiCreateSupplierRequest
*/
func (a *SuppliersAPIService) CreateSupplier(ctx context.Context, companyId int32) ApiCreateSupplierRequest {
	return ApiCreateSupplierRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateSupplierResponse
func (a *SuppliersAPIService) CreateSupplierExecute(r ApiCreateSupplierRequest) (*CreateSupplierResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateSupplierResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppliersAPIService.CreateSupplier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/suppliers"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

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
	localVarPostBody = r.createSupplierRequest
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

type ApiDeleteSupplierRequest struct {
	ctx context.Context
	ApiService *SuppliersAPIService
	companyId int32
	supplierId int32
}

func (r ApiDeleteSupplierRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSupplierExecute(r)
}

/*
DeleteSupplier Delete Supplier

Deletes the specified supplier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param supplierId The ID of the supplier.
 @return ApiDeleteSupplierRequest
*/
func (a *SuppliersAPIService) DeleteSupplier(ctx context.Context, companyId int32, supplierId int32) ApiDeleteSupplierRequest {
	return ApiDeleteSupplierRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		supplierId: supplierId,
	}
}

// Execute executes the request
func (a *SuppliersAPIService) DeleteSupplierExecute(r ApiDeleteSupplierRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppliersAPIService.DeleteSupplier")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/suppliers/{supplier_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"supplier_id"+"}", url.PathEscape(parameterValueToString(r.supplierId, "supplierId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSupplierRequest struct {
	ctx context.Context
	ApiService *SuppliersAPIService
	companyId int32
	supplierId int32
	fields *string
	fieldset *string
}

// List of comma-separated fields.
func (r ApiGetSupplierRequest) Fields(fields string) ApiGetSupplierRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiGetSupplierRequest) Fieldset(fieldset string) ApiGetSupplierRequest {
	r.fieldset = &fieldset
	return r
}

func (r ApiGetSupplierRequest) Execute() (*GetSupplierResponse, *http.Response, error) {
	return r.ApiService.GetSupplierExecute(r)
}

/*
GetSupplier Get Supplier

Gets the specified supplier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param supplierId The ID of the supplier.
 @return ApiGetSupplierRequest
*/
func (a *SuppliersAPIService) GetSupplier(ctx context.Context, companyId int32, supplierId int32) ApiGetSupplierRequest {
	return ApiGetSupplierRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		supplierId: supplierId,
	}
}

// Execute executes the request
//  @return GetSupplierResponse
func (a *SuppliersAPIService) GetSupplierExecute(r ApiGetSupplierRequest) (*GetSupplierResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSupplierResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppliersAPIService.GetSupplier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/suppliers/{supplier_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"supplier_id"+"}", url.PathEscape(parameterValueToString(r.supplierId, "supplierId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.fieldset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldset", r.fieldset, "")
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

type ApiListSuppliersRequest struct {
	ctx context.Context
	ApiService *SuppliersAPIService
	companyId int32
	fields *string
	fieldset *string
	sort *string
	page *int32
	perPage *int32
	q *string
}

// List of comma-separated fields.
func (r ApiListSuppliersRequest) Fields(fields string) ApiListSuppliersRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiListSuppliersRequest) Fieldset(fieldset string) ApiListSuppliersRequest {
	r.fieldset = &fieldset
	return r
}

// List of comma-separated fields for result sorting (minus for desc sorting).
func (r ApiListSuppliersRequest) Sort(sort string) ApiListSuppliersRequest {
	r.sort = &sort
	return r
}

// The page to retrieve.
func (r ApiListSuppliersRequest) Page(page int32) ApiListSuppliersRequest {
	r.page = &page
	return r
}

// The size of the page.
func (r ApiListSuppliersRequest) PerPage(perPage int32) ApiListSuppliersRequest {
	r.perPage = &perPage
	return r
}

// Query for filtering the results.
func (r ApiListSuppliersRequest) Q(q string) ApiListSuppliersRequest {
	r.q = &q
	return r
}

func (r ApiListSuppliersRequest) Execute() (*ListSuppliersResponse, *http.Response, error) {
	return r.ApiService.ListSuppliersExecute(r)
}

/*
ListSuppliers List Suppliers

Lists the suppliers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiListSuppliersRequest
*/
func (a *SuppliersAPIService) ListSuppliers(ctx context.Context, companyId int32) ApiListSuppliersRequest {
	return ApiListSuppliersRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListSuppliersResponse
func (a *SuppliersAPIService) ListSuppliersExecute(r ApiListSuppliersRequest) (*ListSuppliersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListSuppliersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppliersAPIService.ListSuppliers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/suppliers"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.fieldset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldset", r.fieldset, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
 		var defaultValue int32 = 1
 		r.page = &defaultValue
 	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	} else {
 		var defaultValue int32 = 5
 		r.perPage = &defaultValue
 	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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

type ApiModifySupplierRequest struct {
	ctx context.Context
	ApiService *SuppliersAPIService
	companyId int32
	supplierId int32
	modifySupplierRequest *ModifySupplierRequest
}

// The modified Supplier. First level parameters are managed in delta mode.
func (r ApiModifySupplierRequest) ModifySupplierRequest(modifySupplierRequest ModifySupplierRequest) ApiModifySupplierRequest {
	r.modifySupplierRequest = &modifySupplierRequest
	return r
}

func (r ApiModifySupplierRequest) Execute() (*ModifySupplierResponse, *http.Response, error) {
	return r.ApiService.ModifySupplierExecute(r)
}

/*
ModifySupplier Modify Supplier

Modifies the specified supplier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param supplierId The ID of the supplier.
 @return ApiModifySupplierRequest
*/
func (a *SuppliersAPIService) ModifySupplier(ctx context.Context, companyId int32, supplierId int32) ApiModifySupplierRequest {
	return ApiModifySupplierRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		supplierId: supplierId,
	}
}

// Execute executes the request
//  @return ModifySupplierResponse
func (a *SuppliersAPIService) ModifySupplierExecute(r ApiModifySupplierRequest) (*ModifySupplierResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModifySupplierResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuppliersAPIService.ModifySupplier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/entities/suppliers/{supplier_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"supplier_id"+"}", url.PathEscape(parameterValueToString(r.supplierId, "supplierId")), -1)

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
	localVarPostBody = r.modifySupplierRequest
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
