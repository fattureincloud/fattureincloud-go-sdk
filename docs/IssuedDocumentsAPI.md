# \IssuedDocumentsAPI

All URIs are relative to *https://api-v2.fattureincloud.it*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssuedDocument**](IssuedDocumentsAPI.md#CreateIssuedDocument) | **Post** /c/{company_id}/issued_documents | Create Issued Document
[**DeleteIssuedDocument**](IssuedDocumentsAPI.md#DeleteIssuedDocument) | **Delete** /c/{company_id}/issued_documents/{document_id} | Delete Issued Document
[**DeleteIssuedDocumentAttachment**](IssuedDocumentsAPI.md#DeleteIssuedDocumentAttachment) | **Delete** /c/{company_id}/issued_documents/{document_id}/attachment | Delete Issued Document Attachment
[**GetEmailData**](IssuedDocumentsAPI.md#GetEmailData) | **Get** /c/{company_id}/issued_documents/{document_id}/email | Get Email Data
[**GetExistingIssuedDocumentTotals**](IssuedDocumentsAPI.md#GetExistingIssuedDocumentTotals) | **Post** /c/{company_id}/issued_documents/{document_id}/totals | Get Existing Issued Document Totals
[**GetIssuedDocument**](IssuedDocumentsAPI.md#GetIssuedDocument) | **Get** /c/{company_id}/issued_documents/{document_id} | Get Issued Document
[**GetIssuedDocumentPreCreateInfo**](IssuedDocumentsAPI.md#GetIssuedDocumentPreCreateInfo) | **Get** /c/{company_id}/issued_documents/info | Get Issued Document Pre-Create Info
[**GetNewIssuedDocumentTotals**](IssuedDocumentsAPI.md#GetNewIssuedDocumentTotals) | **Post** /c/{company_id}/issued_documents/totals | Get New Issued Document Totals
[**JoinIssuedDocuments**](IssuedDocumentsAPI.md#JoinIssuedDocuments) | **Get** /c/{company_id}/issued_documents/join | Join Issued Documents
[**ListIssuedDocuments**](IssuedDocumentsAPI.md#ListIssuedDocuments) | **Get** /c/{company_id}/issued_documents | List Issued Documents
[**ModifyIssuedDocument**](IssuedDocumentsAPI.md#ModifyIssuedDocument) | **Put** /c/{company_id}/issued_documents/{document_id} | Modify Issued Document
[**ScheduleEmail**](IssuedDocumentsAPI.md#ScheduleEmail) | **Post** /c/{company_id}/issued_documents/{document_id}/email | Schedule Email
[**TransformIssuedDocument**](IssuedDocumentsAPI.md#TransformIssuedDocument) | **Get** /c/{company_id}/issued_documents/transform | Transform Issued Document
[**UploadIssuedDocumentAttachment**](IssuedDocumentsAPI.md#UploadIssuedDocumentAttachment) | **Post** /c/{company_id}/issued_documents/attachment | Upload Issued Document Attachment



## CreateIssuedDocument

> CreateIssuedDocumentResponse CreateIssuedDocument(ctx, companyId).CreateIssuedDocumentRequest(createIssuedDocumentRequest).Execute()

Create Issued Document



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
        createIssuedDocumentRequest := *fattureincloud.NewCreateIssuedDocumentRequest() // CreateIssuedDocumentRequest | The Issued Document (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.CreateIssuedDocument(auth, companyId).CreateIssuedDocumentRequest(createIssuedDocumentRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.CreateIssuedDocument``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `CreateIssuedDocument`: CreateIssuedDocumentResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssuedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIssuedDocumentRequest** | [**CreateIssuedDocumentRequest**](CreateIssuedDocumentRequest.md) | The Issued Document | 

### Return type

[**CreateIssuedDocumentResponse**](CreateIssuedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssuedDocument

> DeleteIssuedDocument(ctx, companyId, documentId).Execute()

Delete Issued Document



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
        documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.DeleteIssuedDocument(auth, companyId, documentId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.DeleteIssuedDocument``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIssuedDocumentRequest struct via the builder pattern


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


## DeleteIssuedDocumentAttachment

> DeleteIssuedDocumentAttachment(ctx, companyId, documentId).Execute()

Delete Issued Document Attachment



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
        documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.DeleteIssuedDocumentAttachment(auth, companyId, documentId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.DeleteIssuedDocumentAttachment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIssuedDocumentAttachmentRequest struct via the builder pattern


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


## GetEmailData

> GetEmailDataResponse GetEmailData(ctx, companyId, documentId).Execute()

Get Email Data



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
        documentId := int32(56) // int32 | The ID of the document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.GetEmailData(auth, companyId, documentId).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.GetEmailData``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetEmailData`: GetEmailDataResponse
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

Other parameters are passed through a pointer to a apiGetEmailDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEmailDataResponse**](GetEmailDataResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExistingIssuedDocumentTotals

> GetExistingIssuedDocumentTotalsResponse GetExistingIssuedDocumentTotals(ctx, companyId, documentId).GetExistingIssuedDocumentTotalsRequest(getExistingIssuedDocumentTotalsRequest).Execute()

Get Existing Issued Document Totals



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
        documentId := int32(56) // int32 | The ID of the document.
        getExistingIssuedDocumentTotalsRequest := *fattureincloud.NewGetExistingIssuedDocumentTotalsRequest() // GetExistingIssuedDocumentTotalsRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.GetExistingIssuedDocumentTotals(auth, companyId, documentId).GetExistingIssuedDocumentTotalsRequest(getExistingIssuedDocumentTotalsRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.GetExistingIssuedDocumentTotals``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetExistingIssuedDocumentTotals`: GetExistingIssuedDocumentTotalsResponse
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

Other parameters are passed through a pointer to a apiGetExistingIssuedDocumentTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getExistingIssuedDocumentTotalsRequest** | [**GetExistingIssuedDocumentTotalsRequest**](GetExistingIssuedDocumentTotalsRequest.md) |  | 

### Return type

[**GetExistingIssuedDocumentTotalsResponse**](GetExistingIssuedDocumentTotalsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuedDocument

> GetIssuedDocumentResponse GetIssuedDocument(ctx, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()

Get Issued Document



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
        documentId := int32(56) // int32 | The ID of the document.
        fields := "fields_example" // string | List of comma-separated fields. (optional)
        fieldset := "fieldset_example" // string | Name of the fieldset. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.GetIssuedDocument(auth, companyId, documentId).Fields(fields).Fieldset(fieldset).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.GetIssuedDocument``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetIssuedDocument`: GetIssuedDocumentResponse
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

Other parameters are passed through a pointer to a apiGetIssuedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 

### Return type

[**GetIssuedDocumentResponse**](GetIssuedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuedDocumentPreCreateInfo

> GetIssuedDocumentPreCreateInfoResponse GetIssuedDocumentPreCreateInfo(ctx, companyId).Type_(type_).Execute()

Get Issued Document Pre-Create Info



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
        type_ := "type__example" // string | The type of the issued document.

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.GetIssuedDocumentPreCreateInfo(auth, companyId).Type_(type_).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.GetIssuedDocumentPreCreateInfo``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetIssuedDocumentPreCreateInfo`: GetIssuedDocumentPreCreateInfoResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuedDocumentPreCreateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of the issued document. | 

### Return type

[**GetIssuedDocumentPreCreateInfoResponse**](GetIssuedDocumentPreCreateInfoResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewIssuedDocumentTotals

> GetNewIssuedDocumentTotalsResponse GetNewIssuedDocumentTotals(ctx, companyId).GetNewIssuedDocumentTotalsRequest(getNewIssuedDocumentTotalsRequest).Execute()

Get New Issued Document Totals



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
        getNewIssuedDocumentTotalsRequest := *fattureincloud.NewGetNewIssuedDocumentTotalsRequest() // GetNewIssuedDocumentTotalsRequest |  (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.GetNewIssuedDocumentTotals(auth, companyId).GetNewIssuedDocumentTotalsRequest(getNewIssuedDocumentTotalsRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.GetNewIssuedDocumentTotals``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `GetNewIssuedDocumentTotals`: GetNewIssuedDocumentTotalsResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewIssuedDocumentTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getNewIssuedDocumentTotalsRequest** | [**GetNewIssuedDocumentTotalsRequest**](GetNewIssuedDocumentTotalsRequest.md) |  | 

### Return type

[**GetNewIssuedDocumentTotalsResponse**](GetNewIssuedDocumentTotalsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinIssuedDocuments

> JoinIssuedDocumentsResponse JoinIssuedDocuments(ctx, companyId).Ids(ids).Group(group).Parameter(parameter).Execute()

Join Issued Documents



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
        ids := "1,2,3,4" // string | Ids of the documents.
        group := int32(56) // int32 | Group items. (optional)
        parameter := "delivery_notes, orders, quotes, work_reports" // string | Type of the parameters to be joined (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.JoinIssuedDocuments(auth, companyId).Ids(ids).Group(group).Parameter(parameter).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.JoinIssuedDocuments``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `JoinIssuedDocuments`: JoinIssuedDocumentsResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinIssuedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | Ids of the documents. | 
 **group** | **int32** | Group items. | 
 **parameter** | **string** | Type of the parameters to be joined | 

### Return type

[**JoinIssuedDocumentsResponse**](JoinIssuedDocumentsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIssuedDocuments

> ListIssuedDocumentsResponse ListIssuedDocuments(ctx, companyId).Type_(type_).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Inclusive(inclusive).Execute()

List Issued Documents



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
        type_ := "type__example" // string | The type of the issued document.
        fields := "fields_example" // string | List of comma-separated fields. (optional)
        fieldset := "fieldset_example" // string | Name of the fieldset. (optional)
        sort := "sort_example" // string | List of comma-separated fields for result sorting (minus for desc sorting). (optional)
        page := int32(56) // int32 | The page to retrieve. (optional) (default to 1)
        perPage := int32(56) // int32 | The size of the page. (optional) (default to 5)
        q := "q_example" // string | Query for filtering the results. (optional)
        inclusive := int32(56) // int32 | (Only for type = delivery_notes) Include invoices delivery notes. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.ListIssuedDocuments(auth, companyId).Type_(type_).Fields(fields).Fieldset(fieldset).Sort(sort).Page(page).PerPage(perPage).Q(q).Inclusive(inclusive).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.ListIssuedDocuments``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `ListIssuedDocuments`: ListIssuedDocumentsResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIssuedDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of the issued document. | 
 **fields** | **string** | List of comma-separated fields. | 
 **fieldset** | **string** | Name of the fieldset. | 
 **sort** | **string** | List of comma-separated fields for result sorting (minus for desc sorting). | 
 **page** | **int32** | The page to retrieve. | [default to 1]
 **perPage** | **int32** | The size of the page. | [default to 5]
 **q** | **string** | Query for filtering the results. | 
 **inclusive** | **int32** | (Only for type &#x3D; delivery_notes) Include invoices delivery notes. | 

### Return type

[**ListIssuedDocumentsResponse**](ListIssuedDocumentsResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIssuedDocument

> ModifyIssuedDocumentResponse ModifyIssuedDocument(ctx, companyId, documentId).ModifyIssuedDocumentRequest(modifyIssuedDocumentRequest).Execute()

Modify Issued Document



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
        documentId := int32(56) // int32 | The ID of the document.
        modifyIssuedDocumentRequest := *fattureincloud.NewModifyIssuedDocumentRequest() // ModifyIssuedDocumentRequest | The modified document (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.ModifyIssuedDocument(auth, companyId, documentId).ModifyIssuedDocumentRequest(modifyIssuedDocumentRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.ModifyIssuedDocument``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `ModifyIssuedDocument`: ModifyIssuedDocumentResponse
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

Other parameters are passed through a pointer to a apiModifyIssuedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyIssuedDocumentRequest** | [**ModifyIssuedDocumentRequest**](ModifyIssuedDocumentRequest.md) | The modified document | 

### Return type

[**ModifyIssuedDocumentResponse**](ModifyIssuedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEmail

> ScheduleEmail(ctx, companyId, documentId).ScheduleEmailRequest(scheduleEmailRequest).Execute()

Schedule Email



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
        documentId := int32(56) // int32 | The ID of the document.
        scheduleEmailRequest := *fattureincloud.NewScheduleEmailRequest() // ScheduleEmailRequest | Email Schedule (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.ScheduleEmail(auth, companyId, documentId).ScheduleEmailRequest(scheduleEmailRequest).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.ScheduleEmail``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScheduleEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scheduleEmailRequest** | [**ScheduleEmailRequest**](ScheduleEmailRequest.md) | Email Schedule | 

### Return type

 (empty response body)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransformIssuedDocument

> TransformIssuedDocumentResponse TransformIssuedDocument(ctx, companyId).OriginalDocumentId(originalDocumentId).NewType(newType).Parameter(parameter).EInvoice(eInvoice).TransformKeepCopy(transformKeepCopy).Execute()

Transform Issued Document



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
        originalDocumentId := int32(56) // int32 | Original document id.
        newType := "newType_example" // string | New document type.
        parameter := "parameter_example" // string | Old document type. (optional)
        eInvoice := int32(56) // int32 | New document e_invoice. (optional)
        transformKeepCopy := int32(56) // int32 | Keep the old document. (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.TransformIssuedDocument(auth, companyId).OriginalDocumentId(originalDocumentId).NewType(newType).Parameter(parameter).EInvoice(eInvoice).TransformKeepCopy(transformKeepCopy).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.TransformIssuedDocument``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `TransformIssuedDocument`: TransformIssuedDocumentResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransformIssuedDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **originalDocumentId** | **int32** | Original document id. | 
 **newType** | **string** | New document type. | 
 **parameter** | **string** | Old document type. | 
 **eInvoice** | **int32** | New document e_invoice. | 
 **transformKeepCopy** | **int32** | Keep the old document. | 

### Return type

[**TransformIssuedDocumentResponse**](TransformIssuedDocumentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIssuedDocumentAttachment

> UploadIssuedDocumentAttachmentResponse UploadIssuedDocumentAttachment(ctx, companyId).Filename(filename).Attachment(attachment).Execute()

Upload Issued Document Attachment



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
        filename := "filename_example" // string | Attachment file name (optional)
        attachment := os.NewFile(1234, "some_file") // *os.File | Attachment file [.png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx] (optional)

    auth := context.WithValue(context.Background(), fattureincloudapi.ContextAccessToken, "ACCESS_TOKEN")
    configuration := fattureincloudapi.NewConfiguration()
    apiClient := fattureincloudapi.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuedDocumentsAPI.UploadIssuedDocumentAttachment(auth, companyId).Filename(filename).Attachment(attachment).Execute()
    if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `IssuedDocumentsAPI.UploadIssuedDocumentAttachment``: %v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        // response from `UploadIssuedDocumentAttachment`: UploadIssuedDocumentAttachmentResponse
        json.NewEncoder(os.Stdout).Encode(resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** | The ID of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIssuedDocumentAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** | Attachment file name | 
 **attachment** | ***os.File** | Attachment file [.png, .jpg, .gif, .pdf, .zip, .xls, .xlsx, .doc, .docx] | 

### Return type

[**UploadIssuedDocumentAttachmentResponse**](UploadIssuedDocumentAttachmentResponse.md)

### Authorization

[OAuth2AuthenticationCodeFlow](../README.md#OAuth2AuthenticationCodeFlow)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

