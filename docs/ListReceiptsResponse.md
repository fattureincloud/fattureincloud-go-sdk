# ListReceiptsResponse

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
**Data** | Pointer to [**[]Receipt**](Receipt.md) |  | [optional] 

## Methods

### NewListReceiptsResponse

`func NewListReceiptsResponse() *ListReceiptsResponse`

NewListReceiptsResponse instantiates a new ListReceiptsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReceiptsResponseWithDefaults

`func NewListReceiptsResponseWithDefaults() *ListReceiptsResponse`

NewListReceiptsResponseWithDefaults instantiates a new ListReceiptsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ListReceiptsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListReceiptsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListReceiptsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListReceiptsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *ListReceiptsResponse) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *ListReceiptsResponse) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetFirstPageUrl

`func (o *ListReceiptsResponse) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *ListReceiptsResponse) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *ListReceiptsResponse) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.

### HasFirstPageUrl

`func (o *ListReceiptsResponse) HasFirstPageUrl() bool`

HasFirstPageUrl returns a boolean if a field has been set.

### SetFirstPageUrlNil

`func (o *ListReceiptsResponse) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *ListReceiptsResponse) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *ListReceiptsResponse) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListReceiptsResponse) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListReceiptsResponse) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListReceiptsResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ListReceiptsResponse) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListReceiptsResponse) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *ListReceiptsResponse) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListReceiptsResponse) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListReceiptsResponse) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ListReceiptsResponse) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### SetLastPageNil

`func (o *ListReceiptsResponse) SetLastPageNil(b bool)`

 SetLastPageNil sets the value for LastPage to be an explicit nil

### UnsetLastPage
`func (o *ListReceiptsResponse) UnsetLastPage()`

UnsetLastPage ensures that no value is present for LastPage, not even an explicit nil
### GetLastPageUrl

`func (o *ListReceiptsResponse) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *ListReceiptsResponse) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *ListReceiptsResponse) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.

### HasLastPageUrl

`func (o *ListReceiptsResponse) HasLastPageUrl() bool`

HasLastPageUrl returns a boolean if a field has been set.

### SetLastPageUrlNil

`func (o *ListReceiptsResponse) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *ListReceiptsResponse) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetNextPageUrl

`func (o *ListReceiptsResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListReceiptsResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListReceiptsResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *ListReceiptsResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *ListReceiptsResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListReceiptsResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *ListReceiptsResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListReceiptsResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListReceiptsResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListReceiptsResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ListReceiptsResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListReceiptsResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListReceiptsResponse) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListReceiptsResponse) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListReceiptsResponse) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListReceiptsResponse) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### SetPerPageNil

`func (o *ListReceiptsResponse) SetPerPageNil(b bool)`

 SetPerPageNil sets the value for PerPage to be an explicit nil

### UnsetPerPage
`func (o *ListReceiptsResponse) UnsetPerPage()`

UnsetPerPage ensures that no value is present for PerPage, not even an explicit nil
### GetPrevPageUrl

`func (o *ListReceiptsResponse) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListReceiptsResponse) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListReceiptsResponse) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *ListReceiptsResponse) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### SetPrevPageUrlNil

`func (o *ListReceiptsResponse) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListReceiptsResponse) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *ListReceiptsResponse) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListReceiptsResponse) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListReceiptsResponse) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ListReceiptsResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ListReceiptsResponse) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListReceiptsResponse) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *ListReceiptsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListReceiptsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListReceiptsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListReceiptsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListReceiptsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListReceiptsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetData

`func (o *ListReceiptsResponse) GetData() []Receipt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListReceiptsResponse) GetDataOk() (*[]Receipt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListReceiptsResponse) SetData(v []Receipt)`

SetData sets Data field to given value.

### HasData

`func (o *ListReceiptsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ListReceiptsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListReceiptsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


