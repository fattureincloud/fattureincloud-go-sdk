# CreateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Client**](Client.md) |  | [optional] 

## Methods

### NewCreateClientRequest

`func NewCreateClientRequest() *CreateClientRequest`

NewCreateClientRequest instantiates a new CreateClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientRequestWithDefaults

`func NewCreateClientRequestWithDefaults() *CreateClientRequest`

NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateClientRequest) GetData() Client`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateClientRequest) GetDataOk() (*Client, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateClientRequest) SetData(v Client)`

SetData sets Data field to given value.

### HasData

`func (o *CreateClientRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


