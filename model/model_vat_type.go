/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// VatType 
type VatType struct {
	// Unique identifier
	Id NullableInt32 `json:"id,omitempty"`
	// [Read Only] Percentual value.
	Value NullableFloat32 `json:"value,omitempty"`
	// Short description.
	Description NullableString `json:"description,omitempty"`
	// Long description and notes shown in documents.
	Notes NullableString `json:"notes,omitempty"`
	// Usable for e-invoices.
	EInvoice NullableBool `json:"e_invoice,omitempty"`
	// E-invoice type (natura).
	EiType NullableString `json:"ei_type,omitempty"`
	// E-invoice description.
	EiDescription NullableString `json:"ei_description,omitempty"`
	// [Read Only] Determine if this vat type is editable.
	Editable NullableBool `json:"editable,omitempty"`
	// Determine if the vat type is disabled.
	IsDisabled NullableBool `json:"is_disabled,omitempty"`
}

// NewVatType instantiates a new VatType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVatType() *VatType {
	this := VatType{}
	return &this
}

// NewVatTypeWithDefaults instantiates a new VatType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVatTypeWithDefaults() *VatType {
	this := VatType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *VatType) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *VatType) SetId(v int32) *VatType {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *VatType) SetIdNil() *VatType {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *VatType) UnsetId() {
	o.Id.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetValue() float32 {
	if o == nil || o.Value.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *VatType) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat32 and assigns it to the Value field.
func (o *VatType) SetValue(v float32) *VatType {
	o.Value.Set(&v)
	return o
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *VatType) SetValueNil() *VatType {
	o.Value.Set(nil)
	return o
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *VatType) UnsetValue() {
	o.Value.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VatType) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VatType) SetDescription(v string) *VatType {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VatType) SetDescriptionNil() *VatType {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VatType) UnsetDescription() {
	o.Description.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *VatType) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *VatType) SetNotes(v string) *VatType {
	o.Notes.Set(&v)
	return o
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *VatType) SetNotesNil() *VatType {
	o.Notes.Set(nil)
	return o
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *VatType) UnsetNotes() {
	o.Notes.Unset()
}

// GetEInvoice returns the EInvoice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetEInvoice() bool {
	if o == nil || o.EInvoice.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EInvoice.Get()
}

// GetEInvoiceOk returns a tuple with the EInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetEInvoiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EInvoice.Get(), o.EInvoice.IsSet()
}

// HasEInvoice returns a boolean if a field has been set.
func (o *VatType) HasEInvoice() bool {
	if o != nil && o.EInvoice.IsSet() {
		return true
	}

	return false
}

// SetEInvoice gets a reference to the given NullableBool and assigns it to the EInvoice field.
func (o *VatType) SetEInvoice(v bool) *VatType {
	o.EInvoice.Set(&v)
	return o
}
// SetEInvoiceNil sets the value for EInvoice to be an explicit nil
func (o *VatType) SetEInvoiceNil() *VatType {
	o.EInvoice.Set(nil)
	return o
}

// UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
func (o *VatType) UnsetEInvoice() {
	o.EInvoice.Unset()
}

// GetEiType returns the EiType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetEiType() string {
	if o == nil || o.EiType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EiType.Get()
}

// GetEiTypeOk returns a tuple with the EiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetEiTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EiType.Get(), o.EiType.IsSet()
}

// HasEiType returns a boolean if a field has been set.
func (o *VatType) HasEiType() bool {
	if o != nil && o.EiType.IsSet() {
		return true
	}

	return false
}

// SetEiType gets a reference to the given NullableString and assigns it to the EiType field.
func (o *VatType) SetEiType(v string) *VatType {
	o.EiType.Set(&v)
	return o
}
// SetEiTypeNil sets the value for EiType to be an explicit nil
func (o *VatType) SetEiTypeNil() *VatType {
	o.EiType.Set(nil)
	return o
}

// UnsetEiType ensures that no value is present for EiType, not even an explicit nil
func (o *VatType) UnsetEiType() {
	o.EiType.Unset()
}

// GetEiDescription returns the EiDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetEiDescription() string {
	if o == nil || o.EiDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.EiDescription.Get()
}

// GetEiDescriptionOk returns a tuple with the EiDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetEiDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EiDescription.Get(), o.EiDescription.IsSet()
}

// HasEiDescription returns a boolean if a field has been set.
func (o *VatType) HasEiDescription() bool {
	if o != nil && o.EiDescription.IsSet() {
		return true
	}

	return false
}

// SetEiDescription gets a reference to the given NullableString and assigns it to the EiDescription field.
func (o *VatType) SetEiDescription(v string) *VatType {
	o.EiDescription.Set(&v)
	return o
}
// SetEiDescriptionNil sets the value for EiDescription to be an explicit nil
func (o *VatType) SetEiDescriptionNil() *VatType {
	o.EiDescription.Set(nil)
	return o
}

// UnsetEiDescription ensures that no value is present for EiDescription, not even an explicit nil
func (o *VatType) UnsetEiDescription() {
	o.EiDescription.Unset()
}

// GetEditable returns the Editable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetEditable() bool {
	if o == nil || o.Editable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Editable.Get()
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Editable.Get(), o.Editable.IsSet()
}

// HasEditable returns a boolean if a field has been set.
func (o *VatType) HasEditable() bool {
	if o != nil && o.Editable.IsSet() {
		return true
	}

	return false
}

// SetEditable gets a reference to the given NullableBool and assigns it to the Editable field.
func (o *VatType) SetEditable(v bool) *VatType {
	o.Editable.Set(&v)
	return o
}
// SetEditableNil sets the value for Editable to be an explicit nil
func (o *VatType) SetEditableNil() *VatType {
	o.Editable.Set(nil)
	return o
}

// UnsetEditable ensures that no value is present for Editable, not even an explicit nil
func (o *VatType) UnsetEditable() {
	o.Editable.Unset()
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VatType) GetIsDisabled() bool {
	if o == nil || o.IsDisabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDisabled.Get()
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VatType) GetIsDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDisabled.Get(), o.IsDisabled.IsSet()
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *VatType) HasIsDisabled() bool {
	if o != nil && o.IsDisabled.IsSet() {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given NullableBool and assigns it to the IsDisabled field.
func (o *VatType) SetIsDisabled(v bool) *VatType {
	o.IsDisabled.Set(&v)
	return o
}
// SetIsDisabledNil sets the value for IsDisabled to be an explicit nil
func (o *VatType) SetIsDisabledNil() *VatType {
	o.IsDisabled.Set(nil)
	return o
}

// UnsetIsDisabled ensures that no value is present for IsDisabled, not even an explicit nil
func (o *VatType) UnsetIsDisabled() {
	o.IsDisabled.Unset()
}

func (o VatType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.EInvoice.IsSet() {
		toSerialize["e_invoice"] = o.EInvoice.Get()
	}
	if o.EiType.IsSet() {
		toSerialize["ei_type"] = o.EiType.Get()
	}
	if o.EiDescription.IsSet() {
		toSerialize["ei_description"] = o.EiDescription.Get()
	}
	if o.Editable.IsSet() {
		toSerialize["editable"] = o.Editable.Get()
	}
	if o.IsDisabled.IsSet() {
		toSerialize["is_disabled"] = o.IsDisabled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVatType struct {
	value *VatType
	isSet bool
}

func (v NullableVatType) Get() *VatType {
	return v.value
}

func (v *NullableVatType) Set(val *VatType) {
	v.value = val
	v.isSet = true
}

func (v NullableVatType) IsSet() bool {
	return v.isSet
}

func (v *NullableVatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVatType(val *VatType) *NullableVatType {
	return &NullableVatType{value: val, isSet: true}
}

func (v NullableVatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


