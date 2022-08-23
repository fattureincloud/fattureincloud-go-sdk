# \ReceiptsApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceipt**](ReceiptsApi.md#CreateReceipt) | **Post** /c/{company_id}/receipts | Create Receipt
[**DeleteReceipt**](ReceiptsApi.md#DeleteReceipt) | **Delete** /c/{company_id}/receipts/{document_id} | Delete Receipt
[**GetReceipt**](ReceiptsApi.md#GetReceipt) | **Get** /c/{company_id}/receipts/{document_id} | Get Receipt
[**GetReceiptPreCreateInfo**](ReceiptsApi.md#GetReceiptPreCreateInfo) | **Get** /c/{company_id}/receipts/info | Get Receipt Pre-Create Info
[**GetReceiptsMonthlyTotals**](ReceiptsApi.md#GetReceiptsMonthlyTotals) | **Get** /c/{company_id}/receipts/monthly_totals | Get Receipts Monthly Totals
[**ListReceipts**](ReceiptsApi.md#ListReceipts) | **Get** /c/{company_id}/receipts | List Receipts
[**ModifyReceipt**](ReceiptsApi.md#ModifyReceipt) | **Put** /c/{company_id}/receipts/{document_id} | Modify Receipt



## CreateReceipt

> CreateReceiptResponse CreateReceipt(ctx, companyId).CreateReceiptRequest(createReceiptRequest).Execute()

Create Receipt



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
    createReceiptRequest := *fattureincloud.NewCreateReceiptRequest() // CreateReceiptRequest | The Receipt to create. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptsApi.CreateReceipt(auth, companyId).CreateReceiptRequest(createReceiptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.CreateReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceipt`: CreateReceiptResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReceiptRequest** | [**CreateReceiptRequest**](CreateReceiptRequest.md) | The Receipt to create. | 

### Return type

[**CreateReceiptResponse**](CreateReceiptResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReceipt

> DeleteReceipt(ctx, companyId, documentId).Execute()

Delete Receipt



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
    resp, r, err := apiClient.ReceiptsApi.DeleteReceipt(auth, companyId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.DeleteReceipt``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReceiptRequest struct via the builder pattern


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


## GetReceipt

> GetReceiptResponse GetReceipt(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get Receipt



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
    resp, r, err := apiClient.ReceiptsApi.GetReceipt(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.GetReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceipt`: GetReceiptResponse
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

Other parameters are passed through a pointer to a apiGetReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetReceiptResponse**](GetReceiptResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptPreCreateInfo

> GetReceiptPreCreateInfoResponse GetReceiptPreCreateInfo(ctx, companyId).Execute()

Get Receipt Pre-Create Info



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

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptsApi.GetReceiptPreCreateInfo(auth, companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.GetReceiptPreCreateInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceiptPreCreateInfo`: GetReceiptPreCreateInfoResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptPreCreateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReceiptPreCreateInfoResponse**](GetReceiptPreCreateInfoResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptsMonthlyTotals

> GetReceiptsMonthlyTotalsResponse GetReceiptsMonthlyTotals(ctx, companyId).Type_(type_).Year(year).Execute()

Get Receipts Monthly Totals



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
    type_ := "type__example" // string | Receipt Type
    year := "year_example" // string | Year for which you want monthly totals

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptsApi.GetReceiptsMonthlyTotals(auth, companyId).Type_(type_).Year(year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.GetReceiptsMonthlyTotals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceiptsMonthlyTotals`: GetReceiptsMonthlyTotalsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptsMonthlyTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Receipt Type | 
 **year** | **string** | Year for which you want monthly totals | 

### Return type

[**GetReceiptsMonthlyTotalsResponse**](GetReceiptsMonthlyTotalsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReceipts

> ListReceiptsResponse ListReceipts(ctx, companyId).Fields(fields).Fieldset(fieldset).Page(page).PerPage(perPage).Sort(sort).Q(q).Execute()

List Receipts



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
    page := int32(56) // int32 | The page to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | The size of the page. (optional) (default to 5)
    sort := "sort_example" // string | List of comma-separated fields for result sorting (minus for desc sorting). (optional)
    q := "q_example" // string | Query for filtering the results. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptsApi.ListReceipts(auth, companyId).Fields(fields).Fieldset(fieldset).Page(page).PerPage(perPage).Sort(sort).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.ListReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReceipts`: ListReceiptsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListReceiptsResponse**](ListReceiptsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyReceipt

> ModifyReceiptResponse ModifyReceipt(ctx, companyId, documentId).ModifyReceiptRequest(modifyReceiptRequest).Execute()

Modify Receipt



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
    modifyReceiptRequest := *fattureincloud.NewModifyReceiptRequest() // ModifyReceiptRequest | Modified receipt. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceiptsApi.ModifyReceipt(auth, companyId, documentId).ModifyReceiptRequest(modifyReceiptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsApi.ModifyReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyReceipt`: ModifyReceiptResponse
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

Other parameters are passed through a pointer to a apiModifyReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyReceiptRequest** | [**ModifyReceiptRequest**](ModifyReceiptRequest.md) | Modified receipt. | 

### Return type

[**ModifyReceiptResponse**](ModifyReceiptResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

