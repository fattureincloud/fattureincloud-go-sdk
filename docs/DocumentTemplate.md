# DocumentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier. | [optional] 
**Name** | Pointer to **NullableString** | Template name. | [optional] 
**Type** | Pointer to **NullableString** | Template type. | [optional] 

## Methods

### NewDocumentTemplate

`func NewDocumentTemplate() *DocumentTemplate`

NewDocumentTemplate instantiates a new DocumentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTemplateWithDefaults

`func NewDocumentTemplateWithDefaults() *DocumentTemplate`

NewDocumentTemplateWithDefaults instantiates a new DocumentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DocumentTemplate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DocumentTemplate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DocumentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocumentTemplate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocumentTemplate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *DocumentTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DocumentTemplate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DocumentTemplate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


