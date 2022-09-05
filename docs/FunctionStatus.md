# FunctionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewFunctionStatus

`func NewFunctionStatus() *FunctionStatus`

NewFunctionStatus instantiates a new FunctionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStatusWithDefaults

`func NewFunctionStatusWithDefaults() *FunctionStatus`

NewFunctionStatusWithDefaults instantiates a new FunctionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *FunctionStatus) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FunctionStatus) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FunctionStatus) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FunctionStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *FunctionStatus) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *FunctionStatus) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


