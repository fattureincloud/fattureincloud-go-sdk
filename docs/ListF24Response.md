# ListF24Response

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
**Data** | Pointer to [**[]F24**](F24.md) |  | [optional] 
**AggregatedData** | Pointer to [**ListF24ResponseAggregatedData**](ListF24ResponseAggregatedData.md) |  | [optional] 

## Methods

### NewListF24Response

`func NewListF24Response() *ListF24Response`

NewListF24Response instantiates a new ListF24Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListF24ResponseWithDefaults

`func NewListF24ResponseWithDefaults() *ListF24Response`

NewListF24ResponseWithDefaults instantiates a new ListF24Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ListF24Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListF24Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListF24Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListF24Response) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### SetCurrentPageNil

`func (o *ListF24Response) SetCurrentPageNil(b bool)`

 SetCurrentPageNil sets the value for CurrentPage to be an explicit nil

### UnsetCurrentPage
`func (o *ListF24Response) UnsetCurrentPage()`

UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil
### GetFirstPageUrl

`func (o *ListF24Response) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *ListF24Response) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *ListF24Response) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.

### HasFirstPageUrl

`func (o *ListF24Response) HasFirstPageUrl() bool`

HasFirstPageUrl returns a boolean if a field has been set.

### SetFirstPageUrlNil

`func (o *ListF24Response) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *ListF24Response) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *ListF24Response) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListF24Response) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListF24Response) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListF24Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ListF24Response) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListF24Response) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *ListF24Response) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListF24Response) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListF24Response) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ListF24Response) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### SetLastPageNil

`func (o *ListF24Response) SetLastPageNil(b bool)`

 SetLastPageNil sets the value for LastPage to be an explicit nil

### UnsetLastPage
`func (o *ListF24Response) UnsetLastPage()`

UnsetLastPage ensures that no value is present for LastPage, not even an explicit nil
### GetLastPageUrl

`func (o *ListF24Response) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *ListF24Response) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *ListF24Response) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.

### HasLastPageUrl

`func (o *ListF24Response) HasLastPageUrl() bool`

HasLastPageUrl returns a boolean if a field has been set.

### SetLastPageUrlNil

`func (o *ListF24Response) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *ListF24Response) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetNextPageUrl

`func (o *ListF24Response) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListF24Response) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListF24Response) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *ListF24Response) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *ListF24Response) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListF24Response) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *ListF24Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListF24Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListF24Response) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListF24Response) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ListF24Response) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListF24Response) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListF24Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListF24Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListF24Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListF24Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### SetPerPageNil

`func (o *ListF24Response) SetPerPageNil(b bool)`

 SetPerPageNil sets the value for PerPage to be an explicit nil

### UnsetPerPage
`func (o *ListF24Response) UnsetPerPage()`

UnsetPerPage ensures that no value is present for PerPage, not even an explicit nil
### GetPrevPageUrl

`func (o *ListF24Response) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListF24Response) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListF24Response) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *ListF24Response) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### SetPrevPageUrlNil

`func (o *ListF24Response) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListF24Response) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *ListF24Response) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListF24Response) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListF24Response) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ListF24Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ListF24Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListF24Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *ListF24Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListF24Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListF24Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListF24Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListF24Response) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListF24Response) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetData

`func (o *ListF24Response) GetData() []F24`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListF24Response) GetDataOk() (*[]F24, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListF24Response) SetData(v []F24)`

SetData sets Data field to given value.

### HasData

`func (o *ListF24Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ListF24Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListF24Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetAggregatedData

`func (o *ListF24Response) GetAggregatedData() ListF24ResponseAggregatedData`

GetAggregatedData returns the AggregatedData field if non-nil, zero value otherwise.

### GetAggregatedDataOk

`func (o *ListF24Response) GetAggregatedDataOk() (*ListF24ResponseAggregatedData, bool)`

GetAggregatedDataOk returns a tuple with the AggregatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedData

`func (o *ListF24Response) SetAggregatedData(v ListF24ResponseAggregatedData)`

SetAggregatedData sets AggregatedData field to given value.

### HasAggregatedData

`func (o *ListF24Response) HasAggregatedData() bool`

HasAggregatedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


