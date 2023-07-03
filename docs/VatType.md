# VatType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Vat type id | [optional] 
**Value** | Pointer to **NullableFloat32** | [Read Only] Vat type percentual value | [optional] 
**Description** | Pointer to **NullableString** | Vat type short description | [optional] 
**Notes** | Pointer to **NullableString** | Vat type notes shown in documents | [optional] 
**EInvoice** | Pointer to **NullableBool** | Vat type is usable for e-invoices | [optional] 
**EiType** | Pointer to **NullableString** | Vat type e-invoice type (natura) | [optional] 
**EiDescription** | Pointer to **NullableString** | Vat type e-invoice description | [optional] 
**Editable** | Pointer to **NullableBool** | [Read Only] Is the vat type is editable. | [optional] [readonly] 
**IsDisabled** | Pointer to **NullableBool** | Is the vat type disabled | [optional] 

## Methods

### NewVatType

`func NewVatType() *VatType`

NewVatType instantiates a new VatType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVatTypeWithDefaults

`func NewVatTypeWithDefaults() *VatType`

NewVatTypeWithDefaults instantiates a new VatType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VatType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VatType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VatType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VatType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VatType) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VatType) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetValue

`func (o *VatType) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VatType) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VatType) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *VatType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VatType) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VatType) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *VatType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VatType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VatType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VatType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VatType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VatType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNotes

`func (o *VatType) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VatType) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VatType) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VatType) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VatType) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VatType) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEInvoice

`func (o *VatType) GetEInvoice() bool`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *VatType) GetEInvoiceOk() (*bool, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *VatType) SetEInvoice(v bool)`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *VatType) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### SetEInvoiceNil

`func (o *VatType) SetEInvoiceNil(b bool)`

 SetEInvoiceNil sets the value for EInvoice to be an explicit nil

### UnsetEInvoice
`func (o *VatType) UnsetEInvoice()`

UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
### GetEiType

`func (o *VatType) GetEiType() string`

GetEiType returns the EiType field if non-nil, zero value otherwise.

### GetEiTypeOk

`func (o *VatType) GetEiTypeOk() (*string, bool)`

GetEiTypeOk returns a tuple with the EiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiType

`func (o *VatType) SetEiType(v string)`

SetEiType sets EiType field to given value.

### HasEiType

`func (o *VatType) HasEiType() bool`

HasEiType returns a boolean if a field has been set.

### SetEiTypeNil

`func (o *VatType) SetEiTypeNil(b bool)`

 SetEiTypeNil sets the value for EiType to be an explicit nil

### UnsetEiType
`func (o *VatType) UnsetEiType()`

UnsetEiType ensures that no value is present for EiType, not even an explicit nil
### GetEiDescription

`func (o *VatType) GetEiDescription() string`

GetEiDescription returns the EiDescription field if non-nil, zero value otherwise.

### GetEiDescriptionOk

`func (o *VatType) GetEiDescriptionOk() (*string, bool)`

GetEiDescriptionOk returns a tuple with the EiDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiDescription

`func (o *VatType) SetEiDescription(v string)`

SetEiDescription sets EiDescription field to given value.

### HasEiDescription

`func (o *VatType) HasEiDescription() bool`

HasEiDescription returns a boolean if a field has been set.

### SetEiDescriptionNil

`func (o *VatType) SetEiDescriptionNil(b bool)`

 SetEiDescriptionNil sets the value for EiDescription to be an explicit nil

### UnsetEiDescription
`func (o *VatType) UnsetEiDescription()`

UnsetEiDescription ensures that no value is present for EiDescription, not even an explicit nil
### GetEditable

`func (o *VatType) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *VatType) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *VatType) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *VatType) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *VatType) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *VatType) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetIsDisabled

`func (o *VatType) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *VatType) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *VatType) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *VatType) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### SetIsDisabledNil

`func (o *VatType) SetIsDisabledNil(b bool)`

 SetIsDisabledNil sets the value for IsDisabled to be an explicit nil

### UnsetIsDisabled
`func (o *VatType) UnsetIsDisabled()`

UnsetIsDisabled ensures that no value is present for IsDisabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


