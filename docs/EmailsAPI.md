# \EmailsAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmails**](EmailsAPI.md#ListEmails) | **Get** /c/{company_id}/emails | List Emails



## ListEmails

> ListEmailsResponse ListEmails(ctx, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()

List Emails



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
	fields := "fields_example" // string | List of comma-separated fields. (optional)
	fieldset := "fieldset_example" // string | Name of the fieldset. (optional)
	sort := "sort_example" // string | List of comma-separated fields for result sorting (minus for desc sorting). (optional)
	page := int32(56) // int32 | The page to retrieve. (optional) (default to 1)
	perPage := int32(56) // int32 | The size of the page. (optional) (default to 5)
	q := "q_example" // string | Query for filtering the results. (optional)

	auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
	configuration := fattureincloudapi.NewConfiguration()
	apiClient := fattureincloudapi.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.ListEmails(auth, companyId).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.ListEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmails`: ListEmailsResponse
	json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 

### Return type

[**ListEmailsResponse**](ListEmailsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

