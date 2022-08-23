# \ArchiveApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchiveDocument**](ArchiveApi.md#CreateArchiveDocument) | **Post** /c/{company_id}/archive | Create Archive Document
[**DeleteArchiveDocument**](ArchiveApi.md#DeleteArchiveDocument) | **Delete** /c/{company_id}/archive/{document_id} | Delete Archive Document
[**GetArchiveDocument**](ArchiveApi.md#GetArchiveDocument) | **Get** /c/{company_id}/archive/{document_id} | Get Archive Document
[**ListArchiveDocuments**](ArchiveApi.md#ListArchiveDocuments) | **Get** /c/{company_id}/archive | List Archive Documents
[**ModifyArchiveDocument**](ArchiveApi.md#ModifyArchiveDocument) | **Put** /c/{company_id}/archive/{document_id} | Modify Archive Document
[**UploadArchiveDocumentAttachment**](ArchiveApi.md#UploadArchiveDocumentAttachment) | **Post** /c/{company_id}/archive/attachment | Upload Archive Document Attachment



## CreateArchiveDocument

> CreateArchiveDocumentResponse CreateArchiveDocument(ctx, companyId).CreateArchiveDocumentRequest(createArchiveDocumentRequest).Execute()

Create Archive Document



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    createArchiveDocumentRequest := *fattureincloud.NewCreateArchiveDocumentRequest() // CreateArchiveDocumentRequest | The Archive Document. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.CreateArchiveDocument(auth, companyId).CreateArchiveDocumentRequest(createArchiveDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.CreateArchiveDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArchiveDocument`: CreateArchiveDocumentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArchiveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createArchiveDocumentRequest** | [**CreateArchiveDocumentRequest**](CreateArchiveDocumentRequest.md) | The Archive Document. | 

### Return type

[**CreateArchiveDocumentResponse**](CreateArchiveDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArchiveDocument

> DeleteArchiveDocument(ctx, companyId, documentId).Execute()

Delete Archive Document



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.DeleteArchiveDocument(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.DeleteArchiveDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **int32** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchiveDocument

> GetArchiveDocumentResponse GetArchiveDocument(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get Archive Document



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.GetArchiveDocument(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.GetArchiveDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveDocument`: GetArchiveDocumentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **int32** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetArchiveDocumentResponse**](GetArchiveDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchiveDocuments

> ListArchiveDocumentsResponse ListArchiveDocuments(ctx, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()

List Archive Documents



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)
    sort := "sort_example" // string | List of comma-separated fields for result sorting (minus for desc sorting). (optional)
    page := int32(56) // int32 | The page to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | The size of the page. (optional) (default to 5)
    q := "q_example" // string | Query for filtering the results. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.ListArchiveDocuments(auth, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.ListArchiveDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveDocuments`: ListArchiveDocumentsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListArchiveDocumentsResponse**](ListArchiveDocumentsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyArchiveDocument

> ModifyArchiveDocumentResponse ModifyArchiveDocument(ctx, companyId, documentId).ModifyArchiveDocumentRequest(modifyArchiveDocumentRequest).Execute()

Modify Archive Document



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.
    modifyArchiveDocumentRequest := *fattureincloud.NewModifyArchiveDocumentRequest() // ModifyArchiveDocumentRequest | Modified Archive Document (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.ModifyArchiveDocument(auth, companyId, documentId).ModifyArchiveDocumentRequest(modifyArchiveDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.ModifyArchiveDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyArchiveDocument`: ModifyArchiveDocumentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **int32** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyArchiveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyArchiveDocumentRequest** | [**ModifyArchiveDocumentRequest**](ModifyArchiveDocumentRequest.md) | Modified Archive Document | 

### Return type

[**ModifyArchiveDocumentResponse**](ModifyArchiveDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadArchiveDocumentAttachment

> UploadArchiveAttachmentResponse UploadArchiveDocumentAttachment(ctx, companyId).Filename(filename).Attachment(attachment).Execute()

Upload Archive Document Attachment



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    filename := "filename_example" // string | Name of the file. (optional)
    attachment := os.NewFile(1234, "some_file") // *os.File | Valid format: .png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.UploadArchiveDocumentAttachment(auth, companyId).Filename(filename).Attachment(attachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.UploadArchiveDocumentAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadArchiveDocumentAttachment`: UploadArchiveAttachmentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadArchiveDocumentAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | Name of the file. | 
 **attachment** | ***os.File** | Valid format: .png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx | 

### Return type

[**UploadArchiveAttachmentResponse**](UploadArchiveAttachmentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

