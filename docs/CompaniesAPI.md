# \CompaniesAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyInfo**](CompaniesAPI.md#GetCompanyInfo) | **Get** /c/{company_id}/company/info | Get Company Info
[**GetCompanyPlanUsage**](CompaniesAPI.md#GetCompanyPlanUsage) | **Get** /c/{company_id}/company/plan_usage | Get Company Plan Usage



## GetCompanyInfo

> GetCompanyInfoResponse GetCompanyInfo(ctx, companyId).Execute()

Get Company Info



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
	resp, r, err := apiClient.CompaniesAPI.GetCompanyInfo(auth, companyId).Execute()
	if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompanyInfo``: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
		// response from `GetCompanyInfo`: GetCompanyInfoResponse
		json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCompanyInfoResponse**](GetCompanyInfoResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPlanUsage

> GetCompanyPlanUsageResponse GetCompanyPlanUsage(ctx, companyId).Category(category).Execute()

Get Company Plan Usage



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
		category := "category_example" // string | Category

	auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
	configuration := fattureincloudapi.NewConfiguration()
	apiClient := fattureincloudapi.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.GetCompanyPlanUsage(auth, companyId).Category(category).Execute()
	if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompanyPlanUsage``: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
		// response from `GetCompanyPlanUsage`: GetCompanyPlanUsageResponse
		json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPlanUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Category | 

### Return type

[**GetCompanyPlanUsageResponse**](GetCompanyPlanUsageResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

