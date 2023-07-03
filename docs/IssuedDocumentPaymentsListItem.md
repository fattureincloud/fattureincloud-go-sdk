# IssuedDocumentPaymentsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Issued document payment item id | [optional] 
**DueDate** | Pointer to **NullableString** | Issued document payment due date | [optional] 
**Amount** | Pointer to **NullableFloat32** | Issued document payment amount | [optional] 
**Status** | Pointer to [**IssuedDocumentStatus**](IssuedDocumentStatus.md) |  | [optional] [default to NOT_PAID]
**PaymentAccount** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 
**PaidDate** | Pointer to **NullableString** | Issued document payment date [Only if status is paid] | [optional] 
**EiRaw** | Pointer to **map[string]interface{}** | Issued document payment advanced raw attributes for e-invoices | [optional] 
**PaymentTerms** | Pointer to [**IssuedDocumentPaymentsListItemPaymentTerms**](IssuedDocumentPaymentsListItemPaymentTerms.md) |  | [optional] 

## Methods

### NewIssuedDocumentPaymentsListItem

`func NewIssuedDocumentPaymentsListItem() *IssuedDocumentPaymentsListItem`

NewIssuedDocumentPaymentsListItem instantiates a new IssuedDocumentPaymentsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentPaymentsListItemWithDefaults

`func NewIssuedDocumentPaymentsListItemWithDefaults() *IssuedDocumentPaymentsListItem`

NewIssuedDocumentPaymentsListItemWithDefaults instantiates a new IssuedDocumentPaymentsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuedDocumentPaymentsListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuedDocumentPaymentsListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuedDocumentPaymentsListItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IssuedDocumentPaymentsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IssuedDocumentPaymentsListItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IssuedDocumentPaymentsListItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDueDate

`func (o *IssuedDocumentPaymentsListItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *IssuedDocumentPaymentsListItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *IssuedDocumentPaymentsListItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *IssuedDocumentPaymentsListItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *IssuedDocumentPaymentsListItem) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *IssuedDocumentPaymentsListItem) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetAmount

`func (o *IssuedDocumentPaymentsListItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IssuedDocumentPaymentsListItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IssuedDocumentPaymentsListItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *IssuedDocumentPaymentsListItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *IssuedDocumentPaymentsListItem) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *IssuedDocumentPaymentsListItem) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetStatus

`func (o *IssuedDocumentPaymentsListItem) GetStatus() IssuedDocumentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IssuedDocumentPaymentsListItem) GetStatusOk() (*IssuedDocumentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IssuedDocumentPaymentsListItem) SetStatus(v IssuedDocumentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IssuedDocumentPaymentsListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *IssuedDocumentPaymentsListItem) GetPaymentAccount() PaymentAccount`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *IssuedDocumentPaymentsListItem) GetPaymentAccountOk() (*PaymentAccount, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *IssuedDocumentPaymentsListItem) SetPaymentAccount(v PaymentAccount)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *IssuedDocumentPaymentsListItem) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### SetPaymentAccountNil

`func (o *IssuedDocumentPaymentsListItem) SetPaymentAccountNil(b bool)`

 SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil

### UnsetPaymentAccount
`func (o *IssuedDocumentPaymentsListItem) UnsetPaymentAccount()`

UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil
### GetPaidDate

`func (o *IssuedDocumentPaymentsListItem) GetPaidDate() string`

GetPaidDate returns the PaidDate field if non-nil, zero value otherwise.

### GetPaidDateOk

`func (o *IssuedDocumentPaymentsListItem) GetPaidDateOk() (*string, bool)`

GetPaidDateOk returns a tuple with the PaidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidDate

`func (o *IssuedDocumentPaymentsListItem) SetPaidDate(v string)`

SetPaidDate sets PaidDate field to given value.

### HasPaidDate

`func (o *IssuedDocumentPaymentsListItem) HasPaidDate() bool`

HasPaidDate returns a boolean if a field has been set.

### SetPaidDateNil

`func (o *IssuedDocumentPaymentsListItem) SetPaidDateNil(b bool)`

 SetPaidDateNil sets the value for PaidDate to be an explicit nil

### UnsetPaidDate
`func (o *IssuedDocumentPaymentsListItem) UnsetPaidDate()`

UnsetPaidDate ensures that no value is present for PaidDate, not even an explicit nil
### GetEiRaw

`func (o *IssuedDocumentPaymentsListItem) GetEiRaw() map[string]interface{}`

GetEiRaw returns the EiRaw field if non-nil, zero value otherwise.

### GetEiRawOk

`func (o *IssuedDocumentPaymentsListItem) GetEiRawOk() (*map[string]interface{}, bool)`

GetEiRawOk returns a tuple with the EiRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiRaw

`func (o *IssuedDocumentPaymentsListItem) SetEiRaw(v map[string]interface{})`

SetEiRaw sets EiRaw field to given value.

### HasEiRaw

`func (o *IssuedDocumentPaymentsListItem) HasEiRaw() bool`

HasEiRaw returns a boolean if a field has been set.

### SetEiRawNil

`func (o *IssuedDocumentPaymentsListItem) SetEiRawNil(b bool)`

 SetEiRawNil sets the value for EiRaw to be an explicit nil

### UnsetEiRaw
`func (o *IssuedDocumentPaymentsListItem) UnsetEiRaw()`

UnsetEiRaw ensures that no value is present for EiRaw, not even an explicit nil
### GetPaymentTerms

`func (o *IssuedDocumentPaymentsListItem) GetPaymentTerms() IssuedDocumentPaymentsListItemPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *IssuedDocumentPaymentsListItem) GetPaymentTermsOk() (*IssuedDocumentPaymentsListItemPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *IssuedDocumentPaymentsListItem) SetPaymentTerms(v IssuedDocumentPaymentsListItemPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *IssuedDocumentPaymentsListItem) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


