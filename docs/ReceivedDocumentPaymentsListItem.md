# ReceivedDocumentPaymentsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier. | [optional] 
**Amount** | Pointer to **NullableFloat32** | Amount of items. | [optional] 
**DueDate** | Pointer to **NullableString** | Due date | [optional] 
**PaidDate** | Pointer to **NullableString** | Paid date | [optional] 
**PaymentTerms** | Pointer to [**NullableReceivedDocumentPaymentsListItemPaymentTerms**](ReceivedDocumentPaymentsListItemPaymentTerms.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Payment status. | [optional] 
**PaymentAccount** | Pointer to [**NullablePaymentAccount**](PaymentAccount.md) |  | [optional] 

## Methods

### NewReceivedDocumentPaymentsListItem

`func NewReceivedDocumentPaymentsListItem() *ReceivedDocumentPaymentsListItem`

NewReceivedDocumentPaymentsListItem instantiates a new ReceivedDocumentPaymentsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentPaymentsListItemWithDefaults

`func NewReceivedDocumentPaymentsListItemWithDefaults() *ReceivedDocumentPaymentsListItem`

NewReceivedDocumentPaymentsListItemWithDefaults instantiates a new ReceivedDocumentPaymentsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReceivedDocumentPaymentsListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivedDocumentPaymentsListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivedDocumentPaymentsListItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivedDocumentPaymentsListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReceivedDocumentPaymentsListItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReceivedDocumentPaymentsListItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAmount

`func (o *ReceivedDocumentPaymentsListItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReceivedDocumentPaymentsListItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReceivedDocumentPaymentsListItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ReceivedDocumentPaymentsListItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ReceivedDocumentPaymentsListItem) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ReceivedDocumentPaymentsListItem) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetDueDate

`func (o *ReceivedDocumentPaymentsListItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ReceivedDocumentPaymentsListItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ReceivedDocumentPaymentsListItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ReceivedDocumentPaymentsListItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *ReceivedDocumentPaymentsListItem) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *ReceivedDocumentPaymentsListItem) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaidDate

`func (o *ReceivedDocumentPaymentsListItem) GetPaidDate() string`

GetPaidDate returns the PaidDate field if non-nil, zero value otherwise.

### GetPaidDateOk

`func (o *ReceivedDocumentPaymentsListItem) GetPaidDateOk() (*string, bool)`

GetPaidDateOk returns a tuple with the PaidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidDate

`func (o *ReceivedDocumentPaymentsListItem) SetPaidDate(v string)`

SetPaidDate sets PaidDate field to given value.

### HasPaidDate

`func (o *ReceivedDocumentPaymentsListItem) HasPaidDate() bool`

HasPaidDate returns a boolean if a field has been set.

### SetPaidDateNil

`func (o *ReceivedDocumentPaymentsListItem) SetPaidDateNil(b bool)`

 SetPaidDateNil sets the value for PaidDate to be an explicit nil

### UnsetPaidDate
`func (o *ReceivedDocumentPaymentsListItem) UnsetPaidDate()`

UnsetPaidDate ensures that no value is present for PaidDate, not even an explicit nil
### GetPaymentTerms

`func (o *ReceivedDocumentPaymentsListItem) GetPaymentTerms() ReceivedDocumentPaymentsListItemPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ReceivedDocumentPaymentsListItem) GetPaymentTermsOk() (*ReceivedDocumentPaymentsListItemPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ReceivedDocumentPaymentsListItem) SetPaymentTerms(v ReceivedDocumentPaymentsListItemPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ReceivedDocumentPaymentsListItem) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ReceivedDocumentPaymentsListItem) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ReceivedDocumentPaymentsListItem) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetStatus

`func (o *ReceivedDocumentPaymentsListItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReceivedDocumentPaymentsListItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReceivedDocumentPaymentsListItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReceivedDocumentPaymentsListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReceivedDocumentPaymentsListItem) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReceivedDocumentPaymentsListItem) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPaymentAccount

`func (o *ReceivedDocumentPaymentsListItem) GetPaymentAccount() PaymentAccount`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *ReceivedDocumentPaymentsListItem) GetPaymentAccountOk() (*PaymentAccount, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *ReceivedDocumentPaymentsListItem) SetPaymentAccount(v PaymentAccount)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *ReceivedDocumentPaymentsListItem) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### SetPaymentAccountNil

`func (o *ReceivedDocumentPaymentsListItem) SetPaymentAccountNil(b bool)`

 SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil

### UnsetPaymentAccount
`func (o *ReceivedDocumentPaymentsListItem) UnsetPaymentAccount()`

UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


