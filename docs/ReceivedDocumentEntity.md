# ReceivedDocumentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Entity unique identifier. | [optional] 
**Name** | Pointer to **NullableString** | Entity name. | [optional] 

## Methods

### NewReceivedDocumentEntity

`func NewReceivedDocumentEntity() *ReceivedDocumentEntity`

NewReceivedDocumentEntity instantiates a new ReceivedDocumentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentEntityWithDefaults

`func NewReceivedDocumentEntityWithDefaults() *ReceivedDocumentEntity`

NewReceivedDocumentEntityWithDefaults instantiates a new ReceivedDocumentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceivedDocumentEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivedDocumentEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivedDocumentEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivedDocumentEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceivedDocumentEntity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceivedDocumentEntity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ReceivedDocumentEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReceivedDocumentEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReceivedDocumentEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReceivedDocumentEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReceivedDocumentEntity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReceivedDocumentEntity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


