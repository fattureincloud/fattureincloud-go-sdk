# ReceivedDocumentPaymentsListItemPaymentTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **NullableInt32** | Received document payment number of days by which the payment must be made | [optional] 
**Type** | Pointer to [**PaymentTermsType**](PaymentTermsType.md) |  | [optional] [default to STANDARD]

## Methods

### NewReceivedDocumentPaymentsListItemPaymentTerms

`func NewReceivedDocumentPaymentsListItemPaymentTerms() *ReceivedDocumentPaymentsListItemPaymentTerms`

NewReceivedDocumentPaymentsListItemPaymentTerms instantiates a new ReceivedDocumentPaymentsListItemPaymentTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentPaymentsListItemPaymentTermsWithDefaults

`func NewReceivedDocumentPaymentsListItemPaymentTermsWithDefaults() *ReceivedDocumentPaymentsListItemPaymentTerms`

NewReceivedDocumentPaymentsListItemPaymentTermsWithDefaults instantiates a new ReceivedDocumentPaymentsListItemPaymentTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetType

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetType() PaymentTermsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) GetTypeOk() (*PaymentTermsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) SetType(v PaymentTermsType)`

SetType sets Type field to given value.

### HasType

`func (o *ReceivedDocumentPaymentsListItemPaymentTerms) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


