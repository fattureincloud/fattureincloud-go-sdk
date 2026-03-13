# ListPendingReceivedDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **NullableInt32** | Current page number. | [optional] 
**FirstPageUrl** | Pointer to **NullableString** | First page url. | [optional] 
**From** | Pointer to **NullableInt32** | First result of the page. | [optional] 
**LastPage** | Pointer to **NullableInt32** | Last page number. | [optional] 
**LastPageUrl** | Pointer to **NullableString** | Last page url. | [optional] 
**NextPageUrl** | Pointer to **NullableString** | Next page url | [optional] 
**Path** | Pointer to **NullableString** | Request path. | [optional] 
**PerPage** | Pointer to **NullableInt32** | Number of result per page. | [optional] 
**PrevPageUrl** | Pointer to **NullableString** | Previous page url. | [optional] 
**To** | Pointer to **NullableInt32** | Last result of the page. | [optional] 
**Total** | Pointer to **NullableInt32** | Total number of results | [optional] 
**Data** | Pointer to [**[]PendingReceivedDocument**](PendingReceivedDocument.md) |  | [optional] 

## Methods

### NewListPendingReceivedDocumentsResponse

`func NewListPendingReceivedDocumentsResponse() *ListPendingReceivedDocumentsResponse`

NewListPendingReceivedDocumentsResponse instantiates a new ListPendingReceivedDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPendingReceivedDocumentsResponseWithDefaults

`func NewListPendingReceivedDocumentsResponseWithDefaults() *ListPendingReceivedDocumentsResponse`

NewListPendingReceivedDocumentsResponseWithDefaults instantiates a new ListPendingReceivedDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ListPendingReceivedDocumentsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListPendingReceivedDocumentsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListPendingReceivedDocumentsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListPendingReceivedDocumentsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *ListPendingReceivedDocumentsResponse) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *ListPendingReceivedDocumentsResponse) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetFirstPageUrl

`func (o *ListPendingReceivedDocumentsResponse) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *ListPendingReceivedDocumentsResponse) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *ListPendingReceivedDocumentsResponse) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.

### HasFirstPageUrl

`func (o *ListPendingReceivedDocumentsResponse) HasFirstPageUrl() bool`

HasFirstPageUrl returns a boolean if a field has been set.

### SetFirstPageUrlNil

`func (o *ListPendingReceivedDocumentsResponse) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *ListPendingReceivedDocumentsResponse) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *ListPendingReceivedDocumentsResponse) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListPendingReceivedDocumentsResponse) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListPendingReceivedDocumentsResponse) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListPendingReceivedDocumentsResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ListPendingReceivedDocumentsResponse) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListPendingReceivedDocumentsResponse) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *ListPendingReceivedDocumentsResponse) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListPendingReceivedDocumentsResponse) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListPendingReceivedDocumentsResponse) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ListPendingReceivedDocumentsResponse) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### SetLastPageNil

`func (o *ListPendingReceivedDocumentsResponse) SetLastPageNil(b bool)`

 SetLastPageNil sets the value for LastPage to be an explicit nil

### UnsetLastPage
`func (o *ListPendingReceivedDocumentsResponse) UnsetLastPage()`

UnsetLastPage ensures that no value is present for LastPage, not even an explicit nil
### GetLastPageUrl

`func (o *ListPendingReceivedDocumentsResponse) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *ListPendingReceivedDocumentsResponse) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *ListPendingReceivedDocumentsResponse) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.

### HasLastPageUrl

`func (o *ListPendingReceivedDocumentsResponse) HasLastPageUrl() bool`

HasLastPageUrl returns a boolean if a field has been set.

### SetLastPageUrlNil

`func (o *ListPendingReceivedDocumentsResponse) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *ListPendingReceivedDocumentsResponse) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetNextPageUrl

`func (o *ListPendingReceivedDocumentsResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListPendingReceivedDocumentsResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListPendingReceivedDocumentsResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *ListPendingReceivedDocumentsResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *ListPendingReceivedDocumentsResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListPendingReceivedDocumentsResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *ListPendingReceivedDocumentsResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListPendingReceivedDocumentsResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListPendingReceivedDocumentsResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListPendingReceivedDocumentsResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ListPendingReceivedDocumentsResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListPendingReceivedDocumentsResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListPendingReceivedDocumentsResponse) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListPendingReceivedDocumentsResponse) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListPendingReceivedDocumentsResponse) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListPendingReceivedDocumentsResponse) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### SetPerPageNil

`func (o *ListPendingReceivedDocumentsResponse) SetPerPageNil(b bool)`

 SetPerPageNil sets the value for PerPage to be an explicit nil

### UnsetPerPage
`func (o *ListPendingReceivedDocumentsResponse) UnsetPerPage()`

UnsetPerPage ensures that no value is present for PerPage, not even an explicit nil
### GetPrevPageUrl

`func (o *ListPendingReceivedDocumentsResponse) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListPendingReceivedDocumentsResponse) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListPendingReceivedDocumentsResponse) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *ListPendingReceivedDocumentsResponse) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### SetPrevPageUrlNil

`func (o *ListPendingReceivedDocumentsResponse) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListPendingReceivedDocumentsResponse) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *ListPendingReceivedDocumentsResponse) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListPendingReceivedDocumentsResponse) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListPendingReceivedDocumentsResponse) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ListPendingReceivedDocumentsResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ListPendingReceivedDocumentsResponse) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListPendingReceivedDocumentsResponse) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *ListPendingReceivedDocumentsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListPendingReceivedDocumentsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListPendingReceivedDocumentsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListPendingReceivedDocumentsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListPendingReceivedDocumentsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListPendingReceivedDocumentsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetData

`func (o *ListPendingReceivedDocumentsResponse) GetData() []PendingReceivedDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPendingReceivedDocumentsResponse) GetDataOk() (*[]PendingReceivedDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPendingReceivedDocumentsResponse) SetData(v []PendingReceivedDocument)`

SetData sets Data field to given value.

### HasData

`func (o *ListPendingReceivedDocumentsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ListPendingReceivedDocumentsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListPendingReceivedDocumentsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


