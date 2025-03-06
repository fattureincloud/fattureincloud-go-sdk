# \PriceListsAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriceListItems**](PriceListsAPI.md#GetPriceListItems) | **Get** /c/{company_id}/price_lists/{price_list_id}/items | Get PriceList Items List
[**GetPriceLists**](PriceListsAPI.md#GetPriceLists) | **Get** /c/{company_id}/price_lists | Get PriceLists



## GetPriceListItems

> GetPriceListItemsResponse GetPriceListItems(ctx, companyId, priceListId).Execute()

Get PriceList Items List



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
	priceListId := "priceListId_example" // string | 

	auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
	configuration := fattureincloudapi.NewConfiguration()
	apiClient := fattureincloudapi.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListItems(auth, companyId, priceListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListItems`: GetPriceListItemsResponse
	json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPriceListItemsResponse**](GetPriceListItemsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceLists

> ListPriceListsResponse GetPriceLists(ctx, companyId).Execute()

Get PriceLists



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

	auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
	configuration := fattureincloudapi.NewConfiguration()
	apiClient := fattureincloudapi.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceLists(auth, companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceLists`: ListPriceListsResponse
	json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPriceListsResponse**](ListPriceListsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

