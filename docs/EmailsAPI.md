# \EmailsAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmails**](EmailsAPI.md#ListEmails) | **Get** /c/{company_id}/emails | List Emails



## ListEmails

> ListEmailsResponse ListEmails(ctx, companyId).Execute()

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

	auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
	configuration := fattureincloudapi.NewConfiguration()
	apiClient := fattureincloudapi.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.ListEmails(auth, companyId).Execute()
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

