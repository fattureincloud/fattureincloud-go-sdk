# \UserAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserInfo**](UserAPI.md#GetUserInfo) | **Get** /user/info | Get User Info
[**ListUserCompanies**](UserAPI.md#ListUserCompanies) | **Get** /user/companies | List User Companies



## GetUserInfo

> GetUserInfoResponse GetUserInfo(ctx).Execute()

Get User Info



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

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetUserInfo(auth).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserInfo``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetUserInfo`: GetUserInfoResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**GetUserInfoResponse**](GetUserInfoResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserCompanies

> ListUserCompaniesResponse ListUserCompanies(ctx).Execute()

List User Companies



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

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ListUserCompanies(auth).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserCompanies``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `ListUserCompanies`: ListUserCompaniesResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserCompaniesRequest struct via the builder pattern


### Return type

[**ListUserCompaniesResponse**](ListUserCompaniesResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

