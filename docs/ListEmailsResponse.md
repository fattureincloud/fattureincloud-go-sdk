# ListEmailsResponse

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
**Data** | Pointer to [**[]Email**](Email.md) |  | [optional] 

## Methods

### NewListEmailsResponse

`func NewListEmailsResponse() *ListEmailsResponse`

NewListEmailsResponse instantiates a new ListEmailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEmailsResponseWithDefaults

`func NewListEmailsResponseWithDefaults() *ListEmailsResponse`

NewListEmailsResponseWithDefaults instantiates a new ListEmailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ListEmailsResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListEmailsResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListEmailsResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListEmailsResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *ListEmailsResponse) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *ListEmailsResponse) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetFirstPageUrl

`func (o *ListEmailsResponse) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *ListEmailsResponse) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *ListEmailsResponse) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.

### HasFirstPageUrl

`func (o *ListEmailsResponse) HasFirstPageUrl() bool`

HasFirstPageUrl returns a boolean if a field has been set.

### SetFirstPageUrlNil

`func (o *ListEmailsResponse) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *ListEmailsResponse) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *ListEmailsResponse) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListEmailsResponse) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListEmailsResponse) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListEmailsResponse) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ListEmailsResponse) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListEmailsResponse) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *ListEmailsResponse) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListEmailsResponse) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListEmailsResponse) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ListEmailsResponse) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### SetLastPageNil

`func (o *ListEmailsResponse) SetLastPageNil(b bool)`

 SetLastPageNil sets the value for LastPage to be an explicit nil

### UnsetLastPage
`func (o *ListEmailsResponse) UnsetLastPage()`

UnsetLastPage ensures that no value is present for LastPage, not even an explicit nil
### GetLastPageUrl

`func (o *ListEmailsResponse) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *ListEmailsResponse) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *ListEmailsResponse) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.

### HasLastPageUrl

`func (o *ListEmailsResponse) HasLastPageUrl() bool`

HasLastPageUrl returns a boolean if a field has been set.

### SetLastPageUrlNil

`func (o *ListEmailsResponse) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *ListEmailsResponse) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetNextPageUrl

`func (o *ListEmailsResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListEmailsResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListEmailsResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *ListEmailsResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *ListEmailsResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListEmailsResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *ListEmailsResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListEmailsResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListEmailsResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListEmailsResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ListEmailsResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListEmailsResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListEmailsResponse) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListEmailsResponse) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListEmailsResponse) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListEmailsResponse) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### SetPerPageNil

`func (o *ListEmailsResponse) SetPerPageNil(b bool)`

 SetPerPageNil sets the value for PerPage to be an explicit nil

### UnsetPerPage
`func (o *ListEmailsResponse) UnsetPerPage()`

UnsetPerPage ensures that no value is present for PerPage, not even an explicit nil
### GetPrevPageUrl

`func (o *ListEmailsResponse) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListEmailsResponse) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListEmailsResponse) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *ListEmailsResponse) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### SetPrevPageUrlNil

`func (o *ListEmailsResponse) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListEmailsResponse) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *ListEmailsResponse) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListEmailsResponse) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListEmailsResponse) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ListEmailsResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ListEmailsResponse) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListEmailsResponse) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *ListEmailsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListEmailsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListEmailsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListEmailsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListEmailsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListEmailsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetData

`func (o *ListEmailsResponse) GetData() []Email`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEmailsResponse) GetDataOk() (*[]Email, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEmailsResponse) SetData(v []Email)`

SetData sets Data field to given value.

### HasData

`func (o *ListEmailsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


