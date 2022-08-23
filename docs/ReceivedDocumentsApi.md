# \ReceivedDocumentsApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceivedDocument**](ReceivedDocumentsApi.md#CreateReceivedDocument) | **Post** /c/{company_id}/received_documents | Create Received Document
[**DeleteReceivedDocument**](ReceivedDocumentsApi.md#DeleteReceivedDocument) | **Delete** /c/{company_id}/received_documents/{document_id} | Delete Received Document
[**DeleteReceivedDocumentAttachment**](ReceivedDocumentsApi.md#DeleteReceivedDocumentAttachment) | **Delete** /c/{company_id}/received_documents/{document_id}/attachment | Delete Received Document Attachment
[**GetExistingReceivedDocumentTotals**](ReceivedDocumentsApi.md#GetExistingReceivedDocumentTotals) | **Post** /c/{company_id}/received_documents/{document_id}/totals | Get Existing Received Document Totals
[**GetNewReceivedDocumentTotals**](ReceivedDocumentsApi.md#GetNewReceivedDocumentTotals) | **Post** /c/{company_id}/received_documents/totals | Get New Received Document Totals
[**GetReceivedDocument**](ReceivedDocumentsApi.md#GetReceivedDocument) | **Get** /c/{company_id}/received_documents/{document_id} | Get Received Document
[**GetReceivedDocumentPreCreateInfo**](ReceivedDocumentsApi.md#GetReceivedDocumentPreCreateInfo) | **Get** /c/{company_id}/received_documents/info | Get Received Document Pre-Create Info
[**ListReceivedDocuments**](ReceivedDocumentsApi.md#ListReceivedDocuments) | **Get** /c/{company_id}/received_documents | List Received Documents
[**ModifyReceivedDocument**](ReceivedDocumentsApi.md#ModifyReceivedDocument) | **Put** /c/{company_id}/received_documents/{document_id} | Modify Received Document
[**UploadReceivedDocumentAttachment**](ReceivedDocumentsApi.md#UploadReceivedDocumentAttachment) | **Post** /c/{company_id}/received_documents/attachment | Upload Received Document Attachment



## CreateReceivedDocument

> CreateReceivedDocumentResponse CreateReceivedDocument(ctx, companyId).CreateReceivedDocumentRequest(createReceivedDocumentRequest).Execute()

Create Received Document



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
    createReceivedDocumentRequest := *fattureincloud.NewCreateReceivedDocumentRequest() // CreateReceivedDocumentRequest | Document to create (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.CreateReceivedDocument(auth, companyId).CreateReceivedDocumentRequest(createReceivedDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.CreateReceivedDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceivedDocument`: CreateReceivedDocumentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceivedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReceivedDocumentRequest** | [**CreateReceivedDocumentRequest**](CreateReceivedDocumentRequest.md) | Document to create | 

### Return type

[**CreateReceivedDocumentResponse**](CreateReceivedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReceivedDocument

> DeleteReceivedDocument(ctx, companyId, documentId).Execute()

Delete Received Document



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
    resp, r, err := apiClient.ReceivedDocumentsApi.DeleteReceivedDocument(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.DeleteReceivedDocument``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReceivedDocumentRequest struct via the builder pattern


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


## DeleteReceivedDocumentAttachment

> DeleteReceivedDocumentAttachment(ctx, companyId, documentId).Execute()

Delete Received Document Attachment



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
    resp, r, err := apiClient.ReceivedDocumentsApi.DeleteReceivedDocumentAttachment(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.DeleteReceivedDocumentAttachment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReceivedDocumentAttachmentRequest struct via the builder pattern


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


## GetExistingReceivedDocumentTotals

> GetExistingReceivedDocumentTotalsResponse GetExistingReceivedDocumentTotals(ctx, companyId, documentId).GetExistingReceivedDocumentTotalsRequest(getExistingReceivedDocumentTotalsRequest).Execute()

Get Existing Received Document Totals



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
    getExistingReceivedDocumentTotalsRequest := *fattureincloud.NewGetExistingReceivedDocumentTotalsRequest() // GetExistingReceivedDocumentTotalsRequest | Received document. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.GetExistingReceivedDocumentTotals(auth, companyId, documentId).GetExistingReceivedDocumentTotalsRequest(getExistingReceivedDocumentTotalsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.GetExistingReceivedDocumentTotals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExistingReceivedDocumentTotals`: GetExistingReceivedDocumentTotalsResponse
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

Other parameters are passed through a pointer to a apiGetExistingReceivedDocumentTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getExistingReceivedDocumentTotalsRequest** | [**GetExistingReceivedDocumentTotalsRequest**](GetExistingReceivedDocumentTotalsRequest.md) | Received document. | 

### Return type

[**GetExistingReceivedDocumentTotalsResponse**](GetExistingReceivedDocumentTotalsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewReceivedDocumentTotals

> GetNewReceivedDocumentTotalsResponse GetNewReceivedDocumentTotals(ctx, companyId).GetNewReceivedDocumentTotalsRequest(getNewReceivedDocumentTotalsRequest).Execute()

Get New Received Document Totals



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
    getNewReceivedDocumentTotalsRequest := *fattureincloud.NewGetNewReceivedDocumentTotalsRequest() // GetNewReceivedDocumentTotalsRequest | Received document. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.GetNewReceivedDocumentTotals(auth, companyId).GetNewReceivedDocumentTotalsRequest(getNewReceivedDocumentTotalsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.GetNewReceivedDocumentTotals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNewReceivedDocumentTotals`: GetNewReceivedDocumentTotalsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewReceivedDocumentTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getNewReceivedDocumentTotalsRequest** | [**GetNewReceivedDocumentTotalsRequest**](GetNewReceivedDocumentTotalsRequest.md) | Received document. | 

### Return type

[**GetNewReceivedDocumentTotalsResponse**](GetNewReceivedDocumentTotalsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivedDocument

> GetReceivedDocumentResponse GetReceivedDocument(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get Received Document



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
    resp, r, err := apiClient.ReceivedDocumentsApi.GetReceivedDocument(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.GetReceivedDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivedDocument`: GetReceivedDocumentResponse
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

Other parameters are passed through a pointer to a apiGetReceivedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetReceivedDocumentResponse**](GetReceivedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivedDocumentPreCreateInfo

> GetReceivedDocumentPreCreateInfoResponse GetReceivedDocumentPreCreateInfo(ctx, companyId).Type_(type_).Execute()

Get Received Document Pre-Create Info



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
    type_ := "type__example" // string | The type of the received document.

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.GetReceivedDocumentPreCreateInfo(auth, companyId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.GetReceivedDocumentPreCreateInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivedDocumentPreCreateInfo`: GetReceivedDocumentPreCreateInfoResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivedDocumentPreCreateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of the received document. | 

### Return type

[**GetReceivedDocumentPreCreateInfoResponse**](GetReceivedDocumentPreCreateInfoResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReceivedDocuments

> ListReceivedDocumentsResponse ListReceivedDocuments(ctx, companyId).Type_(type_).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()

List Received Documents



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
    type_ := "type__example" // string | The type of the received document.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)
    sort := "sort_example" // string | List of comma-separated fields for result sorting (minus for desc sorting). (optional)
    page := int32(56) // int32 | The page to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | The size of the page. (optional) (default to 5)
    q := "q_example" // string | Query for filtering the results. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.ListReceivedDocuments(auth, companyId).Type_(type_).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.ListReceivedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReceivedDocuments`: ListReceivedDocumentsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReceivedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of the received document. | 
 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListReceivedDocumentsResponse**](ListReceivedDocumentsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyReceivedDocument

> ModifyReceivedDocumentResponse ModifyReceivedDocument(ctx, companyId, documentId).ModifyReceivedDocumentRequest(modifyReceivedDocumentRequest).Execute()

Modify Received Document



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
    modifyReceivedDocumentRequest := *fattureincloud.NewModifyReceivedDocumentRequest() // ModifyReceivedDocumentRequest | Modified document. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivedDocumentsApi.ModifyReceivedDocument(auth, companyId, documentId).ModifyReceivedDocumentRequest(modifyReceivedDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.ModifyReceivedDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyReceivedDocument`: ModifyReceivedDocumentResponse
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

Other parameters are passed through a pointer to a apiModifyReceivedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyReceivedDocumentRequest** | [**ModifyReceivedDocumentRequest**](ModifyReceivedDocumentRequest.md) | Modified document. | 

### Return type

[**ModifyReceivedDocumentResponse**](ModifyReceivedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadReceivedDocumentAttachment

> UploadReceivedDocumentAttachmentResponse UploadReceivedDocumentAttachment(ctx, companyId).Filename(filename).Attachment(attachment).Execute()

Upload Received Document Attachment



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
    resp, r, err := apiClient.ReceivedDocumentsApi.UploadReceivedDocumentAttachment(auth, companyId).Filename(filename).Attachment(attachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivedDocumentsApi.UploadReceivedDocumentAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadReceivedDocumentAttachment`: UploadReceivedDocumentAttachmentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadReceivedDocumentAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | Name of the file. | 
 **attachment** | ***os.File** | Valid format: .png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx | 

### Return type

[**UploadReceivedDocumentAttachmentResponse**](UploadReceivedDocumentAttachmentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

