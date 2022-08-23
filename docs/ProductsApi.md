# \ProductsApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductsApi.md#CreateProduct) | **Post** /c/{company_id}/products | Create Product
[**DeleteProduct**](ProductsApi.md#DeleteProduct) | **Delete** /c/{company_id}/products/{product_id} | Delete Product
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /c/{company_id}/products/{product_id} | Get Product
[**ListProducts**](ProductsApi.md#ListProducts) | **Get** /c/{company_id}/products | List Products
[**ModifyProduct**](ProductsApi.md#ModifyProduct) | **Put** /c/{company_id}/products/{product_id} | Modify Product



## CreateProduct

> CreateProductResponse CreateProduct(ctx, companyId).CreateProductRequest(createProductRequest).Execute()

Create Product



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
    createProductRequest := *fattureincloud.NewCreateProductRequest() // CreateProductRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.CreateProduct(auth, companyId).CreateProductRequest(createProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: CreateProductResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) |  | 

### Return type

[**CreateProductResponse**](CreateProductResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProduct(ctx, companyId, productId).Execute()

Delete Product



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
    productId := int32(56) // int32 | The ID of the product.

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.DeleteProduct(auth, companyId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.DeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**productId** | **int32** | The ID of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## GetProduct

> GetProductResponse GetProduct(ctx, companyId, productId).Fields(fields).Fieldset(fieldset).Execute()

Get Product



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
    productId := int32(56) // int32 | The ID of the product.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.GetProduct(auth, companyId, productId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: GetProductResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**productId** | **int32** | The ID of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetProductResponse**](GetProductResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ListProductsResponse ListProducts(ctx, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()

List Products



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
    resp, r, err := apiClient.ProductsApi.ListProducts(auth, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ListProductsResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListProductsResponse**](ListProductsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyProduct

> ModifyProductResponse ModifyProduct(ctx, companyId, productId).ModifyProductRequest(modifyProductRequest).Execute()

Modify Product



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
    productId := int32(56) // int32 | The ID of the product.
    modifyProductRequest := *fattureincloud.NewModifyProductRequest() // ModifyProductRequest | Modified product details. (optional)

    auth := context.WithValue(context.Background(), fattureincloud.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.ModifyProduct(auth, companyId, productId).ModifyProductRequest(modifyProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ModifyProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyProduct`: ModifyProductResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**productId** | **int32** | The ID of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyProductRequest** | [**ModifyProductRequest**](ModifyProductRequest.md) | Modified product details. | 

### Return type

[**ModifyProductResponse**](ModifyProductResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

