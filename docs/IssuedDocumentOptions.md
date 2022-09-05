# IssuedDocumentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixPayments** | Pointer to **NullableBool** | Fixes your last payment amount to match your document total | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


