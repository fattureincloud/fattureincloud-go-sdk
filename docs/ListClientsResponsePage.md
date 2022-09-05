# ListClientsResponsePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Client**](Client.md) |  | [optional] 

## Methods

### NewListClientsResponsePage

`func NewListClientsResponsePage() *ListClientsResponsePage`

NewListClientsResponsePage instantiates a new ListClientsResponsePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClientsResponsePageWithDefaults

`func NewListClientsResponsePageWithDefaults() *ListClientsResponsePage`

NewListClientsResponsePageWithDefaults instantiates a new ListClientsResponsePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListClientsResponsePage) GetData() []Client`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListClientsResponsePage) GetDataOk() (*[]Client, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListClientsResponsePage) SetData(v []Client)`

SetData sets Data field to given value.

### HasData

`func (o *ListClientsResponsePage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ListClientsResponsePage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListClientsResponsePage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


