# ReceivedDocumentTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountNet** | Pointer to **NullableFloat32** | Total net amount. | [optional] 
**AmountVat** | Pointer to **NullableFloat32** | Total vat amount. | [optional] 
**AmountGross** | Pointer to **NullableFloat32** | Total gross amount. | [optional] 
**AmountWithholdingTax** | Pointer to **NullableFloat32** | Total withholding tax amount. | [optional] 
**AmountOtherWithholdingTax** | Pointer to **NullableFloat32** | Total other withholding tax amount. | [optional] 
**AmountDue** | Pointer to **NullableFloat32** | Total amount due. | [optional] 
**PaymentsSum** | Pointer to **NullableFloat32** | Payments sum. | [optional] 

## Methods

### NewReceivedDocumentTotals

`func NewReceivedDocumentTotals() *ReceivedDocumentTotals`

NewReceivedDocumentTotals instantiates a new ReceivedDocumentTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivedDocumentTotalsWithDefaults

`func NewReceivedDocumentTotalsWithDefaults() *ReceivedDocumentTotals`

NewReceivedDocumentTotalsWithDefaults instantiates a new ReceivedDocumentTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountNet

`func (o *ReceivedDocumentTotals) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *ReceivedDocumentTotals) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *ReceivedDocumentTotals) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *ReceivedDocumentTotals) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *ReceivedDocumentTotals) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *ReceivedDocumentTotals) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountVat

`func (o *ReceivedDocumentTotals) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *ReceivedDocumentTotals) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *ReceivedDocumentTotals) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *ReceivedDocumentTotals) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *ReceivedDocumentTotals) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *ReceivedDocumentTotals) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetAmountGross

`func (o *ReceivedDocumentTotals) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *ReceivedDocumentTotals) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *ReceivedDocumentTotals) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *ReceivedDocumentTotals) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *ReceivedDocumentTotals) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *ReceivedDocumentTotals) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetAmountWithholdingTax

`func (o *ReceivedDocumentTotals) GetAmountWithholdingTax() float32`

GetAmountWithholdingTax returns the AmountWithholdingTax field if non-nil, zero value otherwise.

### GetAmountWithholdingTaxOk

`func (o *ReceivedDocumentTotals) GetAmountWithholdingTaxOk() (*float32, bool)`

GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountWithholdingTax

`func (o *ReceivedDocumentTotals) SetAmountWithholdingTax(v float32)`

SetAmountWithholdingTax sets AmountWithholdingTax field to given value.

### HasAmountWithholdingTax

`func (o *ReceivedDocumentTotals) HasAmountWithholdingTax() bool`

HasAmountWithholdingTax returns a boolean if a field has been set.

### SetAmountWithholdingTaxNil

`func (o *ReceivedDocumentTotals) SetAmountWithholdingTaxNil(b bool)`

 SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil

### UnsetAmountWithholdingTax
`func (o *ReceivedDocumentTotals) UnsetAmountWithholdingTax()`

UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
### GetAmountOtherWithholdingTax

`func (o *ReceivedDocumentTotals) GetAmountOtherWithholdingTax() float32`

GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field if non-nil, zero value otherwise.

### GetAmountOtherWithholdingTaxOk

`func (o *ReceivedDocumentTotals) GetAmountOtherWithholdingTaxOk() (*float32, bool)`

GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOtherWithholdingTax

`func (o *ReceivedDocumentTotals) SetAmountOtherWithholdingTax(v float32)`

SetAmountOtherWithholdingTax sets AmountOtherWithholdingTax field to given value.

### HasAmountOtherWithholdingTax

`func (o *ReceivedDocumentTotals) HasAmountOtherWithholdingTax() bool`

HasAmountOtherWithholdingTax returns a boolean if a field has been set.

### SetAmountOtherWithholdingTaxNil

`func (o *ReceivedDocumentTotals) SetAmountOtherWithholdingTaxNil(b bool)`

 SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil

### UnsetAmountOtherWithholdingTax
`func (o *ReceivedDocumentTotals) UnsetAmountOtherWithholdingTax()`

UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
### GetAmountDue

`func (o *ReceivedDocumentTotals) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *ReceivedDocumentTotals) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *ReceivedDocumentTotals) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *ReceivedDocumentTotals) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### SetAmountDueNil

`func (o *ReceivedDocumentTotals) SetAmountDueNil(b bool)`

 SetAmountDueNil sets the value for AmountDue to be an explicit nil

### UnsetAmountDue
`func (o *ReceivedDocumentTotals) UnsetAmountDue()`

UnsetAmountDue ensures that no value is present for AmountDue, not even an explicit nil
### GetPaymentsSum

`func (o *ReceivedDocumentTotals) GetPaymentsSum() float32`

GetPaymentsSum returns the PaymentsSum field if non-nil, zero value otherwise.

### GetPaymentsSumOk

`func (o *ReceivedDocumentTotals) GetPaymentsSumOk() (*float32, bool)`

GetPaymentsSumOk returns a tuple with the PaymentsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsSum

`func (o *ReceivedDocumentTotals) SetPaymentsSum(v float32)`

SetPaymentsSum sets PaymentsSum field to given value.

### HasPaymentsSum

`func (o *ReceivedDocumentTotals) HasPaymentsSum() bool`

HasPaymentsSum returns a boolean if a field has been set.

### SetPaymentsSumNil

`func (o *ReceivedDocumentTotals) SetPaymentsSumNil(b bool)`

 SetPaymentsSumNil sets the value for PaymentsSum to be an explicit nil

### UnsetPaymentsSum
`func (o *ReceivedDocumentTotals) UnsetPaymentsSum()`

UnsetPaymentsSum ensures that no value is present for PaymentsSum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


