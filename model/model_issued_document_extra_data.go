/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.3
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the IssuedDocumentExtraData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentExtraData{}

// IssuedDocumentExtraData Issued document extra data [TS fields follow the technical specifications provided by \"Sistema Tessera Sanitaria\"]
type IssuedDocumentExtraData struct {
MultifattureSent NullableInt32 `json:"multifatture_sent,omitempty"`
	// Send issued document to \"Sistema Tessera Sanitaria\"
TsCommunication NullableBool `json:"ts_communication,omitempty"`
	// Issued document ts \"tipo spesa\" [TK, FC, FV, SV,SP, AD, AS, ECG, SR]
TsFlagTipoSpesa NullableFloat32 `json:"ts_flag_tipo_spesa,omitempty"`
	// Issued document ts traced payment
TsPagamentoTracciato NullableBool `json:"ts_pagamento_tracciato,omitempty"`
	// Can be [ 'TK', 'FC', 'FV', 'SV', 'SP', 'AD', 'AS', 'SR', 'CT', 'PI', 'IC', 'AA' ]. Refer to the technical specifications to learn more.
TsTipoSpesa NullableString `json:"ts_tipo_spesa,omitempty"`
	// Issued document ts \"opposizione\"
TsOpposizione NullableBool `json:"ts_opposizione,omitempty"`
	// Issued document ts status
TsStatus NullableInt32 `json:"ts_status,omitempty"`
	// Issued document ts file id
TsFileId NullableString `json:"ts_file_id,omitempty"`
	// Issued document ts sent date
TsSentDate NullableString `json:"ts_sent_date,omitempty"`
	// Issued document ts total amount
TsFullAmount NullableBool `json:"ts_full_amount,omitempty"`
	// Issued document imported by software
ImportedBy NullableString `json:"imported_by,omitempty"`
}

// NewIssuedDocumentExtraData instantiates a new IssuedDocumentExtraData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentExtraData() *IssuedDocumentExtraData {
	this := IssuedDocumentExtraData{}
	return &this
}

// NewIssuedDocumentExtraDataWithDefaults instantiates a new IssuedDocumentExtraData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentExtraDataWithDefaults() *IssuedDocumentExtraData {
	this := IssuedDocumentExtraData{}
	return &this
}

// GetMultifattureSent returns the MultifattureSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetMultifattureSent() int32 {
	if o == nil || IsNil(o.MultifattureSent.Get()) {
		var ret int32
		return ret
	}
	return *o.MultifattureSent.Get()
}

// GetMultifattureSentOk returns a tuple with the MultifattureSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetMultifattureSentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultifattureSent.Get(), o.MultifattureSent.IsSet()
}

// HasMultifattureSent returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasMultifattureSent() bool {
	if o != nil && o.MultifattureSent.IsSet() {
		return true
	}

	return false
}

// SetMultifattureSent gets a reference to the given NullableInt32 and assigns it to the MultifattureSent field.
func (o *IssuedDocumentExtraData) SetMultifattureSent(v int32) *IssuedDocumentExtraData {
	o.MultifattureSent.Set(&v)
        return o
}
// SetMultifattureSentNil sets the value for MultifattureSent to be an explicit nil
func (o *IssuedDocumentExtraData) SetMultifattureSentNil() *IssuedDocumentExtraData {
	o.MultifattureSent.Set(nil)
    return o
}

// UnsetMultifattureSent ensures that no value is present for MultifattureSent, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetMultifattureSent() {
	o.MultifattureSent.Unset()
}

// GetTsCommunication returns the TsCommunication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsCommunication() bool {
	if o == nil || IsNil(o.TsCommunication.Get()) {
		var ret bool
		return ret
	}
	return *o.TsCommunication.Get()
}

// GetTsCommunicationOk returns a tuple with the TsCommunication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsCommunicationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsCommunication.Get(), o.TsCommunication.IsSet()
}

// HasTsCommunication returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsCommunication() bool {
	if o != nil && o.TsCommunication.IsSet() {
		return true
	}

	return false
}

// SetTsCommunication gets a reference to the given NullableBool and assigns it to the TsCommunication field.
func (o *IssuedDocumentExtraData) SetTsCommunication(v bool) *IssuedDocumentExtraData {
	o.TsCommunication.Set(&v)
        return o
}
// SetTsCommunicationNil sets the value for TsCommunication to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsCommunicationNil() *IssuedDocumentExtraData {
	o.TsCommunication.Set(nil)
    return o
}

// UnsetTsCommunication ensures that no value is present for TsCommunication, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsCommunication() {
	o.TsCommunication.Unset()
}

// GetTsFlagTipoSpesa returns the TsFlagTipoSpesa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsFlagTipoSpesa() float32 {
	if o == nil || IsNil(o.TsFlagTipoSpesa.Get()) {
		var ret float32
		return ret
	}
	return *o.TsFlagTipoSpesa.Get()
}

// GetTsFlagTipoSpesaOk returns a tuple with the TsFlagTipoSpesa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsFlagTipoSpesaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsFlagTipoSpesa.Get(), o.TsFlagTipoSpesa.IsSet()
}

// HasTsFlagTipoSpesa returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsFlagTipoSpesa() bool {
	if o != nil && o.TsFlagTipoSpesa.IsSet() {
		return true
	}

	return false
}

// SetTsFlagTipoSpesa gets a reference to the given NullableFloat32 and assigns it to the TsFlagTipoSpesa field.
func (o *IssuedDocumentExtraData) SetTsFlagTipoSpesa(v float32) *IssuedDocumentExtraData {
	o.TsFlagTipoSpesa.Set(&v)
        return o
}
// SetTsFlagTipoSpesaNil sets the value for TsFlagTipoSpesa to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsFlagTipoSpesaNil() *IssuedDocumentExtraData {
	o.TsFlagTipoSpesa.Set(nil)
    return o
}

// UnsetTsFlagTipoSpesa ensures that no value is present for TsFlagTipoSpesa, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsFlagTipoSpesa() {
	o.TsFlagTipoSpesa.Unset()
}

// GetTsPagamentoTracciato returns the TsPagamentoTracciato field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsPagamentoTracciato() bool {
	if o == nil || IsNil(o.TsPagamentoTracciato.Get()) {
		var ret bool
		return ret
	}
	return *o.TsPagamentoTracciato.Get()
}

// GetTsPagamentoTracciatoOk returns a tuple with the TsPagamentoTracciato field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsPagamentoTracciatoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsPagamentoTracciato.Get(), o.TsPagamentoTracciato.IsSet()
}

// HasTsPagamentoTracciato returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsPagamentoTracciato() bool {
	if o != nil && o.TsPagamentoTracciato.IsSet() {
		return true
	}

	return false
}

// SetTsPagamentoTracciato gets a reference to the given NullableBool and assigns it to the TsPagamentoTracciato field.
func (o *IssuedDocumentExtraData) SetTsPagamentoTracciato(v bool) *IssuedDocumentExtraData {
	o.TsPagamentoTracciato.Set(&v)
        return o
}
// SetTsPagamentoTracciatoNil sets the value for TsPagamentoTracciato to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsPagamentoTracciatoNil() *IssuedDocumentExtraData {
	o.TsPagamentoTracciato.Set(nil)
    return o
}

// UnsetTsPagamentoTracciato ensures that no value is present for TsPagamentoTracciato, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsPagamentoTracciato() {
	o.TsPagamentoTracciato.Unset()
}

// GetTsTipoSpesa returns the TsTipoSpesa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsTipoSpesa() string {
	if o == nil || IsNil(o.TsTipoSpesa.Get()) {
		var ret string
		return ret
	}
	return *o.TsTipoSpesa.Get()
}

// GetTsTipoSpesaOk returns a tuple with the TsTipoSpesa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsTipoSpesaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsTipoSpesa.Get(), o.TsTipoSpesa.IsSet()
}

// HasTsTipoSpesa returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsTipoSpesa() bool {
	if o != nil && o.TsTipoSpesa.IsSet() {
		return true
	}

	return false
}

// SetTsTipoSpesa gets a reference to the given NullableString and assigns it to the TsTipoSpesa field.
func (o *IssuedDocumentExtraData) SetTsTipoSpesa(v string) *IssuedDocumentExtraData {
	o.TsTipoSpesa.Set(&v)
        return o
}
// SetTsTipoSpesaNil sets the value for TsTipoSpesa to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsTipoSpesaNil() *IssuedDocumentExtraData {
	o.TsTipoSpesa.Set(nil)
    return o
}

// UnsetTsTipoSpesa ensures that no value is present for TsTipoSpesa, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsTipoSpesa() {
	o.TsTipoSpesa.Unset()
}

// GetTsOpposizione returns the TsOpposizione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsOpposizione() bool {
	if o == nil || IsNil(o.TsOpposizione.Get()) {
		var ret bool
		return ret
	}
	return *o.TsOpposizione.Get()
}

// GetTsOpposizioneOk returns a tuple with the TsOpposizione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsOpposizioneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsOpposizione.Get(), o.TsOpposizione.IsSet()
}

// HasTsOpposizione returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsOpposizione() bool {
	if o != nil && o.TsOpposizione.IsSet() {
		return true
	}

	return false
}

// SetTsOpposizione gets a reference to the given NullableBool and assigns it to the TsOpposizione field.
func (o *IssuedDocumentExtraData) SetTsOpposizione(v bool) *IssuedDocumentExtraData {
	o.TsOpposizione.Set(&v)
        return o
}
// SetTsOpposizioneNil sets the value for TsOpposizione to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsOpposizioneNil() *IssuedDocumentExtraData {
	o.TsOpposizione.Set(nil)
    return o
}

// UnsetTsOpposizione ensures that no value is present for TsOpposizione, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsOpposizione() {
	o.TsOpposizione.Unset()
}

// GetTsStatus returns the TsStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsStatus() int32 {
	if o == nil || IsNil(o.TsStatus.Get()) {
		var ret int32
		return ret
	}
	return *o.TsStatus.Get()
}

// GetTsStatusOk returns a tuple with the TsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsStatus.Get(), o.TsStatus.IsSet()
}

// HasTsStatus returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsStatus() bool {
	if o != nil && o.TsStatus.IsSet() {
		return true
	}

	return false
}

// SetTsStatus gets a reference to the given NullableInt32 and assigns it to the TsStatus field.
func (o *IssuedDocumentExtraData) SetTsStatus(v int32) *IssuedDocumentExtraData {
	o.TsStatus.Set(&v)
        return o
}
// SetTsStatusNil sets the value for TsStatus to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsStatusNil() *IssuedDocumentExtraData {
	o.TsStatus.Set(nil)
    return o
}

// UnsetTsStatus ensures that no value is present for TsStatus, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsStatus() {
	o.TsStatus.Unset()
}

// GetTsFileId returns the TsFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsFileId() string {
	if o == nil || IsNil(o.TsFileId.Get()) {
		var ret string
		return ret
	}
	return *o.TsFileId.Get()
}

// GetTsFileIdOk returns a tuple with the TsFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsFileId.Get(), o.TsFileId.IsSet()
}

// HasTsFileId returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsFileId() bool {
	if o != nil && o.TsFileId.IsSet() {
		return true
	}

	return false
}

// SetTsFileId gets a reference to the given NullableString and assigns it to the TsFileId field.
func (o *IssuedDocumentExtraData) SetTsFileId(v string) *IssuedDocumentExtraData {
	o.TsFileId.Set(&v)
        return o
}
// SetTsFileIdNil sets the value for TsFileId to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsFileIdNil() *IssuedDocumentExtraData {
	o.TsFileId.Set(nil)
    return o
}

// UnsetTsFileId ensures that no value is present for TsFileId, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsFileId() {
	o.TsFileId.Unset()
}

// GetTsSentDate returns the TsSentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsSentDate() string {
	if o == nil || IsNil(o.TsSentDate.Get()) {
		var ret string
		return ret
	}
	return *o.TsSentDate.Get()
}

// GetTsSentDateOk returns a tuple with the TsSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsSentDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsSentDate.Get(), o.TsSentDate.IsSet()
}

// HasTsSentDate returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsSentDate() bool {
	if o != nil && o.TsSentDate.IsSet() {
		return true
	}

	return false
}

// SetTsSentDate gets a reference to the given NullableString and assigns it to the TsSentDate field.
func (o *IssuedDocumentExtraData) SetTsSentDate(v string) *IssuedDocumentExtraData {
	o.TsSentDate.Set(&v)
        return o
}
// SetTsSentDateNil sets the value for TsSentDate to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsSentDateNil() *IssuedDocumentExtraData {
	o.TsSentDate.Set(nil)
    return o
}

// UnsetTsSentDate ensures that no value is present for TsSentDate, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsSentDate() {
	o.TsSentDate.Unset()
}

// GetTsFullAmount returns the TsFullAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetTsFullAmount() bool {
	if o == nil || IsNil(o.TsFullAmount.Get()) {
		var ret bool
		return ret
	}
	return *o.TsFullAmount.Get()
}

// GetTsFullAmountOk returns a tuple with the TsFullAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetTsFullAmountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TsFullAmount.Get(), o.TsFullAmount.IsSet()
}

// HasTsFullAmount returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasTsFullAmount() bool {
	if o != nil && o.TsFullAmount.IsSet() {
		return true
	}

	return false
}

// SetTsFullAmount gets a reference to the given NullableBool and assigns it to the TsFullAmount field.
func (o *IssuedDocumentExtraData) SetTsFullAmount(v bool) *IssuedDocumentExtraData {
	o.TsFullAmount.Set(&v)
        return o
}
// SetTsFullAmountNil sets the value for TsFullAmount to be an explicit nil
func (o *IssuedDocumentExtraData) SetTsFullAmountNil() *IssuedDocumentExtraData {
	o.TsFullAmount.Set(nil)
    return o
}

// UnsetTsFullAmount ensures that no value is present for TsFullAmount, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetTsFullAmount() {
	o.TsFullAmount.Unset()
}

// GetImportedBy returns the ImportedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentExtraData) GetImportedBy() string {
	if o == nil || IsNil(o.ImportedBy.Get()) {
		var ret string
		return ret
	}
	return *o.ImportedBy.Get()
}

// GetImportedByOk returns a tuple with the ImportedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentExtraData) GetImportedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImportedBy.Get(), o.ImportedBy.IsSet()
}

// HasImportedBy returns a boolean if a field has been set.
func (o *IssuedDocumentExtraData) HasImportedBy() bool {
	if o != nil && o.ImportedBy.IsSet() {
		return true
	}

	return false
}

// SetImportedBy gets a reference to the given NullableString and assigns it to the ImportedBy field.
func (o *IssuedDocumentExtraData) SetImportedBy(v string) *IssuedDocumentExtraData {
	o.ImportedBy.Set(&v)
        return o
}
// SetImportedByNil sets the value for ImportedBy to be an explicit nil
func (o *IssuedDocumentExtraData) SetImportedByNil() *IssuedDocumentExtraData {
	o.ImportedBy.Set(nil)
    return o
}

// UnsetImportedBy ensures that no value is present for ImportedBy, not even an explicit nil
func (o *IssuedDocumentExtraData) UnsetImportedBy() {
	o.ImportedBy.Unset()
}

func (o IssuedDocumentExtraData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentExtraData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MultifattureSent.IsSet() {
		toSerialize["multifatture_sent"] = o.MultifattureSent.Get()
	}
	if o.TsCommunication.IsSet() {
		toSerialize["ts_communication"] = o.TsCommunication.Get()
	}
	if o.TsFlagTipoSpesa.IsSet() {
		toSerialize["ts_flag_tipo_spesa"] = o.TsFlagTipoSpesa.Get()
	}
	if o.TsPagamentoTracciato.IsSet() {
		toSerialize["ts_pagamento_tracciato"] = o.TsPagamentoTracciato.Get()
	}
	if o.TsTipoSpesa.IsSet() {
		toSerialize["ts_tipo_spesa"] = o.TsTipoSpesa.Get()
	}
	if o.TsOpposizione.IsSet() {
		toSerialize["ts_opposizione"] = o.TsOpposizione.Get()
	}
	if o.TsStatus.IsSet() {
		toSerialize["ts_status"] = o.TsStatus.Get()
	}
	if o.TsFileId.IsSet() {
		toSerialize["ts_file_id"] = o.TsFileId.Get()
	}
	if o.TsSentDate.IsSet() {
		toSerialize["ts_sent_date"] = o.TsSentDate.Get()
	}
	if o.TsFullAmount.IsSet() {
		toSerialize["ts_full_amount"] = o.TsFullAmount.Get()
	}
	if o.ImportedBy.IsSet() {
		toSerialize["imported_by"] = o.ImportedBy.Get()
	}
	return toSerialize, nil
}

type NullableIssuedDocumentExtraData struct {
	value *IssuedDocumentExtraData
	isSet bool
}

func (v NullableIssuedDocumentExtraData) Get() *IssuedDocumentExtraData {
	return v.value
}

func (v *NullableIssuedDocumentExtraData) Set(val *IssuedDocumentExtraData) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentExtraData) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentExtraData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentExtraData(val *IssuedDocumentExtraData) *NullableIssuedDocumentExtraData {
	return &NullableIssuedDocumentExtraData{value: val, isSet: true}
}

func (v NullableIssuedDocumentExtraData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentExtraData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


