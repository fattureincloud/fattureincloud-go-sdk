# IssuedDocumentTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountNet** | Pointer to **NullableFloat32** | Total net amount. | [optional] 
**AmountRivalsa** | Pointer to **NullableFloat32** | Rivalsa amount. | [optional] 
**AmountNetWithRivalsa** | Pointer to **NullableFloat32** | Net amount with rivalsa. | [optional] 
**AmountCassa** | Pointer to **NullableFloat32** | Cassa amount. | [optional] 
**TaxableAmount** | Pointer to **NullableFloat32** | Taxable amount. | [optional] 
**NotTaxableAmount** | Pointer to **NullableFloat32** | Not taxable amount. | [optional] 
**AmountVat** | Pointer to **NullableFloat32** | Total vat amount. | [optional] 
**AmountGross** | Pointer to **NullableFloat32** | Total grosas amount. | [optional] 
**TaxableAmountWithholdingTax** | Pointer to **NullableFloat32** | Taxable withholding tax amount. | [optional] 
**AmountWithholdingTax** | Pointer to **NullableFloat32** | Withholding tax amount. | [optional] 
**TaxableAmountOtherWithholdingTax** | Pointer to **NullableFloat32** | Other withholding tax taxable amount. | [optional] 
**AmountOtherWithholdingTax** | Pointer to **NullableFloat32** | Other withholding tax amount. | [optional] 
**StampDuty** | Pointer to **NullableFloat32** | Stamp duty value [0 if not present]. | [optional] 
**AmountDue** | Pointer to **NullableFloat32** | Total amount due. | [optional] 
**IsEnasarcoMaximalExceeded** | Pointer to **NullableBool** |  | [optional] 
**PaymentsSum** | Pointer to **NullableFloat32** | Payments sum. | [optional] 
**VatList** | Pointer to [**map[string]VatItem**](VatItem.md) |  | [optional] 

## Methods

### NewIssuedDocumentTotals

`func NewIssuedDocumentTotals() *IssuedDocumentTotals`

NewIssuedDocumentTotals instantiates a new IssuedDocumentTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentTotalsWithDefaults

`func NewIssuedDocumentTotalsWithDefaults() *IssuedDocumentTotals`

NewIssuedDocumentTotalsWithDefaults instantiates a new IssuedDocumentTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountNet

`func (o *IssuedDocumentTotals) GetAmountNet() float32`

GetAmountNet returns the AmountNet field if non-nil, zero value otherwise.

### GetAmountNetOk

`func (o *IssuedDocumentTotals) GetAmountNetOk() (*float32, bool)`

GetAmountNetOk returns a tuple with the AmountNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNet

`func (o *IssuedDocumentTotals) SetAmountNet(v float32)`

SetAmountNet sets AmountNet field to given value.

### HasAmountNet

`func (o *IssuedDocumentTotals) HasAmountNet() bool`

HasAmountNet returns a boolean if a field has been set.

### SetAmountNetNil

`func (o *IssuedDocumentTotals) SetAmountNetNil(b bool)`

 SetAmountNetNil sets the value for AmountNet to be an explicit nil

### UnsetAmountNet
`func (o *IssuedDocumentTotals) UnsetAmountNet()`

UnsetAmountNet ensures that no value is present for AmountNet, not even an explicit nil
### GetAmountRivalsa

`func (o *IssuedDocumentTotals) GetAmountRivalsa() float32`

GetAmountRivalsa returns the AmountRivalsa field if non-nil, zero value otherwise.

### GetAmountRivalsaOk

`func (o *IssuedDocumentTotals) GetAmountRivalsaOk() (*float32, bool)`

GetAmountRivalsaOk returns a tuple with the AmountRivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRivalsa

`func (o *IssuedDocumentTotals) SetAmountRivalsa(v float32)`

SetAmountRivalsa sets AmountRivalsa field to given value.

### HasAmountRivalsa

`func (o *IssuedDocumentTotals) HasAmountRivalsa() bool`

HasAmountRivalsa returns a boolean if a field has been set.

### SetAmountRivalsaNil

`func (o *IssuedDocumentTotals) SetAmountRivalsaNil(b bool)`

 SetAmountRivalsaNil sets the value for AmountRivalsa to be an explicit nil

### UnsetAmountRivalsa
`func (o *IssuedDocumentTotals) UnsetAmountRivalsa()`

UnsetAmountRivalsa ensures that no value is present for AmountRivalsa, not even an explicit nil
### GetAmountNetWithRivalsa

`func (o *IssuedDocumentTotals) GetAmountNetWithRivalsa() float32`

GetAmountNetWithRivalsa returns the AmountNetWithRivalsa field if non-nil, zero value otherwise.

### GetAmountNetWithRivalsaOk

`func (o *IssuedDocumentTotals) GetAmountNetWithRivalsaOk() (*float32, bool)`

GetAmountNetWithRivalsaOk returns a tuple with the AmountNetWithRivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountNetWithRivalsa

`func (o *IssuedDocumentTotals) SetAmountNetWithRivalsa(v float32)`

SetAmountNetWithRivalsa sets AmountNetWithRivalsa field to given value.

### HasAmountNetWithRivalsa

`func (o *IssuedDocumentTotals) HasAmountNetWithRivalsa() bool`

HasAmountNetWithRivalsa returns a boolean if a field has been set.

### SetAmountNetWithRivalsaNil

`func (o *IssuedDocumentTotals) SetAmountNetWithRivalsaNil(b bool)`

 SetAmountNetWithRivalsaNil sets the value for AmountNetWithRivalsa to be an explicit nil

### UnsetAmountNetWithRivalsa
`func (o *IssuedDocumentTotals) UnsetAmountNetWithRivalsa()`

UnsetAmountNetWithRivalsa ensures that no value is present for AmountNetWithRivalsa, not even an explicit nil
### GetAmountCassa

`func (o *IssuedDocumentTotals) GetAmountCassa() float32`

GetAmountCassa returns the AmountCassa field if non-nil, zero value otherwise.

### GetAmountCassaOk

`func (o *IssuedDocumentTotals) GetAmountCassaOk() (*float32, bool)`

GetAmountCassaOk returns a tuple with the AmountCassa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCassa

`func (o *IssuedDocumentTotals) SetAmountCassa(v float32)`

SetAmountCassa sets AmountCassa field to given value.

### HasAmountCassa

`func (o *IssuedDocumentTotals) HasAmountCassa() bool`

HasAmountCassa returns a boolean if a field has been set.

### SetAmountCassaNil

`func (o *IssuedDocumentTotals) SetAmountCassaNil(b bool)`

 SetAmountCassaNil sets the value for AmountCassa to be an explicit nil

### UnsetAmountCassa
`func (o *IssuedDocumentTotals) UnsetAmountCassa()`

UnsetAmountCassa ensures that no value is present for AmountCassa, not even an explicit nil
### GetTaxableAmount

`func (o *IssuedDocumentTotals) GetTaxableAmount() float32`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *IssuedDocumentTotals) GetTaxableAmountOk() (*float32, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *IssuedDocumentTotals) SetTaxableAmount(v float32)`

SetTaxableAmount sets TaxableAmount field to given value.

### HasTaxableAmount

`func (o *IssuedDocumentTotals) HasTaxableAmount() bool`

HasTaxableAmount returns a boolean if a field has been set.

### SetTaxableAmountNil

`func (o *IssuedDocumentTotals) SetTaxableAmountNil(b bool)`

 SetTaxableAmountNil sets the value for TaxableAmount to be an explicit nil

### UnsetTaxableAmount
`func (o *IssuedDocumentTotals) UnsetTaxableAmount()`

UnsetTaxableAmount ensures that no value is present for TaxableAmount, not even an explicit nil
### GetNotTaxableAmount

`func (o *IssuedDocumentTotals) GetNotTaxableAmount() float32`

GetNotTaxableAmount returns the NotTaxableAmount field if non-nil, zero value otherwise.

### GetNotTaxableAmountOk

`func (o *IssuedDocumentTotals) GetNotTaxableAmountOk() (*float32, bool)`

GetNotTaxableAmountOk returns a tuple with the NotTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotTaxableAmount

`func (o *IssuedDocumentTotals) SetNotTaxableAmount(v float32)`

SetNotTaxableAmount sets NotTaxableAmount field to given value.

### HasNotTaxableAmount

`func (o *IssuedDocumentTotals) HasNotTaxableAmount() bool`

HasNotTaxableAmount returns a boolean if a field has been set.

### SetNotTaxableAmountNil

`func (o *IssuedDocumentTotals) SetNotTaxableAmountNil(b bool)`

 SetNotTaxableAmountNil sets the value for NotTaxableAmount to be an explicit nil

### UnsetNotTaxableAmount
`func (o *IssuedDocumentTotals) UnsetNotTaxableAmount()`

UnsetNotTaxableAmount ensures that no value is present for NotTaxableAmount, not even an explicit nil
### GetAmountVat

`func (o *IssuedDocumentTotals) GetAmountVat() float32`

GetAmountVat returns the AmountVat field if non-nil, zero value otherwise.

### GetAmountVatOk

`func (o *IssuedDocumentTotals) GetAmountVatOk() (*float32, bool)`

GetAmountVatOk returns a tuple with the AmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountVat

`func (o *IssuedDocumentTotals) SetAmountVat(v float32)`

SetAmountVat sets AmountVat field to given value.

### HasAmountVat

`func (o *IssuedDocumentTotals) HasAmountVat() bool`

HasAmountVat returns a boolean if a field has been set.

### SetAmountVatNil

`func (o *IssuedDocumentTotals) SetAmountVatNil(b bool)`

 SetAmountVatNil sets the value for AmountVat to be an explicit nil

### UnsetAmountVat
`func (o *IssuedDocumentTotals) UnsetAmountVat()`

UnsetAmountVat ensures that no value is present for AmountVat, not even an explicit nil
### GetAmountGross

`func (o *IssuedDocumentTotals) GetAmountGross() float32`

GetAmountGross returns the AmountGross field if non-nil, zero value otherwise.

### GetAmountGrossOk

`func (o *IssuedDocumentTotals) GetAmountGrossOk() (*float32, bool)`

GetAmountGrossOk returns a tuple with the AmountGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGross

`func (o *IssuedDocumentTotals) SetAmountGross(v float32)`

SetAmountGross sets AmountGross field to given value.

### HasAmountGross

`func (o *IssuedDocumentTotals) HasAmountGross() bool`

HasAmountGross returns a boolean if a field has been set.

### SetAmountGrossNil

`func (o *IssuedDocumentTotals) SetAmountGrossNil(b bool)`

 SetAmountGrossNil sets the value for AmountGross to be an explicit nil

### UnsetAmountGross
`func (o *IssuedDocumentTotals) UnsetAmountGross()`

UnsetAmountGross ensures that no value is present for AmountGross, not even an explicit nil
### GetTaxableAmountWithholdingTax

`func (o *IssuedDocumentTotals) GetTaxableAmountWithholdingTax() float32`

GetTaxableAmountWithholdingTax returns the TaxableAmountWithholdingTax field if non-nil, zero value otherwise.

### GetTaxableAmountWithholdingTaxOk

`func (o *IssuedDocumentTotals) GetTaxableAmountWithholdingTaxOk() (*float32, bool)`

GetTaxableAmountWithholdingTaxOk returns a tuple with the TaxableAmountWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmountWithholdingTax

`func (o *IssuedDocumentTotals) SetTaxableAmountWithholdingTax(v float32)`

SetTaxableAmountWithholdingTax sets TaxableAmountWithholdingTax field to given value.

### HasTaxableAmountWithholdingTax

`func (o *IssuedDocumentTotals) HasTaxableAmountWithholdingTax() bool`

HasTaxableAmountWithholdingTax returns a boolean if a field has been set.

### SetTaxableAmountWithholdingTaxNil

`func (o *IssuedDocumentTotals) SetTaxableAmountWithholdingTaxNil(b bool)`

 SetTaxableAmountWithholdingTaxNil sets the value for TaxableAmountWithholdingTax to be an explicit nil

### UnsetTaxableAmountWithholdingTax
`func (o *IssuedDocumentTotals) UnsetTaxableAmountWithholdingTax()`

UnsetTaxableAmountWithholdingTax ensures that no value is present for TaxableAmountWithholdingTax, not even an explicit nil
### GetAmountWithholdingTax

`func (o *IssuedDocumentTotals) GetAmountWithholdingTax() float32`

GetAmountWithholdingTax returns the AmountWithholdingTax field if non-nil, zero value otherwise.

### GetAmountWithholdingTaxOk

`func (o *IssuedDocumentTotals) GetAmountWithholdingTaxOk() (*float32, bool)`

GetAmountWithholdingTaxOk returns a tuple with the AmountWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountWithholdingTax

`func (o *IssuedDocumentTotals) SetAmountWithholdingTax(v float32)`

SetAmountWithholdingTax sets AmountWithholdingTax field to given value.

### HasAmountWithholdingTax

`func (o *IssuedDocumentTotals) HasAmountWithholdingTax() bool`

HasAmountWithholdingTax returns a boolean if a field has been set.

### SetAmountWithholdingTaxNil

`func (o *IssuedDocumentTotals) SetAmountWithholdingTaxNil(b bool)`

 SetAmountWithholdingTaxNil sets the value for AmountWithholdingTax to be an explicit nil

### UnsetAmountWithholdingTax
`func (o *IssuedDocumentTotals) UnsetAmountWithholdingTax()`

UnsetAmountWithholdingTax ensures that no value is present for AmountWithholdingTax, not even an explicit nil
### GetTaxableAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) GetTaxableAmountOtherWithholdingTax() float32`

GetTaxableAmountOtherWithholdingTax returns the TaxableAmountOtherWithholdingTax field if non-nil, zero value otherwise.

### GetTaxableAmountOtherWithholdingTaxOk

`func (o *IssuedDocumentTotals) GetTaxableAmountOtherWithholdingTaxOk() (*float32, bool)`

GetTaxableAmountOtherWithholdingTaxOk returns a tuple with the TaxableAmountOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) SetTaxableAmountOtherWithholdingTax(v float32)`

SetTaxableAmountOtherWithholdingTax sets TaxableAmountOtherWithholdingTax field to given value.

### HasTaxableAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) HasTaxableAmountOtherWithholdingTax() bool`

HasTaxableAmountOtherWithholdingTax returns a boolean if a field has been set.

### SetTaxableAmountOtherWithholdingTaxNil

`func (o *IssuedDocumentTotals) SetTaxableAmountOtherWithholdingTaxNil(b bool)`

 SetTaxableAmountOtherWithholdingTaxNil sets the value for TaxableAmountOtherWithholdingTax to be an explicit nil

### UnsetTaxableAmountOtherWithholdingTax
`func (o *IssuedDocumentTotals) UnsetTaxableAmountOtherWithholdingTax()`

UnsetTaxableAmountOtherWithholdingTax ensures that no value is present for TaxableAmountOtherWithholdingTax, not even an explicit nil
### GetAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) GetAmountOtherWithholdingTax() float32`

GetAmountOtherWithholdingTax returns the AmountOtherWithholdingTax field if non-nil, zero value otherwise.

### GetAmountOtherWithholdingTaxOk

`func (o *IssuedDocumentTotals) GetAmountOtherWithholdingTaxOk() (*float32, bool)`

GetAmountOtherWithholdingTaxOk returns a tuple with the AmountOtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) SetAmountOtherWithholdingTax(v float32)`

SetAmountOtherWithholdingTax sets AmountOtherWithholdingTax field to given value.

### HasAmountOtherWithholdingTax

`func (o *IssuedDocumentTotals) HasAmountOtherWithholdingTax() bool`

HasAmountOtherWithholdingTax returns a boolean if a field has been set.

### SetAmountOtherWithholdingTaxNil

`func (o *IssuedDocumentTotals) SetAmountOtherWithholdingTaxNil(b bool)`

 SetAmountOtherWithholdingTaxNil sets the value for AmountOtherWithholdingTax to be an explicit nil

### UnsetAmountOtherWithholdingTax
`func (o *IssuedDocumentTotals) UnsetAmountOtherWithholdingTax()`

UnsetAmountOtherWithholdingTax ensures that no value is present for AmountOtherWithholdingTax, not even an explicit nil
### GetStampDuty

`func (o *IssuedDocumentTotals) GetStampDuty() float32`

GetStampDuty returns the StampDuty field if non-nil, zero value otherwise.

### GetStampDutyOk

`func (o *IssuedDocumentTotals) GetStampDutyOk() (*float32, bool)`

GetStampDutyOk returns a tuple with the StampDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDuty

`func (o *IssuedDocumentTotals) SetStampDuty(v float32)`

SetStampDuty sets StampDuty field to given value.

### HasStampDuty

`func (o *IssuedDocumentTotals) HasStampDuty() bool`

HasStampDuty returns a boolean if a field has been set.

### SetStampDutyNil

`func (o *IssuedDocumentTotals) SetStampDutyNil(b bool)`

 SetStampDutyNil sets the value for StampDuty to be an explicit nil

### UnsetStampDuty
`func (o *IssuedDocumentTotals) UnsetStampDuty()`

UnsetStampDuty ensures that no value is present for StampDuty, not even an explicit nil
### GetAmountDue

`func (o *IssuedDocumentTotals) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *IssuedDocumentTotals) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *IssuedDocumentTotals) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *IssuedDocumentTotals) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### SetAmountDueNil

`func (o *IssuedDocumentTotals) SetAmountDueNil(b bool)`

 SetAmountDueNil sets the value for AmountDue to be an explicit nil

### UnsetAmountDue
`func (o *IssuedDocumentTotals) UnsetAmountDue()`

UnsetAmountDue ensures that no value is present for AmountDue, not even an explicit nil
### GetIsEnasarcoMaximalExceeded

`func (o *IssuedDocumentTotals) GetIsEnasarcoMaximalExceeded() bool`

GetIsEnasarcoMaximalExceeded returns the IsEnasarcoMaximalExceeded field if non-nil, zero value otherwise.

### GetIsEnasarcoMaximalExceededOk

`func (o *IssuedDocumentTotals) GetIsEnasarcoMaximalExceededOk() (*bool, bool)`

GetIsEnasarcoMaximalExceededOk returns a tuple with the IsEnasarcoMaximalExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnasarcoMaximalExceeded

`func (o *IssuedDocumentTotals) SetIsEnasarcoMaximalExceeded(v bool)`

SetIsEnasarcoMaximalExceeded sets IsEnasarcoMaximalExceeded field to given value.

### HasIsEnasarcoMaximalExceeded

`func (o *IssuedDocumentTotals) HasIsEnasarcoMaximalExceeded() bool`

HasIsEnasarcoMaximalExceeded returns a boolean if a field has been set.

### SetIsEnasarcoMaximalExceededNil

`func (o *IssuedDocumentTotals) SetIsEnasarcoMaximalExceededNil(b bool)`

 SetIsEnasarcoMaximalExceededNil sets the value for IsEnasarcoMaximalExceeded to be an explicit nil

### UnsetIsEnasarcoMaximalExceeded
`func (o *IssuedDocumentTotals) UnsetIsEnasarcoMaximalExceeded()`

UnsetIsEnasarcoMaximalExceeded ensures that no value is present for IsEnasarcoMaximalExceeded, not even an explicit nil
### GetPaymentsSum

`func (o *IssuedDocumentTotals) GetPaymentsSum() float32`

GetPaymentsSum returns the PaymentsSum field if non-nil, zero value otherwise.

### GetPaymentsSumOk

`func (o *IssuedDocumentTotals) GetPaymentsSumOk() (*float32, bool)`

GetPaymentsSumOk returns a tuple with the PaymentsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsSum

`func (o *IssuedDocumentTotals) SetPaymentsSum(v float32)`

SetPaymentsSum sets PaymentsSum field to given value.

### HasPaymentsSum

`func (o *IssuedDocumentTotals) HasPaymentsSum() bool`

HasPaymentsSum returns a boolean if a field has been set.

### SetPaymentsSumNil

`func (o *IssuedDocumentTotals) SetPaymentsSumNil(b bool)`

 SetPaymentsSumNil sets the value for PaymentsSum to be an explicit nil

### UnsetPaymentsSum
`func (o *IssuedDocumentTotals) UnsetPaymentsSum()`

UnsetPaymentsSum ensures that no value is present for PaymentsSum, not even an explicit nil
### GetVatList

`func (o *IssuedDocumentTotals) GetVatList() map[string]VatItem`

GetVatList returns the VatList field if non-nil, zero value otherwise.

### GetVatListOk

`func (o *IssuedDocumentTotals) GetVatListOk() (*map[string]VatItem, bool)`

GetVatListOk returns a tuple with the VatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatList

`func (o *IssuedDocumentTotals) SetVatList(v map[string]VatItem)`

SetVatList sets VatList field to given value.

### HasVatList

`func (o *IssuedDocumentTotals) HasVatList() bool`

HasVatList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


