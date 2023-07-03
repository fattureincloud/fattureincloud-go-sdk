# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Language code | [optional] 
**Name** | Pointer to **NullableString** | Language extended name | [optional] 

## Methods

### NewLanguage

`func NewLanguage() *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Language) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Language) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Language) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Language) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Language) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Language) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *Language) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Language) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Language) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Language) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Language) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Language) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


