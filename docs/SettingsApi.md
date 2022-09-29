# \SettingsApi

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentAccount**](SettingsApi.md#CreatePaymentAccount) | **Post** /c/{company_id}/settings/payment_accounts | Create Payment Account
[**CreatePaymentMethod**](SettingsApi.md#CreatePaymentMethod) | **Post** /c/{company_id}/settings/payment_methods | Create Payment Method
[**CreateVatType**](SettingsApi.md#CreateVatType) | **Post** /c/{company_id}/settings/vat_types | Create Vat Type
[**DeletePaymentAccount**](SettingsApi.md#DeletePaymentAccount) | **Delete** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Delete Payment Account
[**DeletePaymentMethod**](SettingsApi.md#DeletePaymentMethod) | **Delete** /c/{company_id}/settings/payment_methods/{payment_method_id} | Delete Payment Method
[**DeleteVatType**](SettingsApi.md#DeleteVatType) | **Delete** /c/{company_id}/settings/vat_types/{vat_type_id} | Delete Vat Type
[**GetPaymentAccount**](SettingsApi.md#GetPaymentAccount) | **Get** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Get Payment Account
[**GetPaymentMethod**](SettingsApi.md#GetPaymentMethod) | **Get** /c/{company_id}/settings/payment_methods/{payment_method_id} | Get Payment Method
[**GetVatType**](SettingsApi.md#GetVatType) | **Get** /c/{company_id}/settings/vat_types/{vat_type_id} | Get Vat Type
[**ModifyPaymentAccount**](SettingsApi.md#ModifyPaymentAccount) | **Put** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Modify Payment Account
[**ModifyPaymentMethod**](SettingsApi.md#ModifyPaymentMethod) | **Put** /c/{company_id}/settings/payment_methods/{payment_method_id} | Modify Payment Method
[**ModifyVatType**](SettingsApi.md#ModifyVatType) | **Put** /c/{company_id}/settings/vat_types/{vat_type_id} | Modify Vat Type



## CreatePaymentAccount

> CreatePaymentAccountResponse CreatePaymentAccount(ctx, companyId).CreatePaymentAccountRequest(createPaymentAccountRequest).Execute()

Create Payment Account



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
    createPaymentAccountRequest := *fattureincloud.NewCreatePaymentAccountRequest() // CreatePaymentAccountRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreatePaymentAccount(auth, companyId).CreatePaymentAccountRequest(createPaymentAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreatePaymentAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentAccount`: CreatePaymentAccountResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPaymentAccountRequest** | [**CreatePaymentAccountRequest**](CreatePaymentAccountRequest.md) |  | 

### Return type

[**CreatePaymentAccountResponse**](CreatePaymentAccountResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentMethod

> CreatePaymentMethodResponse CreatePaymentMethod(ctx, companyId).CreatePaymentMethodRequest(createPaymentMethodRequest).Execute()

Create Payment Method



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
    createPaymentMethodRequest := *fattureincloud.NewCreatePaymentMethodRequest() // CreatePaymentMethodRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreatePaymentMethod(auth, companyId).CreatePaymentMethodRequest(createPaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreatePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentMethod`: CreatePaymentMethodResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPaymentMethodRequest** | [**CreatePaymentMethodRequest**](CreatePaymentMethodRequest.md) |  | 

### Return type

[**CreatePaymentMethodResponse**](CreatePaymentMethodResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVatType

> CreateVatTypeResponse CreateVatType(ctx, companyId).CreateVatTypeRequest(createVatTypeRequest).Execute()

Create Vat Type



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
    createVatTypeRequest := *fattureincloud.NewCreateVatTypeRequest() // CreateVatTypeRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateVatType(auth, companyId).CreateVatTypeRequest(createVatTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateVatType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVatType`: CreateVatTypeResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVatTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVatTypeRequest** | [**CreateVatTypeRequest**](CreateVatTypeRequest.md) |  | 

### Return type

[**CreateVatTypeResponse**](CreateVatTypeResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentAccount

> DeletePaymentAccount(ctx, companyId, paymentAccountId).Execute()

Delete Payment Account



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
    paymentAccountId := int32(56) // int32 | The Referred Payment Account Id.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeletePaymentAccount(auth, companyId, paymentAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeletePaymentAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentAccountId** | **int32** | The Referred Payment Account Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentAccountRequest struct via the builder pattern


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


## DeletePaymentMethod

> DeletePaymentMethod(ctx, companyId, paymentMethodId).Execute()

Delete Payment Method



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
    paymentMethodId := int32(56) // int32 | The Referred Payment Method Id.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeletePaymentMethod(auth, companyId, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeletePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentMethodId** | **int32** | The Referred Payment Method Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentMethodRequest struct via the builder pattern


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


## DeleteVatType

> DeleteVatType(ctx, companyId, vatTypeId).Execute()

Delete Vat Type



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
    vatTypeId := int32(56) // int32 | The Referred Vat Type Id.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeleteVatType(auth, companyId, vatTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteVatType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**vatTypeId** | **int32** | The Referred Vat Type Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVatTypeRequest struct via the builder pattern


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


## GetPaymentAccount

> GetPaymentAccountResponse GetPaymentAccount(ctx, companyId, paymentAccountId).Fields(fields).Fieldset(fieldset).Execute()

Get Payment Account



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
    paymentAccountId := int32(56) // int32 | The Referred Payment Account Id.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetPaymentAccount(auth, companyId, paymentAccountId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetPaymentAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentAccount`: GetPaymentAccountResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentAccountId** | **int32** | The Referred Payment Account Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetPaymentAccountResponse**](GetPaymentAccountResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> GetPaymentMethodResponse GetPaymentMethod(ctx, companyId, paymentMethodId).Fields(fields).Fieldset(fieldset).Execute()

Get Payment Method



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
    paymentMethodId := int32(56) // int32 | The Referred Payment Method Id.
    fields := "fields_example" // string | List of comma-separated fields. (optional)
    fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetPaymentMethod(auth, companyId, paymentMethodId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethod`: GetPaymentMethodResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentMethodId** | **int32** | The Referred Payment Method Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetPaymentMethodResponse**](GetPaymentMethodResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVatType

> GetVatTypeResponse GetVatType(ctx, companyId, vatTypeId).Execute()

Get Vat Type



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
    vatTypeId := int32(56) // int32 | The Referred Vat Type Id.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.GetVatType(auth, companyId, vatTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetVatType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVatType`: GetVatTypeResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**vatTypeId** | **int32** | The Referred Vat Type Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVatTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVatTypeResponse**](GetVatTypeResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPaymentAccount

> ModifyPaymentAccountResponse ModifyPaymentAccount(ctx, companyId, paymentAccountId).ModifyPaymentAccountRequest(modifyPaymentAccountRequest).Execute()

Modify Payment Account



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
    paymentAccountId := int32(56) // int32 | The Referred Payment Account Id.
    modifyPaymentAccountRequest := *fattureincloud.NewModifyPaymentAccountRequest() // ModifyPaymentAccountRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ModifyPaymentAccount(auth, companyId, paymentAccountId).ModifyPaymentAccountRequest(modifyPaymentAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ModifyPaymentAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyPaymentAccount`: ModifyPaymentAccountResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentAccountId** | **int32** | The Referred Payment Account Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPaymentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyPaymentAccountRequest** | [**ModifyPaymentAccountRequest**](ModifyPaymentAccountRequest.md) |  | 

### Return type

[**ModifyPaymentAccountResponse**](ModifyPaymentAccountResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPaymentMethod

> ModifyPaymentMethodResponse ModifyPaymentMethod(ctx, companyId, paymentMethodId).ModifyPaymentMethodRequest(modifyPaymentMethodRequest).Execute()

Modify Payment Method



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
    paymentMethodId := int32(56) // int32 | The Referred Payment Method Id.
    modifyPaymentMethodRequest := *fattureincloud.NewModifyPaymentMethodRequest() // ModifyPaymentMethodRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ModifyPaymentMethod(auth, companyId, paymentMethodId).ModifyPaymentMethodRequest(modifyPaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ModifyPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyPaymentMethod`: ModifyPaymentMethodResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**paymentMethodId** | **int32** | The Referred Payment Method Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyPaymentMethodRequest** | [**ModifyPaymentMethodRequest**](ModifyPaymentMethodRequest.md) |  | 

### Return type

[**ModifyPaymentMethodResponse**](ModifyPaymentMethodResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyVatType

> ModifyVatTypeResponse ModifyVatType(ctx, companyId, vatTypeId).ModifyVatTypeRequest(modifyVatTypeRequest).Execute()

Modify Vat Type



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
    vatTypeId := int32(56) // int32 | The Referred Vat Type Id.
    modifyVatTypeRequest := *fattureincloud.NewModifyVatTypeRequest() // ModifyVatTypeRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ModifyVatType(auth, companyId, vatTypeId).ModifyVatTypeRequest(modifyVatTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ModifyVatType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyVatType`: ModifyVatTypeResponse
    json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 
**vatTypeId** | **int32** | The Referred Vat Type Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyVatTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyVatTypeRequest** | [**ModifyVatTypeRequest**](ModifyVatTypeRequest.md) |  | 

### Return type

[**ModifyVatTypeResponse**](ModifyVatTypeResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

