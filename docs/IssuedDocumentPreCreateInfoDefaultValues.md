# IssuedDocumentPreCreateInfoDefaultValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTemplate** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**DnTemplate** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**AiTemplate** | Pointer to [**DocumentTemplate**](DocumentTemplate.md) |  | [optional] 
**Notes** | Pointer to **NullableString** | Default notes. | [optional] 
**Rivalsa** | Pointer to **NullableFloat32** | Default rivalsa percentage. | [optional] 
**Cassa** | Pointer to **NullableFloat32** | Default cassa percentage. | [optional] 
**WithholdingTax** | Pointer to **NullableFloat32** | Default withholding tax percentage. | [optional] 
**WithholdingTaxTaxable** | Pointer to **NullableFloat32** | Default withholding tax taxable percentage. | [optional] 
**OtherWithholdingTax** | Pointer to **NullableFloat32** | Default other withholding tax percentage. | [optional] 
**UseGrossPrices** | Pointer to **NullableBool** | Use gross price by default. | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 

## Methods

### NewIssuedDocumentPreCreateInfoDefaultValues

`func NewIssuedDocumentPreCreateInfoDefaultValues() *IssuedDocumentPreCreateInfoDefaultValues`

NewIssuedDocumentPreCreateInfoDefaultValues instantiates a new IssuedDocumentPreCreateInfoDefaultValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentPreCreateInfoDefaultValuesWithDefaults

`func NewIssuedDocumentPreCreateInfoDefaultValuesWithDefaults() *IssuedDocumentPreCreateInfoDefaultValues`

NewIssuedDocumentPreCreateInfoDefaultValuesWithDefaults instantiates a new IssuedDocumentPreCreateInfoDefaultValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDefaultTemplate() DocumentTemplate`

GetDefaultTemplate returns the DefaultTemplate field if non-nil, zero value otherwise.

### GetDefaultTemplateOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDefaultTemplateOk() (*DocumentTemplate, bool)`

GetDefaultTemplateOk returns a tuple with the DefaultTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetDefaultTemplate(v DocumentTemplate)`

SetDefaultTemplate sets DefaultTemplate field to given value.

### HasDefaultTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasDefaultTemplate() bool`

HasDefaultTemplate returns a boolean if a field has been set.

### GetDnTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDnTemplate() DocumentTemplate`

GetDnTemplate returns the DnTemplate field if non-nil, zero value otherwise.

### GetDnTemplateOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetDnTemplateOk() (*DocumentTemplate, bool)`

GetDnTemplateOk returns a tuple with the DnTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetDnTemplate(v DocumentTemplate)`

SetDnTemplate sets DnTemplate field to given value.

### HasDnTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasDnTemplate() bool`

HasDnTemplate returns a boolean if a field has been set.

### GetAiTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetAiTemplate() DocumentTemplate`

GetAiTemplate returns the AiTemplate field if non-nil, zero value otherwise.

### GetAiTemplateOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetAiTemplateOk() (*DocumentTemplate, bool)`

GetAiTemplateOk returns a tuple with the AiTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetAiTemplate(v DocumentTemplate)`

SetAiTemplate sets AiTemplate field to given value.

### HasAiTemplate

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasAiTemplate() bool`

HasAiTemplate returns a boolean if a field has been set.

### GetNotes

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRivalsa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetRivalsa() float32`

GetRivalsa returns the Rivalsa field if non-nil, zero value otherwise.

### GetRivalsaOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetRivalsaOk() (*float32, bool)`

GetRivalsaOk returns a tuple with the Rivalsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRivalsa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetRivalsa(v float32)`

SetRivalsa sets Rivalsa field to given value.

### HasRivalsa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasRivalsa() bool`

HasRivalsa returns a boolean if a field has been set.

### SetRivalsaNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetRivalsaNil(b bool)`

 SetRivalsaNil sets the value for Rivalsa to be an explicit nil

### UnsetRivalsa
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetRivalsa()`

UnsetRivalsa ensures that no value is present for Rivalsa, not even an explicit nil
### GetCassa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetCassa() float32`

GetCassa returns the Cassa field if non-nil, zero value otherwise.

### GetCassaOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetCassaOk() (*float32, bool)`

GetCassaOk returns a tuple with the Cassa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetCassa(v float32)`

SetCassa sets Cassa field to given value.

### HasCassa

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasCassa() bool`

HasCassa returns a boolean if a field has been set.

### SetCassaNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetCassaNil(b bool)`

 SetCassaNil sets the value for Cassa to be an explicit nil

### UnsetCassa
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetCassa()`

UnsetCassa ensures that no value is present for Cassa, not even an explicit nil
### GetWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTax() float32`

GetWithholdingTax returns the WithholdingTax field if non-nil, zero value otherwise.

### GetWithholdingTaxOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxOk() (*float32, bool)`

GetWithholdingTaxOk returns a tuple with the WithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTax(v float32)`

SetWithholdingTax sets WithholdingTax field to given value.

### HasWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasWithholdingTax() bool`

HasWithholdingTax returns a boolean if a field has been set.

### SetWithholdingTaxNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxNil(b bool)`

 SetWithholdingTaxNil sets the value for WithholdingTax to be an explicit nil

### UnsetWithholdingTax
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetWithholdingTax()`

UnsetWithholdingTax ensures that no value is present for WithholdingTax, not even an explicit nil
### GetWithholdingTaxTaxable

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxTaxable() float32`

GetWithholdingTaxTaxable returns the WithholdingTaxTaxable field if non-nil, zero value otherwise.

### GetWithholdingTaxTaxableOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetWithholdingTaxTaxableOk() (*float32, bool)`

GetWithholdingTaxTaxableOk returns a tuple with the WithholdingTaxTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholdingTaxTaxable

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxTaxable(v float32)`

SetWithholdingTaxTaxable sets WithholdingTaxTaxable field to given value.

### HasWithholdingTaxTaxable

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasWithholdingTaxTaxable() bool`

HasWithholdingTaxTaxable returns a boolean if a field has been set.

### SetWithholdingTaxTaxableNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetWithholdingTaxTaxableNil(b bool)`

 SetWithholdingTaxTaxableNil sets the value for WithholdingTaxTaxable to be an explicit nil

### UnsetWithholdingTaxTaxable
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetWithholdingTaxTaxable()`

UnsetWithholdingTaxTaxable ensures that no value is present for WithholdingTaxTaxable, not even an explicit nil
### GetOtherWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetOtherWithholdingTax() float32`

GetOtherWithholdingTax returns the OtherWithholdingTax field if non-nil, zero value otherwise.

### GetOtherWithholdingTaxOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetOtherWithholdingTaxOk() (*float32, bool)`

GetOtherWithholdingTaxOk returns a tuple with the OtherWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetOtherWithholdingTax(v float32)`

SetOtherWithholdingTax sets OtherWithholdingTax field to given value.

### HasOtherWithholdingTax

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasOtherWithholdingTax() bool`

HasOtherWithholdingTax returns a boolean if a field has been set.

### SetOtherWithholdingTaxNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetOtherWithholdingTaxNil(b bool)`

 SetOtherWithholdingTaxNil sets the value for OtherWithholdingTax to be an explicit nil

### UnsetOtherWithholdingTax
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetOtherWithholdingTax()`

UnsetOtherWithholdingTax ensures that no value is present for OtherWithholdingTax, not even an explicit nil
### GetUseGrossPrices

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetUseGrossPrices() bool`

GetUseGrossPrices returns the UseGrossPrices field if non-nil, zero value otherwise.

### GetUseGrossPricesOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetUseGrossPricesOk() (*bool, bool)`

GetUseGrossPricesOk returns a tuple with the UseGrossPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGrossPrices

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetUseGrossPrices(v bool)`

SetUseGrossPrices sets UseGrossPrices field to given value.

### HasUseGrossPrices

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasUseGrossPrices() bool`

HasUseGrossPrices returns a boolean if a field has been set.

### SetUseGrossPricesNil

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetUseGrossPricesNil(b bool)`

 SetUseGrossPricesNil sets the value for UseGrossPrices to be an explicit nil

### UnsetUseGrossPrices
`func (o *IssuedDocumentPreCreateInfoDefaultValues) UnsetUseGrossPrices()`

UnsetUseGrossPrices ensures that no value is present for UseGrossPrices, not even an explicit nil
### GetPaymentMethod

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *IssuedDocumentPreCreateInfoDefaultValues) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *IssuedDocumentPreCreateInfoDefaultValues) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *IssuedDocumentPreCreateInfoDefaultValues) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


