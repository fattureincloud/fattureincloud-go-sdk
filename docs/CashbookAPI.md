# \CashbookAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCashbookEntry**](CashbookAPI.md#CreateCashbookEntry) | **Post** /c/{company_id}/cashbook | Create Cashbook Entry
[**DeleteCashbookEntry**](CashbookAPI.md#DeleteCashbookEntry) | **Delete** /c/{company_id}/cashbook/{document_id} | Delete Cashbook Entry
[**GetCashbookEntry**](CashbookAPI.md#GetCashbookEntry) | **Get** /c/{company_id}/cashbook/{document_id} | Get Cashbook Entry
[**ListCashbookEntries**](CashbookAPI.md#ListCashbookEntries) | **Get** /c/{company_id}/cashbook | List Cashbook Entries
[**ModifyCashbookEntry**](CashbookAPI.md#ModifyCashbookEntry) | **Put** /c/{company_id}/cashbook/{document_id} | Modify Cashbook Entry



## CreateCashbookEntry

> CreateCashbookEntryResponse CreateCashbookEntry(ctx, companyId).CreateCashbookEntryRequest(createCashbookEntryRequest).Execute()

Create Cashbook Entry



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
    createCashbookEntryRequest := *fattureincloud.NewCreateCashbookEntryRequest() // CreateCashbookEntryRequest | Cashbook entry.  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.CashbookAPI.CreateCashbookEntry(auth, companyId).CreateCashbookEntryRequest(createCashbookEntryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashbookAPI.CreateCashbookEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCashbookEntry`: CreateCashbookEntryResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCashbookEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCashbookEntryRequest** | [**CreateCashbookEntryRequest**](CreateCashbookEntryRequest.md) | Cashbook entry.  | 

### Return type

[**CreateCashbookEntryResponse**](CreateCashbookEntryResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCashbookEntry

> DeleteCashbookEntry(ctx, companyId, documentId).Execute()

Delete Cashbook Entry



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
    documentId := "documentId_example" // string | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.CashbookAPI.DeleteCashbookEntry(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashbookAPI.DeleteCashbookEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **string** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCashbookEntryRequest struct via the builder pattern


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


## GetCashbookEntry

> GetCashbookEntryResponse GetCashbookEntry(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get Cashbook Entry



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
    documentId := "documentId_example" // string | The ID of the document.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.CashbookAPI.GetCashbookEntry(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashbookAPI.GetCashbookEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCashbookEntry`: GetCashbookEntryResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **string** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCashbookEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetCashbookEntryResponse**](GetCashbookEntryResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCashbookEntries

> ListCashbookEntriesResponse ListCashbookEntries(ctx, companyId).DateFrom(dateFrom).DateTo(dateTo).Year(year).Type_(type_).PaymentAccountId(paymentAccountId).Execute()

List Cashbook Entries



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
    dateFrom := "dateFrom_example" // string | Start date.
    dateTo := "dateTo_example" // string | End date.
    year := int32(56) // int32 | Filter cashbook by year. (optional)
    type_ := "type__example" // string | Filter cashbook by type. (optional)
    paymentAccountId := int32(56) // int32 | Filter by payment account. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.CashbookAPI.ListCashbookEntries(auth, companyId).DateFrom(dateFrom).DateTo(dateTo).Year(year).Type_(type_).PaymentAccountId(paymentAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashbookAPI.ListCashbookEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCashbookEntries`: ListCashbookEntriesResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCashbookEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateFrom** | **string** | Start date. | 
 **dateTo** | **string** | End date. | 
 **year** | **int32** | Filter cashbook by year. | 
 **type_** | **string** | Filter cashbook by type. | 
 **paymentAccountId** | **int32** | Filter by payment account. | 

### Return type

[**ListCashbookEntriesResponse**](ListCashbookEntriesResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCashbookEntry

> ModifyCashbookEntryResponse ModifyCashbookEntry(ctx, companyId, documentId).ModifyCashbookEntryRequest(modifyCashbookEntryRequest).Execute()

Modify Cashbook Entry



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
    documentId := "documentId_example" // string | The ID of the document.
    modifyCashbookEntryRequest := *fattureincloud.NewModifyCashbookEntryRequest() // ModifyCashbookEntryRequest | Cashbook Entry (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.CashbookAPI.ModifyCashbookEntry(auth, companyId, documentId).ModifyCashbookEntryRequest(modifyCashbookEntryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CashbookAPI.ModifyCashbookEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyCashbookEntry`: ModifyCashbookEntryResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**documentId** | **string** | The ID of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCashbookEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyCashbookEntryRequest** | [**ModifyCashbookEntryRequest**](ModifyCashbookEntryRequest.md) | Cashbook Entry | 

### Return type

[**ModifyCashbookEntryResponse**](ModifyCashbookEntryResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

