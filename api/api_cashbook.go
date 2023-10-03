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


// CashbookAPIService CashbookAPI service
type CashbookAPIService service

type ApiCreateCashbookEntryRequest struct {
	ctx context.Context
	ApiService *CashbookAPIService
	companyId int32
	createCashbookEntryRequest *CreateCashbookEntryRequest
}

// Cashbook entry. 
func (r ApiCreateCashbookEntryRequest) CreateCashbookEntryRequest(createCashbookEntryRequest CreateCashbookEntryRequest) ApiCreateCashbookEntryRequest {
	r.createCashbookEntryRequest = &createCashbookEntryRequest
	return r
}

func (r ApiCreateCashbookEntryRequest) Execute() (*CreateCashbookEntryResponse, *http.Response, error) {
	return r.ApiService.CreateCashbookEntryExecute(r)
}

/*
CreateCashbookEntry Create Cashbook Entry

Creates a new cashbook entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiCreateCashbookEntryRequest
*/
func (a *CashbookAPIService) CreateCashbookEntry(ctx context.Context, companyId int32) ApiCreateCashbookEntryRequest {
	return ApiCreateCashbookEntryRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateCashbookEntryResponse
func (a *CashbookAPIService) CreateCashbookEntryExecute(r ApiCreateCashbookEntryRequest) (*CreateCashbookEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateCashbookEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CashbookAPIService.CreateCashbookEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/cashbook"
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
	localVarPostBody = r.createCashbookEntryRequest
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

type ApiDeleteCashbookEntryRequest struct {
	ctx context.Context
	ApiService *CashbookAPIService
	companyId int32
	documentId string
}

func (r ApiDeleteCashbookEntryRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCashbookEntryExecute(r)
}

/*
DeleteCashbookEntry Delete Cashbook Entry

Deletes the specified cashbook entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiDeleteCashbookEntryRequest
*/
func (a *CashbookAPIService) DeleteCashbookEntry(ctx context.Context, companyId int32, documentId string) ApiDeleteCashbookEntryRequest {
	return ApiDeleteCashbookEntryRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *CashbookAPIService) DeleteCashbookEntryExecute(r ApiDeleteCashbookEntryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CashbookAPIService.DeleteCashbookEntry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/cashbook/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

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

type ApiGetCashbookEntryRequest struct {
	ctx context.Context
	ApiService *CashbookAPIService
	companyId int32
	documentId string
	fields *string
	fieldset *string
}

// List of comma-separated fields.
func (r ApiGetCashbookEntryRequest) Fields(fields string) ApiGetCashbookEntryRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiGetCashbookEntryRequest) Fieldset(fieldset string) ApiGetCashbookEntryRequest {
	r.fieldset = &fieldset
	return r
}

func (r ApiGetCashbookEntryRequest) Execute() (*GetCashbookEntryResponse, *http.Response, error) {
	return r.ApiService.GetCashbookEntryExecute(r)
}

/*
GetCashbookEntry Get Cashbook Entry

Gets the specified cashbook entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiGetCashbookEntryRequest
*/
func (a *CashbookAPIService) GetCashbookEntry(ctx context.Context, companyId int32, documentId string) ApiGetCashbookEntryRequest {
	return ApiGetCashbookEntryRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return GetCashbookEntryResponse
func (a *CashbookAPIService) GetCashbookEntryExecute(r ApiGetCashbookEntryRequest) (*GetCashbookEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCashbookEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CashbookAPIService.GetCashbookEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/cashbook/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	}
	if r.fieldset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldset", r.fieldset, "")
	}
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

type ApiListCashbookEntriesRequest struct {
	ctx context.Context
	ApiService *CashbookAPIService
	companyId int32
	dateFrom *string
	dateTo *string
	year *int32
	type_ *string
	paymentAccountId *int32
}

// Start date.
func (r ApiListCashbookEntriesRequest) DateFrom(dateFrom string) ApiListCashbookEntriesRequest {
	r.dateFrom = &dateFrom
	return r
}

// End date.
func (r ApiListCashbookEntriesRequest) DateTo(dateTo string) ApiListCashbookEntriesRequest {
	r.dateTo = &dateTo
	return r
}

// Filter cashbook by year.
func (r ApiListCashbookEntriesRequest) Year(year int32) ApiListCashbookEntriesRequest {
	r.year = &year
	return r
}

// Filter cashbook by type.
func (r ApiListCashbookEntriesRequest) Type_(type_ string) ApiListCashbookEntriesRequest {
	r.type_ = &type_
	return r
}

// Filter by payment account.
func (r ApiListCashbookEntriesRequest) PaymentAccountId(paymentAccountId int32) ApiListCashbookEntriesRequest {
	r.paymentAccountId = &paymentAccountId
	return r
}

func (r ApiListCashbookEntriesRequest) Execute() (*ListCashbookEntriesResponse, *http.Response, error) {
	return r.ApiService.ListCashbookEntriesExecute(r)
}

/*
ListCashbookEntries List Cashbook Entries

Lists the cashbook entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiListCashbookEntriesRequest
*/
func (a *CashbookAPIService) ListCashbookEntries(ctx context.Context, companyId int32) ApiListCashbookEntriesRequest {
	return ApiListCashbookEntriesRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListCashbookEntriesResponse
func (a *CashbookAPIService) ListCashbookEntriesExecute(r ApiListCashbookEntriesRequest) (*ListCashbookEntriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCashbookEntriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CashbookAPIService.ListCashbookEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/cashbook"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dateFrom == nil {
		return localVarReturnValue, nil, reportError("dateFrom is required and must be specified")
	}
	if r.dateTo == nil {
		return localVarReturnValue, nil, reportError("dateTo is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "date_from", r.dateFrom, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "date_to", r.dateTo, "")
	if r.year != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "year", r.year, "")
	}
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	}
	if r.paymentAccountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payment_account_id", r.paymentAccountId, "")
	}
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

type ApiModifyCashbookEntryRequest struct {
	ctx context.Context
	ApiService *CashbookAPIService
	companyId int32
	documentId string
	modifyCashbookEntryRequest *ModifyCashbookEntryRequest
}

// Cashbook Entry
func (r ApiModifyCashbookEntryRequest) ModifyCashbookEntryRequest(modifyCashbookEntryRequest ModifyCashbookEntryRequest) ApiModifyCashbookEntryRequest {
	r.modifyCashbookEntryRequest = &modifyCashbookEntryRequest
	return r
}

func (r ApiModifyCashbookEntryRequest) Execute() (*ModifyCashbookEntryResponse, *http.Response, error) {
	return r.ApiService.ModifyCashbookEntryExecute(r)
}

/*
ModifyCashbookEntry Modify Cashbook Entry

Modifies the specified cashbook entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiModifyCashbookEntryRequest
*/
func (a *CashbookAPIService) ModifyCashbookEntry(ctx context.Context, companyId int32, documentId string) ApiModifyCashbookEntryRequest {
	return ApiModifyCashbookEntryRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return ModifyCashbookEntryResponse
func (a *CashbookAPIService) ModifyCashbookEntryExecute(r ApiModifyCashbookEntryRequest) (*ModifyCashbookEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModifyCashbookEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CashbookAPIService.ModifyCashbookEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/cashbook/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

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
	localVarPostBody = r.modifyCashbookEntryRequest
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
