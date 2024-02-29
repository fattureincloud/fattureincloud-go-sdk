/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.32
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the EmailSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSchedule{}

// EmailSchedule struct for EmailSchedule
type EmailSchedule struct {
	// Email sender id [required if **sender_email** is not specified]
	SenderId NullableInt32 `json:"sender_id,omitempty"`
	// Email sender address [required if **sender_id** is not specified]
	SenderEmail NullableString `json:"sender_email,omitempty"`
	// Email recipient emails [comma separated]
	RecipientEmail NullableString `json:"recipient_email,omitempty"`
	// Email subject
	Subject NullableString `json:"subject,omitempty"`
	// Email body [HTML Escaped] [max size 50KiB]
	Body NullableString `json:"body,omitempty"`
	Include NullableEmailScheduleInclude `json:"include,omitempty"`
	// Attach the pdf of the document
	AttachPdf NullableBool `json:"attach_pdf,omitempty"`
	// Send a copy of the email to the **cc_email** specified by **Get email data**
	SendCopy NullableBool `json:"send_copy,omitempty"`
}

// NewEmailSchedule instantiates a new EmailSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSchedule() *EmailSchedule {
	this := EmailSchedule{}
	return &this
}

// NewEmailScheduleWithDefaults instantiates a new EmailSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailScheduleWithDefaults() *EmailSchedule {
	this := EmailSchedule{}
	return &this
}

// GetSenderId returns the SenderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetSenderId() int32 {
	if o == nil || IsNil(o.SenderId.Get()) {
		var ret int32
		return ret
	}
	return *o.SenderId.Get()
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetSenderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderId.Get(), o.SenderId.IsSet()
}

// HasSenderId returns a boolean if a field has been set.
func (o *EmailSchedule) HasSenderId() bool {
	if o != nil && o.SenderId.IsSet() {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given NullableInt32 and assigns it to the SenderId field.
func (o *EmailSchedule) SetSenderId(v int32) *EmailSchedule {
	o.SenderId.Set(&v)
	return o
}
// SetSenderIdNil sets the value for SenderId to be an explicit nil
func (o *EmailSchedule) SetSenderIdNil() *EmailSchedule {
	o.SenderId.Set(nil)
	return o
}

// UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
func (o *EmailSchedule) UnsetSenderId() {
	o.SenderId.Unset()
}

// GetSenderEmail returns the SenderEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetSenderEmail() string {
	if o == nil || IsNil(o.SenderEmail.Get()) {
		var ret string
		return ret
	}
	return *o.SenderEmail.Get()
}

// GetSenderEmailOk returns a tuple with the SenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetSenderEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderEmail.Get(), o.SenderEmail.IsSet()
}

// HasSenderEmail returns a boolean if a field has been set.
func (o *EmailSchedule) HasSenderEmail() bool {
	if o != nil && o.SenderEmail.IsSet() {
		return true
	}

	return false
}

// SetSenderEmail gets a reference to the given NullableString and assigns it to the SenderEmail field.
func (o *EmailSchedule) SetSenderEmail(v string) *EmailSchedule {
	o.SenderEmail.Set(&v)
	return o
}
// SetSenderEmailNil sets the value for SenderEmail to be an explicit nil
func (o *EmailSchedule) SetSenderEmailNil() *EmailSchedule {
	o.SenderEmail.Set(nil)
	return o
}

// UnsetSenderEmail ensures that no value is present for SenderEmail, not even an explicit nil
func (o *EmailSchedule) UnsetSenderEmail() {
	o.SenderEmail.Unset()
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetRecipientEmail() string {
	if o == nil || IsNil(o.RecipientEmail.Get()) {
		var ret string
		return ret
	}
	return *o.RecipientEmail.Get()
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetRecipientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientEmail.Get(), o.RecipientEmail.IsSet()
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *EmailSchedule) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail.IsSet() {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given NullableString and assigns it to the RecipientEmail field.
func (o *EmailSchedule) SetRecipientEmail(v string) *EmailSchedule {
	o.RecipientEmail.Set(&v)
	return o
}
// SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil
func (o *EmailSchedule) SetRecipientEmailNil() *EmailSchedule {
	o.RecipientEmail.Set(nil)
	return o
}

// UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
func (o *EmailSchedule) UnsetRecipientEmail() {
	o.RecipientEmail.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailSchedule) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *EmailSchedule) SetSubject(v string) *EmailSchedule {
	o.Subject.Set(&v)
	return o
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *EmailSchedule) SetSubjectNil() *EmailSchedule {
	o.Subject.Set(nil)
	return o
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *EmailSchedule) UnsetSubject() {
	o.Subject.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *EmailSchedule) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *EmailSchedule) SetBody(v string) *EmailSchedule {
	o.Body.Set(&v)
	return o
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *EmailSchedule) SetBodyNil() *EmailSchedule {
	o.Body.Set(nil)
	return o
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *EmailSchedule) UnsetBody() {
	o.Body.Unset()
}

// GetInclude returns the Include field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetInclude() EmailScheduleInclude {
	if o == nil || IsNil(o.Include.Get()) {
		var ret EmailScheduleInclude
		return ret
	}
	return *o.Include.Get()
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetIncludeOk() (*EmailScheduleInclude, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include.Get(), o.Include.IsSet()
}

// HasInclude returns a boolean if a field has been set.
func (o *EmailSchedule) HasInclude() bool {
	if o != nil && o.Include.IsSet() {
		return true
	}

	return false
}

// SetInclude gets a reference to the given NullableEmailScheduleInclude and assigns it to the Include field.
func (o *EmailSchedule) SetInclude(v EmailScheduleInclude) *EmailSchedule {
	o.Include.Set(&v)
	return o
}
// SetIncludeNil sets the value for Include to be an explicit nil
func (o *EmailSchedule) SetIncludeNil() *EmailSchedule {
	o.Include.Set(nil)
	return o
}

// UnsetInclude ensures that no value is present for Include, not even an explicit nil
func (o *EmailSchedule) UnsetInclude() {
	o.Include.Unset()
}

// GetAttachPdf returns the AttachPdf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetAttachPdf() bool {
	if o == nil || IsNil(o.AttachPdf.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPdf.Get()
}

// GetAttachPdfOk returns a tuple with the AttachPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetAttachPdfOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPdf.Get(), o.AttachPdf.IsSet()
}

// HasAttachPdf returns a boolean if a field has been set.
func (o *EmailSchedule) HasAttachPdf() bool {
	if o != nil && o.AttachPdf.IsSet() {
		return true
	}

	return false
}

// SetAttachPdf gets a reference to the given NullableBool and assigns it to the AttachPdf field.
func (o *EmailSchedule) SetAttachPdf(v bool) *EmailSchedule {
	o.AttachPdf.Set(&v)
	return o
}
// SetAttachPdfNil sets the value for AttachPdf to be an explicit nil
func (o *EmailSchedule) SetAttachPdfNil() *EmailSchedule {
	o.AttachPdf.Set(nil)
	return o
}

// UnsetAttachPdf ensures that no value is present for AttachPdf, not even an explicit nil
func (o *EmailSchedule) UnsetAttachPdf() {
	o.AttachPdf.Unset()
}

// GetSendCopy returns the SendCopy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailSchedule) GetSendCopy() bool {
	if o == nil || IsNil(o.SendCopy.Get()) {
		var ret bool
		return ret
	}
	return *o.SendCopy.Get()
}

// GetSendCopyOk returns a tuple with the SendCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailSchedule) GetSendCopyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendCopy.Get(), o.SendCopy.IsSet()
}

// HasSendCopy returns a boolean if a field has been set.
func (o *EmailSchedule) HasSendCopy() bool {
	if o != nil && o.SendCopy.IsSet() {
		return true
	}

	return false
}

// SetSendCopy gets a reference to the given NullableBool and assigns it to the SendCopy field.
func (o *EmailSchedule) SetSendCopy(v bool) *EmailSchedule {
	o.SendCopy.Set(&v)
	return o
}
// SetSendCopyNil sets the value for SendCopy to be an explicit nil
func (o *EmailSchedule) SetSendCopyNil() *EmailSchedule {
	o.SendCopy.Set(nil)
	return o
}

// UnsetSendCopy ensures that no value is present for SendCopy, not even an explicit nil
func (o *EmailSchedule) UnsetSendCopy() {
	o.SendCopy.Unset()
}

func (o EmailSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SenderId.IsSet() {
		toSerialize["sender_id"] = o.SenderId.Get()
	}
	if o.SenderEmail.IsSet() {
		toSerialize["sender_email"] = o.SenderEmail.Get()
	}
	if o.RecipientEmail.IsSet() {
		toSerialize["recipient_email"] = o.RecipientEmail.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.Include.IsSet() {
		toSerialize["include"] = o.Include.Get()
	}
	if o.AttachPdf.IsSet() {
		toSerialize["attach_pdf"] = o.AttachPdf.Get()
	}
	if o.SendCopy.IsSet() {
		toSerialize["send_copy"] = o.SendCopy.Get()
	}
	return toSerialize, nil
}

type NullableEmailSchedule struct {
	value *EmailSchedule
	isSet bool
}

func (v NullableEmailSchedule) Get() *EmailSchedule {
	return v.value
}

func (v *NullableEmailSchedule) Set(val *EmailSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSchedule(val *EmailSchedule) *NullableEmailSchedule {
	return &NullableEmailSchedule{value: val, isSet: true}
}

func (v NullableEmailSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


