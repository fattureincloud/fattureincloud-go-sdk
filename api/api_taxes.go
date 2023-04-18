/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
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
	"os"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)


// TaxesApiService TaxesApi service
type TaxesApiService service

type ApiCreateF24Request struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	createF24Request *CreateF24Request
}

// The F24 to create
func (r ApiCreateF24Request) CreateF24Request(createF24Request CreateF24Request) ApiCreateF24Request {
	r.createF24Request = &createF24Request
	return r
}

func (r ApiCreateF24Request) Execute() (*CreateF24Response, *http.Response, error) {
	return r.ApiService.CreateF24Execute(r)
}

/*
CreateF24 Create F24

Creates a new F24.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiCreateF24Request
*/
func (a *TaxesApiService) CreateF24(ctx context.Context, companyId int32) ApiCreateF24Request {
	return ApiCreateF24Request{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateF24Response
func (a *TaxesApiService) CreateF24Execute(r ApiCreateF24Request) (*CreateF24Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue *CreateF24Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.CreateF24")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes"
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
	localVarPostBody = r.createF24Request
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

type ApiDeleteF24Request struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	documentId int32
}

func (r ApiDeleteF24Request) Execute() (*http.Response, error) {
	return r.ApiService.DeleteF24Execute(r)
}

/*
DeleteF24 Delete F24

Removes the specified F24.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiDeleteF24Request
*/
func (a *TaxesApiService) DeleteF24(ctx context.Context, companyId int32, documentId int32) ApiDeleteF24Request {
	return ApiDeleteF24Request{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *TaxesApiService) DeleteF24Execute(r ApiDeleteF24Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.DeleteF24")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes/{document_id}"
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

type ApiDeleteF24AttachmentRequest struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	documentId int32
}

func (r ApiDeleteF24AttachmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteF24AttachmentExecute(r)
}

/*
DeleteF24Attachment Delete F24 Attachment

Removes the attachment of the specified F24.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiDeleteF24AttachmentRequest
*/
func (a *TaxesApiService) DeleteF24Attachment(ctx context.Context, companyId int32, documentId int32) ApiDeleteF24AttachmentRequest {
	return ApiDeleteF24AttachmentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *TaxesApiService) DeleteF24AttachmentExecute(r ApiDeleteF24AttachmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.DeleteF24Attachment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes/{document_id}/attachment"
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

type ApiGetF24Request struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	documentId int32
	fields *string
	fieldset *string
}

// List of comma-separated fields.
func (r ApiGetF24Request) Fields(fields string) ApiGetF24Request {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiGetF24Request) Fieldset(fieldset string) ApiGetF24Request {
	r.fieldset = &fieldset
	return r
}

func (r ApiGetF24Request) Execute() (*GetF24Response, *http.Response, error) {
	return r.ApiService.GetF24Execute(r)
}

/*
GetF24 Get F24

Gets the specified F24.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiGetF24Request
*/
func (a *TaxesApiService) GetF24(ctx context.Context, companyId int32, documentId int32) ApiGetF24Request {
	return ApiGetF24Request{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return GetF24Response
func (a *TaxesApiService) GetF24Execute(r ApiGetF24Request) (*GetF24Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue *GetF24Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.GetF24")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

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

type ApiListF24Request struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	fields *string
	fieldset *string
	sort *string
	page *int32
	perPage *int32
	q *string
}

// List of comma-separated fields.
func (r ApiListF24Request) Fields(fields string) ApiListF24Request {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiListF24Request) Fieldset(fieldset string) ApiListF24Request {
	r.fieldset = &fieldset
	return r
}

// List of comma-separated fields for result sorting (minus for desc sorting).
func (r ApiListF24Request) Sort(sort string) ApiListF24Request {
	r.sort = &sort
	return r
}

// The page to retrieve.
func (r ApiListF24Request) Page(page int32) ApiListF24Request {
	r.page = &page
	return r
}

// The size of the page.
func (r ApiListF24Request) PerPage(perPage int32) ApiListF24Request {
	r.perPage = &perPage
	return r
}

// Query for filtering the results.
func (r ApiListF24Request) Q(q string) ApiListF24Request {
	r.q = &q
	return r
}

func (r ApiListF24Request) Execute() (*ListF24Response, *http.Response, error) {
	return r.ApiService.ListF24Execute(r)
}

/*
ListF24 List F24

Lists the F24s.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiListF24Request
*/
func (a *TaxesApiService) ListF24(ctx context.Context, companyId int32) ApiListF24Request {
	return ApiListF24Request{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListF24Response
func (a *TaxesApiService) ListF24Execute(r ApiListF24Request) (*ListF24Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue *ListF24Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.ListF24")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes"
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
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
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

type ApiModifyF24Request struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	documentId int32
	modifyF24Request *ModifyF24Request
}

// The F24
func (r ApiModifyF24Request) ModifyF24Request(modifyF24Request ModifyF24Request) ApiModifyF24Request {
	r.modifyF24Request = &modifyF24Request
	return r
}

func (r ApiModifyF24Request) Execute() (*ModifyF24Response, *http.Response, error) {
	return r.ApiService.ModifyF24Execute(r)
}

/*
ModifyF24 Modify F24

Modifies the specified F24.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiModifyF24Request
*/
func (a *TaxesApiService) ModifyF24(ctx context.Context, companyId int32, documentId int32) ApiModifyF24Request {
	return ApiModifyF24Request{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return ModifyF24Response
func (a *TaxesApiService) ModifyF24Execute(r ApiModifyF24Request) (*ModifyF24Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue *ModifyF24Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.ModifyF24")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes/{document_id}"
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
	localVarPostBody = r.modifyF24Request
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

type ApiUploadF24AttachmentRequest struct {
	ctx context.Context
	ApiService *TaxesApiService
	companyId int32
	filename *string
	attachment *os.File
}

// Name of the file.
func (r ApiUploadF24AttachmentRequest) Filename(filename string) ApiUploadF24AttachmentRequest {
	r.filename = &filename
	return r
}

// Valid format: .png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx
func (r ApiUploadF24AttachmentRequest) Attachment(attachment *os.File) ApiUploadF24AttachmentRequest {
	r.attachment = attachment
	return r
}

func (r ApiUploadF24AttachmentRequest) Execute() (*UploadF24AttachmentResponse, *http.Response, error) {
	return r.ApiService.UploadF24AttachmentExecute(r)
}

/*
UploadF24Attachment Upload F24 Attachment

Uploads an attachment destined to a F24. The actual association between the document and the attachment must be implemented separately, using the returned token.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiUploadF24AttachmentRequest
*/
func (a *TaxesApiService) UploadF24Attachment(ctx context.Context, companyId int32) ApiUploadF24AttachmentRequest {
	return ApiUploadF24AttachmentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return UploadF24AttachmentResponse
func (a *TaxesApiService) UploadF24AttachmentExecute(r ApiUploadF24AttachmentRequest) (*UploadF24AttachmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue *UploadF24AttachmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxesApiService.UploadF24Attachment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/taxes/attachment"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "filename", r.filename, "")
	}
	var attachmentLocalVarFormFileName string
	var attachmentLocalVarFileName     string
	var attachmentLocalVarFileBytes    []byte

	attachmentLocalVarFormFileName = "attachment"


	attachmentLocalVarFile := r.attachment

	if attachmentLocalVarFile != nil {
		fbs, _ := io.ReadAll(attachmentLocalVarFile)

		attachmentLocalVarFileBytes = fbs
		attachmentLocalVarFileName = attachmentLocalVarFile.Name()
		attachmentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: attachmentLocalVarFileBytes, fileName: attachmentLocalVarFileName, formFileName: attachmentLocalVarFormFileName})
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
