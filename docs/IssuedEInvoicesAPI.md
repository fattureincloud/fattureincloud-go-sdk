# \IssuedEInvoicesAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEInvoiceRejectionReason**](IssuedEInvoicesAPI.md#GetEInvoiceRejectionReason) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/error_reason | Get E-Invoice Rejection Reason
[**GetEInvoiceXml**](IssuedEInvoicesAPI.md#GetEInvoiceXml) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/xml | Get E-Invoice XML
[**SendEInvoice**](IssuedEInvoicesAPI.md#SendEInvoice) | **Post** /c/{company_id}/issued_documents/{document_id}/e_invoice/send | Send E-Invoice
[**VerifyEInvoiceXml**](IssuedEInvoicesAPI.md#VerifyEInvoiceXml) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/xml_verify | Verify E-Invoice XML



## GetEInvoiceRejectionReason

> GetEInvoiceRejectionReasonResponse GetEInvoiceRejectionReason(ctx, companyId, documentId).Execute()

Get E-Invoice Rejection Reason



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloudapi "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedEInvoicesAPI.GetEInvoiceRejectionReason(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedEInvoicesAPI.GetEInvoiceRejectionReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEInvoiceRejectionReason`: GetEInvoiceRejectionReasonResponse
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

Other parameters are passed through a pointer to a apiGetEInvoiceRejectionReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEInvoiceRejectionReasonResponse**](GetEInvoiceRejectionReasonResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEInvoiceXml

> string GetEInvoiceXml(ctx, companyId, documentId).IncludeAttachment(includeAttachment).Execute()

Get E-Invoice XML



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloudapi "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.
    includeAttachment := true // bool | Include the attachment to the XML e-invoice. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedEInvoicesAPI.GetEInvoiceXml(auth, companyId, documentId).IncludeAttachment(includeAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedEInvoicesAPI.GetEInvoiceXml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEInvoiceXml`: string
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

Other parameters are passed through a pointer to a apiGetEInvoiceXmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAttachment** | **bool** | Include the attachment to the XML e-invoice. | 

### Return type

**string**

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEInvoice

> SendEInvoiceResponse SendEInvoice(ctx, companyId, documentId).SendEInvoiceRequest(sendEInvoiceRequest).Execute()

Send E-Invoice



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloudapi "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.
    sendEInvoiceRequest := *fattureincloud.NewSendEInvoiceRequest() // SendEInvoiceRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedEInvoicesAPI.SendEInvoice(auth, companyId, documentId).SendEInvoiceRequest(sendEInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedEInvoicesAPI.SendEInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEInvoice`: SendEInvoiceResponse
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

Other parameters are passed through a pointer to a apiSendEInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendEInvoiceRequest** | [**SendEInvoiceRequest**](SendEInvoiceRequest.md) |  | 

### Return type

[**SendEInvoiceResponse**](SendEInvoiceResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEInvoiceXml

> VerifyEInvoiceXmlResponse VerifyEInvoiceXml(ctx, companyId, documentId).Execute()

Verify E-Invoice XML



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    fattureincloudapi "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
    fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
)

func main() {
    companyId := int32(12345) // int32 | The ID of the company.
    documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedEInvoicesAPI.VerifyEInvoiceXml(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuedEInvoicesAPI.VerifyEInvoiceXml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEInvoiceXml`: VerifyEInvoiceXmlResponse
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

Other parameters are passed through a pointer to a apiVerifyEInvoiceXmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VerifyEInvoiceXmlResponse**](VerifyEInvoiceXmlResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

