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
	"time"
)

// checks if the Email type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Email{}

// Email 
type Email struct {
	// Email unique identifier.
	Id NullableInt32 `json:"id,omitempty"`
	Status *EmailStatus `json:"status,omitempty"`
	// Email sent date.
	SentDate NullableTime `json:"sent_date,omitempty"`
	// Errors count.
	ErrorsCount NullableInt32 `json:"errors_count,omitempty"`
	// Error log.
	ErrorLog NullableString `json:"error_log,omitempty"`
	// Sender email.
	FromEmail NullableString `json:"from_email,omitempty"`
	// Sender name.
	FromName NullableString `json:"from_name,omitempty"`
	// Recipient email.
	ToEmail NullableString `json:"to_email,omitempty"`
	// Receipient email.
	ToName NullableString `json:"to_name,omitempty"`
	// Email subject.
	Subject NullableString `json:"subject,omitempty"`
	// Email content.
	Content NullableString `json:"content,omitempty"`
	CopyTo NullableString `json:"copy_to,omitempty"`
	RecipientStatus *EmailRecipientStatus `json:"recipient_status,omitempty"`
	RecipientDate NullableTime `json:"recipient_date,omitempty"`
	// Email kind.
	Kind NullableString `json:"kind,omitempty"`
	// Email attachments.
	Attachments []EmailAttachment `json:"attachments,omitempty"`
}

// NewEmail instantiates a new Email object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmail() *Email {
	this := Email{}
	return &this
}

// NewEmailWithDefaults instantiates a new Email object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailWithDefaults() *Email {
	this := Email{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Email) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Email) SetId(v int32) *Email {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Email) SetIdNil() *Email {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Email) UnsetId() {
	o.Id.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Email) GetStatus() EmailStatus {
	if o == nil || IsNil(o.Status) {
		var ret EmailStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetStatusOk() (*EmailStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Email) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmailStatus and assigns it to the Status field.
func (o *Email) SetStatus(v EmailStatus) *Email {
	o.Status = &v
	return o
}

// GetSentDate returns the SentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetSentDate() time.Time {
	if o == nil || IsNil(o.SentDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SentDate.Get()
}

// GetSentDateOk returns a tuple with the SentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetSentDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentDate.Get(), o.SentDate.IsSet()
}

// HasSentDate returns a boolean if a field has been set.
func (o *Email) HasSentDate() bool {
	if o != nil && o.SentDate.IsSet() {
		return true
	}

	return false
}

// SetSentDate gets a reference to the given NullableTime and assigns it to the SentDate field.
func (o *Email) SetSentDate(v time.Time) *Email {
	o.SentDate.Set(&v)
	return o
}
// SetSentDateNil sets the value for SentDate to be an explicit nil
func (o *Email) SetSentDateNil() *Email {
	o.SentDate.Set(nil)
	return o
}

// UnsetSentDate ensures that no value is present for SentDate, not even an explicit nil
func (o *Email) UnsetSentDate() {
	o.SentDate.Unset()
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetErrorsCount() int32 {
	if o == nil || IsNil(o.ErrorsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ErrorsCount.Get()
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetErrorsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorsCount.Get(), o.ErrorsCount.IsSet()
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *Email) HasErrorsCount() bool {
	if o != nil && o.ErrorsCount.IsSet() {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given NullableInt32 and assigns it to the ErrorsCount field.
func (o *Email) SetErrorsCount(v int32) *Email {
	o.ErrorsCount.Set(&v)
	return o
}
// SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil
func (o *Email) SetErrorsCountNil() *Email {
	o.ErrorsCount.Set(nil)
	return o
}

// UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
func (o *Email) UnsetErrorsCount() {
	o.ErrorsCount.Unset()
}

// GetErrorLog returns the ErrorLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetErrorLog() string {
	if o == nil || IsNil(o.ErrorLog.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorLog.Get()
}

// GetErrorLogOk returns a tuple with the ErrorLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetErrorLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorLog.Get(), o.ErrorLog.IsSet()
}

// HasErrorLog returns a boolean if a field has been set.
func (o *Email) HasErrorLog() bool {
	if o != nil && o.ErrorLog.IsSet() {
		return true
	}

	return false
}

// SetErrorLog gets a reference to the given NullableString and assigns it to the ErrorLog field.
func (o *Email) SetErrorLog(v string) *Email {
	o.ErrorLog.Set(&v)
	return o
}
// SetErrorLogNil sets the value for ErrorLog to be an explicit nil
func (o *Email) SetErrorLogNil() *Email {
	o.ErrorLog.Set(nil)
	return o
}

// UnsetErrorLog ensures that no value is present for ErrorLog, not even an explicit nil
func (o *Email) UnsetErrorLog() {
	o.ErrorLog.Unset()
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail.Get()) {
		var ret string
		return ret
	}
	return *o.FromEmail.Get()
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetFromEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromEmail.Get(), o.FromEmail.IsSet()
}

// HasFromEmail returns a boolean if a field has been set.
func (o *Email) HasFromEmail() bool {
	if o != nil && o.FromEmail.IsSet() {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given NullableString and assigns it to the FromEmail field.
func (o *Email) SetFromEmail(v string) *Email {
	o.FromEmail.Set(&v)
	return o
}
// SetFromEmailNil sets the value for FromEmail to be an explicit nil
func (o *Email) SetFromEmailNil() *Email {
	o.FromEmail.Set(nil)
	return o
}

// UnsetFromEmail ensures that no value is present for FromEmail, not even an explicit nil
func (o *Email) UnsetFromEmail() {
	o.FromEmail.Unset()
}

// GetFromName returns the FromName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetFromName() string {
	if o == nil || IsNil(o.FromName.Get()) {
		var ret string
		return ret
	}
	return *o.FromName.Get()
}

// GetFromNameOk returns a tuple with the FromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetFromNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromName.Get(), o.FromName.IsSet()
}

// HasFromName returns a boolean if a field has been set.
func (o *Email) HasFromName() bool {
	if o != nil && o.FromName.IsSet() {
		return true
	}

	return false
}

// SetFromName gets a reference to the given NullableString and assigns it to the FromName field.
func (o *Email) SetFromName(v string) *Email {
	o.FromName.Set(&v)
	return o
}
// SetFromNameNil sets the value for FromName to be an explicit nil
func (o *Email) SetFromNameNil() *Email {
	o.FromName.Set(nil)
	return o
}

// UnsetFromName ensures that no value is present for FromName, not even an explicit nil
func (o *Email) UnsetFromName() {
	o.FromName.Unset()
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetToEmail() string {
	if o == nil || IsNil(o.ToEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ToEmail.Get()
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetToEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToEmail.Get(), o.ToEmail.IsSet()
}

// HasToEmail returns a boolean if a field has been set.
func (o *Email) HasToEmail() bool {
	if o != nil && o.ToEmail.IsSet() {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given NullableString and assigns it to the ToEmail field.
func (o *Email) SetToEmail(v string) *Email {
	o.ToEmail.Set(&v)
	return o
}
// SetToEmailNil sets the value for ToEmail to be an explicit nil
func (o *Email) SetToEmailNil() *Email {
	o.ToEmail.Set(nil)
	return o
}

// UnsetToEmail ensures that no value is present for ToEmail, not even an explicit nil
func (o *Email) UnsetToEmail() {
	o.ToEmail.Unset()
}

// GetToName returns the ToName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetToName() string {
	if o == nil || IsNil(o.ToName.Get()) {
		var ret string
		return ret
	}
	return *o.ToName.Get()
}

// GetToNameOk returns a tuple with the ToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetToNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToName.Get(), o.ToName.IsSet()
}

// HasToName returns a boolean if a field has been set.
func (o *Email) HasToName() bool {
	if o != nil && o.ToName.IsSet() {
		return true
	}

	return false
}

// SetToName gets a reference to the given NullableString and assigns it to the ToName field.
func (o *Email) SetToName(v string) *Email {
	o.ToName.Set(&v)
	return o
}
// SetToNameNil sets the value for ToName to be an explicit nil
func (o *Email) SetToNameNil() *Email {
	o.ToName.Set(nil)
	return o
}

// UnsetToName ensures that no value is present for ToName, not even an explicit nil
func (o *Email) UnsetToName() {
	o.ToName.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *Email) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *Email) SetSubject(v string) *Email {
	o.Subject.Set(&v)
	return o
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *Email) SetSubjectNil() *Email {
	o.Subject.Set(nil)
	return o
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *Email) UnsetSubject() {
	o.Subject.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *Email) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *Email) SetContent(v string) *Email {
	o.Content.Set(&v)
	return o
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *Email) SetContentNil() *Email {
	o.Content.Set(nil)
	return o
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *Email) UnsetContent() {
	o.Content.Unset()
}

// GetCopyTo returns the CopyTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetCopyTo() string {
	if o == nil || IsNil(o.CopyTo.Get()) {
		var ret string
		return ret
	}
	return *o.CopyTo.Get()
}

// GetCopyToOk returns a tuple with the CopyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetCopyToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyTo.Get(), o.CopyTo.IsSet()
}

// HasCopyTo returns a boolean if a field has been set.
func (o *Email) HasCopyTo() bool {
	if o != nil && o.CopyTo.IsSet() {
		return true
	}

	return false
}

// SetCopyTo gets a reference to the given NullableString and assigns it to the CopyTo field.
func (o *Email) SetCopyTo(v string) *Email {
	o.CopyTo.Set(&v)
	return o
}
// SetCopyToNil sets the value for CopyTo to be an explicit nil
func (o *Email) SetCopyToNil() *Email {
	o.CopyTo.Set(nil)
	return o
}

// UnsetCopyTo ensures that no value is present for CopyTo, not even an explicit nil
func (o *Email) UnsetCopyTo() {
	o.CopyTo.Unset()
}

// GetRecipientStatus returns the RecipientStatus field value if set, zero value otherwise.
func (o *Email) GetRecipientStatus() EmailRecipientStatus {
	if o == nil || IsNil(o.RecipientStatus) {
		var ret EmailRecipientStatus
		return ret
	}
	return *o.RecipientStatus
}

// GetRecipientStatusOk returns a tuple with the RecipientStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetRecipientStatusOk() (*EmailRecipientStatus, bool) {
	if o == nil || IsNil(o.RecipientStatus) {
		return nil, false
	}
	return o.RecipientStatus, true
}

// HasRecipientStatus returns a boolean if a field has been set.
func (o *Email) HasRecipientStatus() bool {
	if o != nil && !IsNil(o.RecipientStatus) {
		return true
	}

	return false
}

// SetRecipientStatus gets a reference to the given EmailRecipientStatus and assigns it to the RecipientStatus field.
func (o *Email) SetRecipientStatus(v EmailRecipientStatus) *Email {
	o.RecipientStatus = &v
	return o
}

// GetRecipientDate returns the RecipientDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetRecipientDate() time.Time {
	if o == nil || IsNil(o.RecipientDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecipientDate.Get()
}

// GetRecipientDateOk returns a tuple with the RecipientDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetRecipientDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientDate.Get(), o.RecipientDate.IsSet()
}

// HasRecipientDate returns a boolean if a field has been set.
func (o *Email) HasRecipientDate() bool {
	if o != nil && o.RecipientDate.IsSet() {
		return true
	}

	return false
}

// SetRecipientDate gets a reference to the given NullableTime and assigns it to the RecipientDate field.
func (o *Email) SetRecipientDate(v time.Time) *Email {
	o.RecipientDate.Set(&v)
	return o
}
// SetRecipientDateNil sets the value for RecipientDate to be an explicit nil
func (o *Email) SetRecipientDateNil() *Email {
	o.RecipientDate.Set(nil)
	return o
}

// UnsetRecipientDate ensures that no value is present for RecipientDate, not even an explicit nil
func (o *Email) UnsetRecipientDate() {
	o.RecipientDate.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetKind() string {
	if o == nil || IsNil(o.Kind.Get()) {
		var ret string
		return ret
	}
	return *o.Kind.Get()
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kind.Get(), o.Kind.IsSet()
}

// HasKind returns a boolean if a field has been set.
func (o *Email) HasKind() bool {
	if o != nil && o.Kind.IsSet() {
		return true
	}

	return false
}

// SetKind gets a reference to the given NullableString and assigns it to the Kind field.
func (o *Email) SetKind(v string) *Email {
	o.Kind.Set(&v)
	return o
}
// SetKindNil sets the value for Kind to be an explicit nil
func (o *Email) SetKindNil() *Email {
	o.Kind.Set(nil)
	return o
}

// UnsetKind ensures that no value is present for Kind, not even an explicit nil
func (o *Email) UnsetKind() {
	o.Kind.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Email) GetAttachments() []EmailAttachment {
	if o == nil {
		var ret []EmailAttachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetAttachmentsOk() ([]EmailAttachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Email) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []EmailAttachment and assigns it to the Attachments field.
func (o *Email) SetAttachments(v []EmailAttachment) *Email {
	o.Attachments = v
	return o
}

func (o Email) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Email) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.SentDate.IsSet() {
		toSerialize["sent_date"] = o.SentDate.Get()
	}
	if o.ErrorsCount.IsSet() {
		toSerialize["errors_count"] = o.ErrorsCount.Get()
	}
	if o.ErrorLog.IsSet() {
		toSerialize["error_log"] = o.ErrorLog.Get()
	}
	if o.FromEmail.IsSet() {
		toSerialize["from_email"] = o.FromEmail.Get()
	}
	if o.FromName.IsSet() {
		toSerialize["from_name"] = o.FromName.Get()
	}
	if o.ToEmail.IsSet() {
		toSerialize["to_email"] = o.ToEmail.Get()
	}
	if o.ToName.IsSet() {
		toSerialize["to_name"] = o.ToName.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.CopyTo.IsSet() {
		toSerialize["copy_to"] = o.CopyTo.Get()
	}
	if !IsNil(o.RecipientStatus) {
		toSerialize["recipient_status"] = o.RecipientStatus
	}
	if o.RecipientDate.IsSet() {
		toSerialize["recipient_date"] = o.RecipientDate.Get()
	}
	if o.Kind.IsSet() {
		toSerialize["kind"] = o.Kind.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableEmail struct {
	value *Email
	isSet bool
}

func (v NullableEmail) Get() *Email {
	return v.value
}

func (v *NullableEmail) Set(val *Email) {
	v.value = val
	v.isSet = true
}

func (v NullableEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmail(val *Email) *NullableEmail {
	return &NullableEmail{value: val, isSet: true}
}

func (v NullableEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


