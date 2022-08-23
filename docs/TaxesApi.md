# \TaxesApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateF24**](TaxesApi.md#CreateF24) | **Post** /c/{company_id}/taxes | Create F24
[**DeleteF24**](TaxesApi.md#DeleteF24) | **Delete** /c/{company_id}/taxes/{document_id} | Delete F24
[**DeleteF24Attachment**](TaxesApi.md#DeleteF24Attachment) | **Delete** /c/{company_id}/taxes/{document_id}/attachment | Delete F24 Attachment
[**GetF24**](TaxesApi.md#GetF24) | **Get** /c/{company_id}/taxes/{document_id} | Get F24
[**ListF24**](TaxesApi.md#ListF24) | **Get** /c/{company_id}/taxes | List F24
[**ModifyF24**](TaxesApi.md#ModifyF24) | **Put** /c/{company_id}/taxes/{document_id} | Modify F24
[**UploadF24Attachment**](TaxesApi.md#UploadF24Attachment) | **Post** /c/{company_id}/taxes/attachment | Upload F24 Attachment



## CreateF24

> CreateF24Response CreateF24(ctx, companyId).CreateF24Request(createF24Request).Execute()

Create F24



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
    createF24Request := *fattureincloud.NewCreateF24Request() // CreateF24Request | The F24 to create (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.CreateF24(auth, companyId).CreateF24Request(createF24Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.CreateF24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateF24`: CreateF24Response
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateF24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createF24Request** | [**CreateF24Request**](CreateF24Request.md) | The F24 to create | 

### Return type

[**CreateF24Response**](CreateF24Response.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteF24

> DeleteF24(ctx, companyId, documentId).Execute()

Delete F24



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
    resp, r, err := apiClient.TaxesApi.DeleteF24(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.DeleteF24``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteF24Request struct via the builder pattern


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


## DeleteF24Attachment

> DeleteF24Attachment(ctx, companyId, documentId).Execute()

Delete F24 Attachment



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
    resp, r, err := apiClient.TaxesApi.DeleteF24Attachment(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.DeleteF24Attachment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteF24AttachmentRequest struct via the builder pattern


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


## GetF24

> GetF24Response GetF24(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get F24



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
    resp, r, err := apiClient.TaxesApi.GetF24(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.GetF24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetF24`: GetF24Response
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

Other parameters are passed through a pointer to a apiGetF24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetF24Response**](GetF24Response.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListF24

> ListF24Response ListF24(ctx, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()

List F24



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
    resp, r, err := apiClient.TaxesApi.ListF24(auth, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.ListF24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListF24`: ListF24Response
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListF24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListF24Response**](ListF24Response.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyF24

> ModifyF24Response ModifyF24(ctx, companyId, documentId).ModifyF24Request(modifyF24Request).Execute()

Modify F24



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
    modifyF24Request := *fattureincloud.NewModifyF24Request() // ModifyF24Request | The F24 (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.ModifyF24(auth, companyId, documentId).ModifyF24Request(modifyF24Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.ModifyF24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyF24`: ModifyF24Response
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

Other parameters are passed through a pointer to a apiModifyF24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyF24Request** | [**ModifyF24Request**](ModifyF24Request.md) | The F24 | 

### Return type

[**ModifyF24Response**](ModifyF24Response.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadF24Attachment

> UploadF24AttachmentResponse UploadF24Attachment(ctx, companyId).Filename(filename).Attachment(attachment).Execute()

Upload F24 Attachment



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
    resp, r, err := apiClient.TaxesApi.UploadF24Attachment(auth, companyId).Filename(filename).Attachment(attachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.UploadF24Attachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadF24Attachment`: UploadF24AttachmentResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadF24AttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | Name of the file. | 
 **attachment** | ***os.File** | Valid format: .png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx | 

### Return type

[**UploadF24AttachmentResponse**](UploadF24AttachmentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

