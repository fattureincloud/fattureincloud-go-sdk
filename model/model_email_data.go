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

// checks if the EmailData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailData{}

// EmailData struct for EmailData
type EmailData struct {
	// Email recipient
	RecipientEmail NullableString `json:"recipient_email,omitempty"`
	DefaultSenderEmail NullableEmailDataDefaultSenderEmail `json:"default_sender_email,omitempty"`
	// List of all emails from which the document can be sent
	SenderEmailsList []SenderEmail `json:"sender_emails_list,omitempty"`
	// Email cc [by default is the logged company email]
	CcEmail NullableString `json:"cc_email,omitempty"`
	// Email subject
	Subject NullableString `json:"subject,omitempty"`
	// Email body
	Body NullableString `json:"body,omitempty"`
	// Document exists if it is not a delivery note
	DocumentExists NullableBool `json:"document_exists,omitempty"`
	// Document is a delivery note
	DeliveryNoteExists NullableBool `json:"delivery_note_exists,omitempty"`
	// Document has attachment
	AttachmentExists NullableBool `json:"attachment_exists,omitempty"`
	// Document has accompanying invoice
	AccompanyingInvoiceExists NullableBool `json:"accompanying_invoice_exists,omitempty"`
	// Attach document pdf
	DefaultAttachPdf NullableBool `json:"default_attach_pdf,omitempty"`
}

// NewEmailData instantiates a new EmailData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailData() *EmailData {
	this := EmailData{}
	return &this
}

// NewEmailDataWithDefaults instantiates a new EmailData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDataWithDefaults() *EmailData {
	this := EmailData{}
	return &this
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetRecipientEmail() string {
	if o == nil || IsNil(o.RecipientEmail.Get()) {
		var ret string
		return ret
	}
	return *o.RecipientEmail.Get()
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetRecipientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientEmail.Get(), o.RecipientEmail.IsSet()
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *EmailData) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail.IsSet() {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given NullableString and assigns it to the RecipientEmail field.
func (o *EmailData) SetRecipientEmail(v string) *EmailData {
	o.RecipientEmail.Set(&v)
		return o
}
// SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil
func (o *EmailData) SetRecipientEmailNil() *EmailData {
	o.RecipientEmail.Set(nil)
	return o
}

// UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
func (o *EmailData) UnsetRecipientEmail() {
	o.RecipientEmail.Unset()
}

// GetDefaultSenderEmail returns the DefaultSenderEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetDefaultSenderEmail() EmailDataDefaultSenderEmail {
	if o == nil || IsNil(o.DefaultSenderEmail.Get()) {
		var ret EmailDataDefaultSenderEmail
		return ret
	}
	return *o.DefaultSenderEmail.Get()
}

// GetDefaultSenderEmailOk returns a tuple with the DefaultSenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetDefaultSenderEmailOk() (*EmailDataDefaultSenderEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultSenderEmail.Get(), o.DefaultSenderEmail.IsSet()
}

// HasDefaultSenderEmail returns a boolean if a field has been set.
func (o *EmailData) HasDefaultSenderEmail() bool {
	if o != nil && o.DefaultSenderEmail.IsSet() {
		return true
	}

	return false
}

// SetDefaultSenderEmail gets a reference to the given NullableEmailDataDefaultSenderEmail and assigns it to the DefaultSenderEmail field.
func (o *EmailData) SetDefaultSenderEmail(v EmailDataDefaultSenderEmail) *EmailData {
	o.DefaultSenderEmail.Set(&v)
		return o
}
// SetDefaultSenderEmailNil sets the value for DefaultSenderEmail to be an explicit nil
func (o *EmailData) SetDefaultSenderEmailNil() *EmailData {
	o.DefaultSenderEmail.Set(nil)
	return o
}

// UnsetDefaultSenderEmail ensures that no value is present for DefaultSenderEmail, not even an explicit nil
func (o *EmailData) UnsetDefaultSenderEmail() {
	o.DefaultSenderEmail.Unset()
}

// GetSenderEmailsList returns the SenderEmailsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetSenderEmailsList() []SenderEmail {
	if o == nil {
		var ret []SenderEmail
		return ret
	}
	return o.SenderEmailsList
}

// GetSenderEmailsListOk returns a tuple with the SenderEmailsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetSenderEmailsListOk() ([]SenderEmail, bool) {
	if o == nil || IsNil(o.SenderEmailsList) {
		return nil, false
	}
	return o.SenderEmailsList, true
}

// HasSenderEmailsList returns a boolean if a field has been set.
func (o *EmailData) HasSenderEmailsList() bool {
	if o != nil && !IsNil(o.SenderEmailsList) {
		return true
	}

	return false
}

// SetSenderEmailsList gets a reference to the given []SenderEmail and assigns it to the SenderEmailsList field.
func (o *EmailData) SetSenderEmailsList(v []SenderEmail) *EmailData {
	o.SenderEmailsList = v
		return o
}

// GetCcEmail returns the CcEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetCcEmail() string {
	if o == nil || IsNil(o.CcEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CcEmail.Get()
}

// GetCcEmailOk returns a tuple with the CcEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetCcEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CcEmail.Get(), o.CcEmail.IsSet()
}

// HasCcEmail returns a boolean if a field has been set.
func (o *EmailData) HasCcEmail() bool {
	if o != nil && o.CcEmail.IsSet() {
		return true
	}

	return false
}

// SetCcEmail gets a reference to the given NullableString and assigns it to the CcEmail field.
func (o *EmailData) SetCcEmail(v string) *EmailData {
	o.CcEmail.Set(&v)
		return o
}
// SetCcEmailNil sets the value for CcEmail to be an explicit nil
func (o *EmailData) SetCcEmailNil() *EmailData {
	o.CcEmail.Set(nil)
	return o
}

// UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
func (o *EmailData) UnsetCcEmail() {
	o.CcEmail.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailData) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *EmailData) SetSubject(v string) *EmailData {
	o.Subject.Set(&v)
		return o
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *EmailData) SetSubjectNil() *EmailData {
	o.Subject.Set(nil)
	return o
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *EmailData) UnsetSubject() {
	o.Subject.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *EmailData) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *EmailData) SetBody(v string) *EmailData {
	o.Body.Set(&v)
		return o
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *EmailData) SetBodyNil() *EmailData {
	o.Body.Set(nil)
	return o
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *EmailData) UnsetBody() {
	o.Body.Unset()
}

// GetDocumentExists returns the DocumentExists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetDocumentExists() bool {
	if o == nil || IsNil(o.DocumentExists.Get()) {
		var ret bool
		return ret
	}
	return *o.DocumentExists.Get()
}

// GetDocumentExistsOk returns a tuple with the DocumentExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetDocumentExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentExists.Get(), o.DocumentExists.IsSet()
}

// HasDocumentExists returns a boolean if a field has been set.
func (o *EmailData) HasDocumentExists() bool {
	if o != nil && o.DocumentExists.IsSet() {
		return true
	}

	return false
}

// SetDocumentExists gets a reference to the given NullableBool and assigns it to the DocumentExists field.
func (o *EmailData) SetDocumentExists(v bool) *EmailData {
	o.DocumentExists.Set(&v)
		return o
}
// SetDocumentExistsNil sets the value for DocumentExists to be an explicit nil
func (o *EmailData) SetDocumentExistsNil() *EmailData {
	o.DocumentExists.Set(nil)
	return o
}

// UnsetDocumentExists ensures that no value is present for DocumentExists, not even an explicit nil
func (o *EmailData) UnsetDocumentExists() {
	o.DocumentExists.Unset()
}

// GetDeliveryNoteExists returns the DeliveryNoteExists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetDeliveryNoteExists() bool {
	if o == nil || IsNil(o.DeliveryNoteExists.Get()) {
		var ret bool
		return ret
	}
	return *o.DeliveryNoteExists.Get()
}

// GetDeliveryNoteExistsOk returns a tuple with the DeliveryNoteExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetDeliveryNoteExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryNoteExists.Get(), o.DeliveryNoteExists.IsSet()
}

// HasDeliveryNoteExists returns a boolean if a field has been set.
func (o *EmailData) HasDeliveryNoteExists() bool {
	if o != nil && o.DeliveryNoteExists.IsSet() {
		return true
	}

	return false
}

// SetDeliveryNoteExists gets a reference to the given NullableBool and assigns it to the DeliveryNoteExists field.
func (o *EmailData) SetDeliveryNoteExists(v bool) *EmailData {
	o.DeliveryNoteExists.Set(&v)
		return o
}
// SetDeliveryNoteExistsNil sets the value for DeliveryNoteExists to be an explicit nil
func (o *EmailData) SetDeliveryNoteExistsNil() *EmailData {
	o.DeliveryNoteExists.Set(nil)
	return o
}

// UnsetDeliveryNoteExists ensures that no value is present for DeliveryNoteExists, not even an explicit nil
func (o *EmailData) UnsetDeliveryNoteExists() {
	o.DeliveryNoteExists.Unset()
}

// GetAttachmentExists returns the AttachmentExists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetAttachmentExists() bool {
	if o == nil || IsNil(o.AttachmentExists.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachmentExists.Get()
}

// GetAttachmentExistsOk returns a tuple with the AttachmentExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetAttachmentExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentExists.Get(), o.AttachmentExists.IsSet()
}

// HasAttachmentExists returns a boolean if a field has been set.
func (o *EmailData) HasAttachmentExists() bool {
	if o != nil && o.AttachmentExists.IsSet() {
		return true
	}

	return false
}

// SetAttachmentExists gets a reference to the given NullableBool and assigns it to the AttachmentExists field.
func (o *EmailData) SetAttachmentExists(v bool) *EmailData {
	o.AttachmentExists.Set(&v)
		return o
}
// SetAttachmentExistsNil sets the value for AttachmentExists to be an explicit nil
func (o *EmailData) SetAttachmentExistsNil() *EmailData {
	o.AttachmentExists.Set(nil)
	return o
}

// UnsetAttachmentExists ensures that no value is present for AttachmentExists, not even an explicit nil
func (o *EmailData) UnsetAttachmentExists() {
	o.AttachmentExists.Unset()
}

// GetAccompanyingInvoiceExists returns the AccompanyingInvoiceExists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetAccompanyingInvoiceExists() bool {
	if o == nil || IsNil(o.AccompanyingInvoiceExists.Get()) {
		var ret bool
		return ret
	}
	return *o.AccompanyingInvoiceExists.Get()
}

// GetAccompanyingInvoiceExistsOk returns a tuple with the AccompanyingInvoiceExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetAccompanyingInvoiceExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccompanyingInvoiceExists.Get(), o.AccompanyingInvoiceExists.IsSet()
}

// HasAccompanyingInvoiceExists returns a boolean if a field has been set.
func (o *EmailData) HasAccompanyingInvoiceExists() bool {
	if o != nil && o.AccompanyingInvoiceExists.IsSet() {
		return true
	}

	return false
}

// SetAccompanyingInvoiceExists gets a reference to the given NullableBool and assigns it to the AccompanyingInvoiceExists field.
func (o *EmailData) SetAccompanyingInvoiceExists(v bool) *EmailData {
	o.AccompanyingInvoiceExists.Set(&v)
		return o
}
// SetAccompanyingInvoiceExistsNil sets the value for AccompanyingInvoiceExists to be an explicit nil
func (o *EmailData) SetAccompanyingInvoiceExistsNil() *EmailData {
	o.AccompanyingInvoiceExists.Set(nil)
	return o
}

// UnsetAccompanyingInvoiceExists ensures that no value is present for AccompanyingInvoiceExists, not even an explicit nil
func (o *EmailData) UnsetAccompanyingInvoiceExists() {
	o.AccompanyingInvoiceExists.Unset()
}

// GetDefaultAttachPdf returns the DefaultAttachPdf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailData) GetDefaultAttachPdf() bool {
	if o == nil || IsNil(o.DefaultAttachPdf.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultAttachPdf.Get()
}

// GetDefaultAttachPdfOk returns a tuple with the DefaultAttachPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailData) GetDefaultAttachPdfOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultAttachPdf.Get(), o.DefaultAttachPdf.IsSet()
}

// HasDefaultAttachPdf returns a boolean if a field has been set.
func (o *EmailData) HasDefaultAttachPdf() bool {
	if o != nil && o.DefaultAttachPdf.IsSet() {
		return true
	}

	return false
}

// SetDefaultAttachPdf gets a reference to the given NullableBool and assigns it to the DefaultAttachPdf field.
func (o *EmailData) SetDefaultAttachPdf(v bool) *EmailData {
	o.DefaultAttachPdf.Set(&v)
		return o
}
// SetDefaultAttachPdfNil sets the value for DefaultAttachPdf to be an explicit nil
func (o *EmailData) SetDefaultAttachPdfNil() *EmailData {
	o.DefaultAttachPdf.Set(nil)
	return o
}

// UnsetDefaultAttachPdf ensures that no value is present for DefaultAttachPdf, not even an explicit nil
func (o *EmailData) UnsetDefaultAttachPdf() {
	o.DefaultAttachPdf.Unset()
}

func (o EmailData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecipientEmail.IsSet() {
		toSerialize["recipient_email"] = o.RecipientEmail.Get()
	}
	if o.DefaultSenderEmail.IsSet() {
		toSerialize["default_sender_email"] = o.DefaultSenderEmail.Get()
	}
	if o.SenderEmailsList != nil {
		toSerialize["sender_emails_list"] = o.SenderEmailsList
	}
	if o.CcEmail.IsSet() {
		toSerialize["cc_email"] = o.CcEmail.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.DocumentExists.IsSet() {
		toSerialize["document_exists"] = o.DocumentExists.Get()
	}
	if o.DeliveryNoteExists.IsSet() {
		toSerialize["delivery_note_exists"] = o.DeliveryNoteExists.Get()
	}
	if o.AttachmentExists.IsSet() {
		toSerialize["attachment_exists"] = o.AttachmentExists.Get()
	}
	if o.AccompanyingInvoiceExists.IsSet() {
		toSerialize["accompanying_invoice_exists"] = o.AccompanyingInvoiceExists.Get()
	}
	if o.DefaultAttachPdf.IsSet() {
		toSerialize["default_attach_pdf"] = o.DefaultAttachPdf.Get()
	}
	return toSerialize, nil
}

type NullableEmailData struct {
	value *EmailData
	isSet bool
}

func (v NullableEmailData) Get() *EmailData {
	return v.value
}

func (v *NullableEmailData) Set(val *EmailData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailData(val *EmailData) *NullableEmailData {
	return &NullableEmailData{value: val, isSet: true}
}

func (v NullableEmailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


