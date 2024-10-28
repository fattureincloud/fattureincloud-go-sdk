# \WebhooksAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhooksSubscription**](WebhooksAPI.md#CreateWebhooksSubscription) | **Post** /c/{company_id}/subscriptions | Create a Webhook Subscription
[**DeleteWebhooksSubscription**](WebhooksAPI.md#DeleteWebhooksSubscription) | **Delete** /c/{company_id}/subscriptions/{subscription_id} | Delete Webhooks Subscription
[**GetWebhooksSubscription**](WebhooksAPI.md#GetWebhooksSubscription) | **Get** /c/{company_id}/subscriptions/{subscription_id} | Get Webhooks Subscription
[**ListWebhooksSubscriptions**](WebhooksAPI.md#ListWebhooksSubscriptions) | **Get** /c/{company_id}/subscriptions | List Webhooks Subscriptions
[**ModifyWebhooksSubscription**](WebhooksAPI.md#ModifyWebhooksSubscription) | **Put** /c/{company_id}/subscriptions/{subscription_id} | Modify Webhooks Subscription



## CreateWebhooksSubscription

> CreateWebhooksSubscriptionResponse CreateWebhooksSubscription(ctx, companyId).CreateWebhooksSubscriptionRequest(createWebhooksSubscriptionRequest).Execute()

Create a Webhook Subscription



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
        createWebhooksSubscriptionRequest := *fattureincloud.NewCreateWebhooksSubscriptionRequest() // CreateWebhooksSubscriptionRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.CreateWebhooksSubscription(auth, companyId).CreateWebhooksSubscriptionRequest(createWebhooksSubscriptionRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhooksSubscription``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `CreateWebhooksSubscription`: CreateWebhooksSubscriptionResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhooksSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWebhooksSubscriptionRequest** | [**CreateWebhooksSubscriptionRequest**](CreateWebhooksSubscriptionRequest.md) |  | 

### Return type

[**CreateWebhooksSubscriptionResponse**](CreateWebhooksSubscriptionResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhooksSubscription

> DeleteWebhooksSubscription(ctx, companyId, subscriptionId).Execute()

Delete Webhooks Subscription



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
        subscriptionId := "SUB123" // string | The ID of the subscription.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.DeleteWebhooksSubscription(auth, companyId, subscriptionId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhooksSubscription``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**subscriptionId** | **string** | The ID of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksSubscriptionRequest struct via the builder pattern


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


## GetWebhooksSubscription

> GetWebhooksSubscriptionResponse GetWebhooksSubscription(ctx, companyId, subscriptionId).Execute()

Get Webhooks Subscription



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
        subscriptionId := "SUB123" // string | The ID of the subscription.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.GetWebhooksSubscription(auth, companyId, subscriptionId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooksSubscription``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetWebhooksSubscription`: GetWebhooksSubscriptionResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**subscriptionId** | **string** | The ID of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetWebhooksSubscriptionResponse**](GetWebhooksSubscriptionResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooksSubscriptions

> ListWebhooksSubscriptionsResponse ListWebhooksSubscriptions(ctx, companyId).Execute()

List Webhooks Subscriptions



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
    resp, r, err := apiClient.WebhooksAPI.ListWebhooksSubscriptions(auth, companyId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooksSubscriptions``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `ListWebhooksSubscriptions`: ListWebhooksSubscriptionsResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWebhooksSubscriptionsResponse**](ListWebhooksSubscriptionsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWebhooksSubscription

> ModifyWebhooksSubscriptionResponse ModifyWebhooksSubscription(ctx, companyId, subscriptionId).ModifyWebhooksSubscriptionRequest(modifyWebhooksSubscriptionRequest).Execute()

Modify Webhooks Subscription



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
        subscriptionId := "SUB123" // string | The ID of the subscription.
        modifyWebhooksSubscriptionRequest := *fattureincloud.NewModifyWebhooksSubscriptionRequest() // ModifyWebhooksSubscriptionRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.ModifyWebhooksSubscription(auth, companyId, subscriptionId).ModifyWebhooksSubscriptionRequest(modifyWebhooksSubscriptionRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ModifyWebhooksSubscription``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `ModifyWebhooksSubscription`: ModifyWebhooksSubscriptionResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**subscriptionId** | **string** | The ID of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWebhooksSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyWebhooksSubscriptionRequest** | [**ModifyWebhooksSubscriptionRequest**](ModifyWebhooksSubscriptionRequest.md) |  | 

### Return type

[**ModifyWebhooksSubscriptionResponse**](ModifyWebhooksSubscriptionResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

