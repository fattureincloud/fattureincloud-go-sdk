# IssuedDocumentPaymentsListItemPaymentTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **NullableInt32** | The number of days by which the payment must be made. | [optional] 
**Type** | Pointer to **NullableString** | Payment terms type. | [optional] 

## Methods

### NewIssuedDocumentPaymentsListItemPaymentTerms

`func NewIssuedDocumentPaymentsListItemPaymentTerms() *IssuedDocumentPaymentsListItemPaymentTerms`

NewIssuedDocumentPaymentsListItemPaymentTerms instantiates a new IssuedDocumentPaymentsListItemPaymentTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentPaymentsListItemPaymentTermsWithDefaults

`func NewIssuedDocumentPaymentsListItemPaymentTermsWithDefaults() *IssuedDocumentPaymentsListItemPaymentTerms`

NewIssuedDocumentPaymentsListItemPaymentTermsWithDefaults instantiates a new IssuedDocumentPaymentsListItemPaymentTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *IssuedDocumentPaymentsListItemPaymentTerms) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetType

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *IssuedDocumentPaymentsListItemPaymentTerms) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IssuedDocumentPaymentsListItemPaymentTerms) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


