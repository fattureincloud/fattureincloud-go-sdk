/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
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


// ArchiveAPIService ArchiveAPI service
type ArchiveAPIService service

type ApiCreateArchiveDocumentRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	createArchiveDocumentRequest *CreateArchiveDocumentRequest
}

// The Archive Document.
func (r ApiCreateArchiveDocumentRequest) CreateArchiveDocumentRequest(createArchiveDocumentRequest CreateArchiveDocumentRequest) ApiCreateArchiveDocumentRequest {
	r.createArchiveDocumentRequest = &createArchiveDocumentRequest
	return r
}

func (r ApiCreateArchiveDocumentRequest) Execute() (*CreateArchiveDocumentResponse, *http.Response, error) {
	return r.ApiService.CreateArchiveDocumentExecute(r)
}

/*
CreateArchiveDocument Create Archive Document

Creates a new archive document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiCreateArchiveDocumentRequest
*/
func (a *ArchiveAPIService) CreateArchiveDocument(ctx context.Context, companyId int32) ApiCreateArchiveDocumentRequest {
	return ApiCreateArchiveDocumentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return CreateArchiveDocumentResponse
func (a *ArchiveAPIService) CreateArchiveDocumentExecute(r ApiCreateArchiveDocumentRequest) (*CreateArchiveDocumentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateArchiveDocumentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.CreateArchiveDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive"
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
	localVarPostBody = r.createArchiveDocumentRequest
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

type ApiDeleteArchiveDocumentRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	documentId int32
}

func (r ApiDeleteArchiveDocumentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArchiveDocumentExecute(r)
}

/*
DeleteArchiveDocument Delete Archive Document

Deletes the specified archive document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiDeleteArchiveDocumentRequest
*/
func (a *ArchiveAPIService) DeleteArchiveDocument(ctx context.Context, companyId int32, documentId int32) ApiDeleteArchiveDocumentRequest {
	return ApiDeleteArchiveDocumentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
func (a *ArchiveAPIService) DeleteArchiveDocumentExecute(r ApiDeleteArchiveDocumentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.DeleteArchiveDocument")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive/{document_id}"
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

type ApiGetArchiveDocumentRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	documentId int32
	fields *string
	fieldset *string
}

// List of comma-separated fields.
func (r ApiGetArchiveDocumentRequest) Fields(fields string) ApiGetArchiveDocumentRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiGetArchiveDocumentRequest) Fieldset(fieldset string) ApiGetArchiveDocumentRequest {
	r.fieldset = &fieldset
	return r
}

func (r ApiGetArchiveDocumentRequest) Execute() (*GetArchiveDocumentResponse, *http.Response, error) {
	return r.ApiService.GetArchiveDocumentExecute(r)
}

/*
GetArchiveDocument Get Archive Document

Gets the specified archive document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiGetArchiveDocumentRequest
*/
func (a *ArchiveAPIService) GetArchiveDocument(ctx context.Context, companyId int32, documentId int32) ApiGetArchiveDocumentRequest {
	return ApiGetArchiveDocumentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return GetArchiveDocumentResponse
func (a *ArchiveAPIService) GetArchiveDocumentExecute(r ApiGetArchiveDocumentRequest) (*GetArchiveDocumentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetArchiveDocumentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.GetArchiveDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive/{document_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"document_id"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.fieldset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldset", r.fieldset, "form", "")
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

type ApiListArchiveDocumentsRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	fields *string
	fieldset *string
	sort *string
	page *int32
	perPage *int32
	q *string
}

// List of comma-separated fields.
func (r ApiListArchiveDocumentsRequest) Fields(fields string) ApiListArchiveDocumentsRequest {
	r.fields = &fields
	return r
}

// Name of the fieldset.
func (r ApiListArchiveDocumentsRequest) Fieldset(fieldset string) ApiListArchiveDocumentsRequest {
	r.fieldset = &fieldset
	return r
}

// List of comma-separated fields for result sorting (minus for desc sorting).
func (r ApiListArchiveDocumentsRequest) Sort(sort string) ApiListArchiveDocumentsRequest {
	r.sort = &sort
	return r
}

// The page to retrieve.
func (r ApiListArchiveDocumentsRequest) Page(page int32) ApiListArchiveDocumentsRequest {
	r.page = &page
	return r
}

// The size of the page.
func (r ApiListArchiveDocumentsRequest) PerPage(perPage int32) ApiListArchiveDocumentsRequest {
	r.perPage = &perPage
	return r
}

// Query for filtering the results.
func (r ApiListArchiveDocumentsRequest) Q(q string) ApiListArchiveDocumentsRequest {
	r.q = &q
	return r
}

func (r ApiListArchiveDocumentsRequest) Execute() (*ListArchiveDocumentsResponse, *http.Response, error) {
	return r.ApiService.ListArchiveDocumentsExecute(r)
}

/*
ListArchiveDocuments List Archive Documents

Lists the archive documents.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiListArchiveDocumentsRequest
*/
func (a *ArchiveAPIService) ListArchiveDocuments(ctx context.Context, companyId int32) ApiListArchiveDocumentsRequest {
	return ApiListArchiveDocumentsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListArchiveDocumentsResponse
func (a *ArchiveAPIService) ListArchiveDocumentsExecute(r ApiListArchiveDocumentsRequest) (*ListArchiveDocumentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListArchiveDocumentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.ListArchiveDocuments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", url.PathEscape(parameterValueToString(r.companyId, "companyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.fieldset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fieldset", r.fieldset, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "form", "")
	} else {
		var defaultValue int32 = 5
		r.perPage = &defaultValue
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
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

type ApiModifyArchiveDocumentRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	documentId int32
	modifyArchiveDocumentRequest *ModifyArchiveDocumentRequest
}

// Modified Archive Document
func (r ApiModifyArchiveDocumentRequest) ModifyArchiveDocumentRequest(modifyArchiveDocumentRequest ModifyArchiveDocumentRequest) ApiModifyArchiveDocumentRequest {
	r.modifyArchiveDocumentRequest = &modifyArchiveDocumentRequest
	return r
}

func (r ApiModifyArchiveDocumentRequest) Execute() (*ModifyArchiveDocumentResponse, *http.Response, error) {
	return r.ApiService.ModifyArchiveDocumentExecute(r)
}

/*
ModifyArchiveDocument Modify Archive Document

Modifies the specified archive document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @param documentId The ID of the document.
 @return ApiModifyArchiveDocumentRequest
*/
func (a *ArchiveAPIService) ModifyArchiveDocument(ctx context.Context, companyId int32, documentId int32) ApiModifyArchiveDocumentRequest {
	return ApiModifyArchiveDocumentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		documentId: documentId,
	}
}

// Execute executes the request
//  @return ModifyArchiveDocumentResponse
func (a *ArchiveAPIService) ModifyArchiveDocumentExecute(r ApiModifyArchiveDocumentRequest) (*ModifyArchiveDocumentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModifyArchiveDocumentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.ModifyArchiveDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive/{document_id}"
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
	localVarPostBody = r.modifyArchiveDocumentRequest
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

type ApiUploadArchiveDocumentAttachmentRequest struct {
	ctx context.Context
	ApiService *ArchiveAPIService
	companyId int32
	filename *string
	attachment *os.File
}

// Attachment file name
func (r ApiUploadArchiveDocumentAttachmentRequest) Filename(filename string) ApiUploadArchiveDocumentAttachmentRequest {
	r.filename = &filename
	return r
}

// Attachment file [.png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx]
func (r ApiUploadArchiveDocumentAttachmentRequest) Attachment(attachment *os.File) ApiUploadArchiveDocumentAttachmentRequest {
	r.attachment = attachment
	return r
}

func (r ApiUploadArchiveDocumentAttachmentRequest) Execute() (*UploadArchiveAttachmentResponse, *http.Response, error) {
	return r.ApiService.UploadArchiveDocumentAttachmentExecute(r)
}

/*
UploadArchiveDocumentAttachment Upload Archive Document Attachment

Uploads an attachment destined to an archive document. The actual association between the document and the attachment must be implemented separately, using the returned token.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The ID of the company.
 @return ApiUploadArchiveDocumentAttachmentRequest
*/
func (a *ArchiveAPIService) UploadArchiveDocumentAttachment(ctx context.Context, companyId int32) ApiUploadArchiveDocumentAttachmentRequest {
	return ApiUploadArchiveDocumentAttachmentRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return UploadArchiveAttachmentResponse
func (a *ArchiveAPIService) UploadArchiveDocumentAttachmentExecute(r ApiUploadArchiveDocumentAttachmentRequest) (*UploadArchiveAttachmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadArchiveAttachmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveAPIService.UploadArchiveDocumentAttachment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/c/{company_id}/archive/attachment"
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
		parameterAddToHeaderOrQuery(localVarFormParams, "filename", r.filename, "", "")
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
