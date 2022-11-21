# IssuedDocumentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixPayments** | Pointer to **NullableBool** | Fixes your last payment amount to match your document total | [optional] 
**CreateFrom** | Pointer to **[]string** | Original documents ids [only for join/transform] | [optional] 
**Transform** | Pointer to **NullableBool** | Tranform a document. [only for transform] | [optional] 
**KeepCopy** | Pointer to **NullableBool** | Keep original document [only for transform] | [optional] 
**JoinType** | Pointer to **NullableString** | Join type [only for join] | [optional] 

## Methods

### NewIssuedDocumentOptions

`func NewIssuedDocumentOptions() *IssuedDocumentOptions`

NewIssuedDocumentOptions instantiates a new IssuedDocumentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentOptionsWithDefaults

`func NewIssuedDocumentOptionsWithDefaults() *IssuedDocumentOptions`

NewIssuedDocumentOptionsWithDefaults instantiates a new IssuedDocumentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixPayments

`func (o *IssuedDocumentOptions) GetFixPayments() bool`

GetFixPayments returns the FixPayments field if non-nil, zero value otherwise.

### GetFixPaymentsOk

`func (o *IssuedDocumentOptions) GetFixPaymentsOk() (*bool, bool)`

GetFixPaymentsOk returns a tuple with the FixPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixPayments

`func (o *IssuedDocumentOptions) SetFixPayments(v bool)`

SetFixPayments sets FixPayments field to given value.

### HasFixPayments

`func (o *IssuedDocumentOptions) HasFixPayments() bool`

HasFixPayments returns a boolean if a field has been set.

### SetFixPaymentsNil

`func (o *IssuedDocumentOptions) SetFixPaymentsNil(b bool)`

 SetFixPaymentsNil sets the value for FixPayments to be an explicit nil

### UnsetFixPayments
`func (o *IssuedDocumentOptions) UnsetFixPayments()`

UnsetFixPayments ensures that no value is present for FixPayments, not even an explicit nil
### GetCreateFrom

`func (o *IssuedDocumentOptions) GetCreateFrom() []string`

GetCreateFrom returns the CreateFrom field if non-nil, zero value otherwise.

### GetCreateFromOk

`func (o *IssuedDocumentOptions) GetCreateFromOk() (*[]string, bool)`

GetCreateFromOk returns a tuple with the CreateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFrom

`func (o *IssuedDocumentOptions) SetCreateFrom(v []string)`

SetCreateFrom sets CreateFrom field to given value.

### HasCreateFrom

`func (o *IssuedDocumentOptions) HasCreateFrom() bool`

HasCreateFrom returns a boolean if a field has been set.

### SetCreateFromNil

`func (o *IssuedDocumentOptions) SetCreateFromNil(b bool)`

 SetCreateFromNil sets the value for CreateFrom to be an explicit nil

### UnsetCreateFrom
`func (o *IssuedDocumentOptions) UnsetCreateFrom()`

UnsetCreateFrom ensures that no value is present for CreateFrom, not even an explicit nil
### GetTransform

`func (o *IssuedDocumentOptions) GetTransform() bool`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *IssuedDocumentOptions) GetTransformOk() (*bool, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *IssuedDocumentOptions) SetTransform(v bool)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *IssuedDocumentOptions) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### SetTransformNil

`func (o *IssuedDocumentOptions) SetTransformNil(b bool)`

 SetTransformNil sets the value for Transform to be an explicit nil

### UnsetTransform
`func (o *IssuedDocumentOptions) UnsetTransform()`

UnsetTransform ensures that no value is present for Transform, not even an explicit nil
### GetKeepCopy

`func (o *IssuedDocumentOptions) GetKeepCopy() bool`

GetKeepCopy returns the KeepCopy field if non-nil, zero value otherwise.

### GetKeepCopyOk

`func (o *IssuedDocumentOptions) GetKeepCopyOk() (*bool, bool)`

GetKeepCopyOk returns a tuple with the KeepCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCopy

`func (o *IssuedDocumentOptions) SetKeepCopy(v bool)`

SetKeepCopy sets KeepCopy field to given value.

### HasKeepCopy

`func (o *IssuedDocumentOptions) HasKeepCopy() bool`

HasKeepCopy returns a boolean if a field has been set.

### SetKeepCopyNil

`func (o *IssuedDocumentOptions) SetKeepCopyNil(b bool)`

 SetKeepCopyNil sets the value for KeepCopy to be an explicit nil

### UnsetKeepCopy
`func (o *IssuedDocumentOptions) UnsetKeepCopy()`

UnsetKeepCopy ensures that no value is present for KeepCopy, not even an explicit nil
### GetJoinType

`func (o *IssuedDocumentOptions) GetJoinType() string`

GetJoinType returns the JoinType field if non-nil, zero value otherwise.

### GetJoinTypeOk

`func (o *IssuedDocumentOptions) GetJoinTypeOk() (*string, bool)`

GetJoinTypeOk returns a tuple with the JoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinType

`func (o *IssuedDocumentOptions) SetJoinType(v string)`

SetJoinType sets JoinType field to given value.

### HasJoinType

`func (o *IssuedDocumentOptions) HasJoinType() bool`

HasJoinType returns a boolean if a field has been set.

### SetJoinTypeNil

`func (o *IssuedDocumentOptions) SetJoinTypeNil(b bool)`

 SetJoinTypeNil sets the value for JoinType to be an explicit nil

### UnsetJoinType
`func (o *IssuedDocumentOptions) UnsetJoinType()`

UnsetJoinType ensures that no value is present for JoinType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


