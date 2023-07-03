/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.28
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the EmailScheduleInclude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailScheduleInclude{}

// EmailScheduleInclude struct for EmailScheduleInclude
type EmailScheduleInclude struct {
	// Include a button to view the document
	Document NullableBool `json:"document,omitempty"`
	// Include a button to view the delivery note
	DeliveryNote NullableBool `json:"delivery_note,omitempty"`
	// Include a button to view the attachment
	Attachment NullableBool `json:"attachment,omitempty"`
	// Include a button to view the accompanying invoice
	AccompanyingInvoice NullableBool `json:"accompanying_invoice,omitempty"`
}

// NewEmailScheduleInclude instantiates a new EmailScheduleInclude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailScheduleInclude() *EmailScheduleInclude {
	this := EmailScheduleInclude{}
	return &this
}

// NewEmailScheduleIncludeWithDefaults instantiates a new EmailScheduleInclude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailScheduleIncludeWithDefaults() *EmailScheduleInclude {
	this := EmailScheduleInclude{}
	return &this
}

// GetDocument returns the Document field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailScheduleInclude) GetDocument() bool {
	if o == nil || IsNil(o.Document.Get()) {
		var ret bool
		return ret
	}
	return *o.Document.Get()
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailScheduleInclude) GetDocumentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Document.Get(), o.Document.IsSet()
}

// HasDocument returns a boolean if a field has been set.
func (o *EmailScheduleInclude) HasDocument() bool {
	if o != nil && o.Document.IsSet() {
		return true
	}

	return false
}

// SetDocument gets a reference to the given NullableBool and assigns it to the Document field.
func (o *EmailScheduleInclude) SetDocument(v bool) *EmailScheduleInclude {
	o.Document.Set(&v)
	return o
}
// SetDocumentNil sets the value for Document to be an explicit nil
func (o *EmailScheduleInclude) SetDocumentNil() *EmailScheduleInclude {
	o.Document.Set(nil)
	return o
}

// UnsetDocument ensures that no value is present for Document, not even an explicit nil
func (o *EmailScheduleInclude) UnsetDocument() {
	o.Document.Unset()
}

// GetDeliveryNote returns the DeliveryNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailScheduleInclude) GetDeliveryNote() bool {
	if o == nil || IsNil(o.DeliveryNote.Get()) {
		var ret bool
		return ret
	}
	return *o.DeliveryNote.Get()
}

// GetDeliveryNoteOk returns a tuple with the DeliveryNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailScheduleInclude) GetDeliveryNoteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryNote.Get(), o.DeliveryNote.IsSet()
}

// HasDeliveryNote returns a boolean if a field has been set.
func (o *EmailScheduleInclude) HasDeliveryNote() bool {
	if o != nil && o.DeliveryNote.IsSet() {
		return true
	}

	return false
}

// SetDeliveryNote gets a reference to the given NullableBool and assigns it to the DeliveryNote field.
func (o *EmailScheduleInclude) SetDeliveryNote(v bool) *EmailScheduleInclude {
	o.DeliveryNote.Set(&v)
	return o
}
// SetDeliveryNoteNil sets the value for DeliveryNote to be an explicit nil
func (o *EmailScheduleInclude) SetDeliveryNoteNil() *EmailScheduleInclude {
	o.DeliveryNote.Set(nil)
	return o
}

// UnsetDeliveryNote ensures that no value is present for DeliveryNote, not even an explicit nil
func (o *EmailScheduleInclude) UnsetDeliveryNote() {
	o.DeliveryNote.Unset()
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailScheduleInclude) GetAttachment() bool {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret bool
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailScheduleInclude) GetAttachmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *EmailScheduleInclude) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableBool and assigns it to the Attachment field.
func (o *EmailScheduleInclude) SetAttachment(v bool) *EmailScheduleInclude {
	o.Attachment.Set(&v)
	return o
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *EmailScheduleInclude) SetAttachmentNil() *EmailScheduleInclude {
	o.Attachment.Set(nil)
	return o
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *EmailScheduleInclude) UnsetAttachment() {
	o.Attachment.Unset()
}

// GetAccompanyingInvoice returns the AccompanyingInvoice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailScheduleInclude) GetAccompanyingInvoice() bool {
	if o == nil || IsNil(o.AccompanyingInvoice.Get()) {
		var ret bool
		return ret
	}
	return *o.AccompanyingInvoice.Get()
}

// GetAccompanyingInvoiceOk returns a tuple with the AccompanyingInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailScheduleInclude) GetAccompanyingInvoiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccompanyingInvoice.Get(), o.AccompanyingInvoice.IsSet()
}

// HasAccompanyingInvoice returns a boolean if a field has been set.
func (o *EmailScheduleInclude) HasAccompanyingInvoice() bool {
	if o != nil && o.AccompanyingInvoice.IsSet() {
		return true
	}

	return false
}

// SetAccompanyingInvoice gets a reference to the given NullableBool and assigns it to the AccompanyingInvoice field.
func (o *EmailScheduleInclude) SetAccompanyingInvoice(v bool) *EmailScheduleInclude {
	o.AccompanyingInvoice.Set(&v)
	return o
}
// SetAccompanyingInvoiceNil sets the value for AccompanyingInvoice to be an explicit nil
func (o *EmailScheduleInclude) SetAccompanyingInvoiceNil() *EmailScheduleInclude {
	o.AccompanyingInvoice.Set(nil)
	return o
}

// UnsetAccompanyingInvoice ensures that no value is present for AccompanyingInvoice, not even an explicit nil
func (o *EmailScheduleInclude) UnsetAccompanyingInvoice() {
	o.AccompanyingInvoice.Unset()
}

func (o EmailScheduleInclude) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailScheduleInclude) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Document.IsSet() {
		toSerialize["document"] = o.Document.Get()
	}
	if o.DeliveryNote.IsSet() {
		toSerialize["delivery_note"] = o.DeliveryNote.Get()
	}
	if o.Attachment.IsSet() {
		toSerialize["attachment"] = o.Attachment.Get()
	}
	if o.AccompanyingInvoice.IsSet() {
		toSerialize["accompanying_invoice"] = o.AccompanyingInvoice.Get()
	}
	return toSerialize, nil
}

type NullableEmailScheduleInclude struct {
	value *EmailScheduleInclude
	isSet bool
}

func (v NullableEmailScheduleInclude) Get() *EmailScheduleInclude {
	return v.value
}

func (v *NullableEmailScheduleInclude) Set(val *EmailScheduleInclude) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailScheduleInclude) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailScheduleInclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailScheduleInclude(val *EmailScheduleInclude) *NullableEmailScheduleInclude {
	return &NullableEmailScheduleInclude{value: val, isSet: true}
}

func (v NullableEmailScheduleInclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailScheduleInclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


