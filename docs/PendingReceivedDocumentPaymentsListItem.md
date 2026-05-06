# PendingReceivedDocumentPaymentsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat32** | Pending received document payment total amount | [optional] 
**DueDate** | Pointer to **NullableString** | Due date | [optional] 
**PaidDate** | Pointer to **NullableString** | Pending received document payment paid date | [optional] 
**PaymentTerms** | Pointer to [**NullablePendingReceivedDocumentPaymentsListItemPaymentTerms**](PendingReceivedDocumentPaymentsListItemPaymentTerms.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Pending received document payment status | [optional] 
**PaidWithTsPay** | Pointer to **NullableBool** | True if paid with TS Pay | [optional] 
**PaymentAccount** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 

## Methods

### NewPendingReceivedDocumentPaymentsListItem

`func NewPendingReceivedDocumentPaymentsListItem() *PendingReceivedDocumentPaymentsListItem`

NewPendingReceivedDocumentPaymentsListItem instantiates a new PendingReceivedDocumentPaymentsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingReceivedDocumentPaymentsListItemWithDefaults

`func NewPendingReceivedDocumentPaymentsListItemWithDefaults() *PendingReceivedDocumentPaymentsListItem`

NewPendingReceivedDocumentPaymentsListItemWithDefaults instantiates a new PendingReceivedDocumentPaymentsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PendingReceivedDocumentPaymentsListItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PendingReceivedDocumentPaymentsListItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PendingReceivedDocumentPaymentsListItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetDueDate

`func (o *PendingReceivedDocumentPaymentsListItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PendingReceivedDocumentPaymentsListItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *PendingReceivedDocumentPaymentsListItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaidDate

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaidDate() string`

GetPaidDate returns the PaidDate field if non-nil, zero value otherwise.

### GetPaidDateOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaidDateOk() (*string, bool)`

GetPaidDateOk returns a tuple with the PaidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidDate

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaidDate(v string)`

SetPaidDate sets PaidDate field to given value.

### HasPaidDate

`func (o *PendingReceivedDocumentPaymentsListItem) HasPaidDate() bool`

HasPaidDate returns a boolean if a field has been set.

### SetPaidDateNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaidDateNil(b bool)`

 SetPaidDateNil sets the value for PaidDate to be an explicit nil

### UnsetPaidDate
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetPaidDate()`

UnsetPaidDate ensures that no value is present for PaidDate, not even an explicit nil
### GetPaymentTerms

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaymentTerms() PendingReceivedDocumentPaymentsListItemPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaymentTermsOk() (*PendingReceivedDocumentPaymentsListItemPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaymentTerms(v PendingReceivedDocumentPaymentsListItemPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *PendingReceivedDocumentPaymentsListItem) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetStatus

`func (o *PendingReceivedDocumentPaymentsListItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PendingReceivedDocumentPaymentsListItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PendingReceivedDocumentPaymentsListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPaidWithTsPay

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaidWithTsPay() bool`

GetPaidWithTsPay returns the PaidWithTsPay field if non-nil, zero value otherwise.

### GetPaidWithTsPayOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaidWithTsPayOk() (*bool, bool)`

GetPaidWithTsPayOk returns a tuple with the PaidWithTsPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidWithTsPay

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaidWithTsPay(v bool)`

SetPaidWithTsPay sets PaidWithTsPay field to given value.

### HasPaidWithTsPay

`func (o *PendingReceivedDocumentPaymentsListItem) HasPaidWithTsPay() bool`

HasPaidWithTsPay returns a boolean if a field has been set.

### SetPaidWithTsPayNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaidWithTsPayNil(b bool)`

 SetPaidWithTsPayNil sets the value for PaidWithTsPay to be an explicit nil

### UnsetPaidWithTsPay
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetPaidWithTsPay()`

UnsetPaidWithTsPay ensures that no value is present for PaidWithTsPay, not even an explicit nil
### GetPaymentAccount

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaymentAccount() PaymentAccount`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *PendingReceivedDocumentPaymentsListItem) GetPaymentAccountOk() (*PaymentAccount, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaymentAccount(v PaymentAccount)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *PendingReceivedDocumentPaymentsListItem) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### SetPaymentAccountNil

`func (o *PendingReceivedDocumentPaymentsListItem) SetPaymentAccountNil(b bool)`

 SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil

### UnsetPaymentAccount
`func (o *PendingReceivedDocumentPaymentsListItem) UnsetPaymentAccount()`

UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


