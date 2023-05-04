/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
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
	Id *int32 `json:"id,omitempty"`
	Status *EmailStatus `json:"status,omitempty"`
	// Email sent date.
	SentDate *time.Time `json:"sent_date,omitempty"`
	// Errors count.
	ErrorsCount *int32 `json:"errors_count,omitempty"`
	// Error log.
	ErrorLog *string `json:"error_log,omitempty"`
	// Sender email.
	FromEmail *string `json:"from_email,omitempty"`
	// Sender name.
	FromName *string `json:"from_name,omitempty"`
	// Recipient email.
	ToEmail *string `json:"to_email,omitempty"`
	// Receipient email.
	ToName *string `json:"to_name,omitempty"`
	// Email subject.
	Subject *string `json:"subject,omitempty"`
	// Email content.
	Content *string `json:"content,omitempty"`
	CopyTo *string `json:"copy_to,omitempty"`
	RecipientStatus *EmailRecipientStatus `json:"recipient_status,omitempty"`
	RecipientDate *time.Time `json:"recipient_date,omitempty"`
	// Email kind.
	Kind *string `json:"kind,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *Email) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Email) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Email) SetId(v int32) *Email {
	o.Id = &v
	return o
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

// GetSentDate returns the SentDate field value if set, zero value otherwise.
func (o *Email) GetSentDate() time.Time {
	if o == nil || IsNil(o.SentDate) {
		var ret time.Time
		return ret
	}
	return *o.SentDate
}

// GetSentDateOk returns a tuple with the SentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetSentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentDate) {
		return nil, false
	}
	return o.SentDate, true
}

// HasSentDate returns a boolean if a field has been set.
func (o *Email) HasSentDate() bool {
	if o != nil && !IsNil(o.SentDate) {
		return true
	}

	return false
}

// SetSentDate gets a reference to the given time.Time and assigns it to the SentDate field.
func (o *Email) SetSentDate(v time.Time) *Email {
	o.SentDate = &v
	return o
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise.
func (o *Email) GetErrorsCount() int32 {
	if o == nil || IsNil(o.ErrorsCount) {
		var ret int32
		return ret
	}
	return *o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetErrorsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorsCount) {
		return nil, false
	}
	return o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *Email) HasErrorsCount() bool {
	if o != nil && !IsNil(o.ErrorsCount) {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given int32 and assigns it to the ErrorsCount field.
func (o *Email) SetErrorsCount(v int32) *Email {
	o.ErrorsCount = &v
	return o
}

// GetErrorLog returns the ErrorLog field value if set, zero value otherwise.
func (o *Email) GetErrorLog() string {
	if o == nil || IsNil(o.ErrorLog) {
		var ret string
		return ret
	}
	return *o.ErrorLog
}

// GetErrorLogOk returns a tuple with the ErrorLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetErrorLogOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorLog) {
		return nil, false
	}
	return o.ErrorLog, true
}

// HasErrorLog returns a boolean if a field has been set.
func (o *Email) HasErrorLog() bool {
	if o != nil && !IsNil(o.ErrorLog) {
		return true
	}

	return false
}

// SetErrorLog gets a reference to the given string and assigns it to the ErrorLog field.
func (o *Email) SetErrorLog(v string) *Email {
	o.ErrorLog = &v
	return o
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *Email) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetFromEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *Email) HasFromEmail() bool {
	if o != nil && !IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *Email) SetFromEmail(v string) *Email {
	o.FromEmail = &v
	return o
}

// GetFromName returns the FromName field value if set, zero value otherwise.
func (o *Email) GetFromName() string {
	if o == nil || IsNil(o.FromName) {
		var ret string
		return ret
	}
	return *o.FromName
}

// GetFromNameOk returns a tuple with the FromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetFromNameOk() (*string, bool) {
	if o == nil || IsNil(o.FromName) {
		return nil, false
	}
	return o.FromName, true
}

// HasFromName returns a boolean if a field has been set.
func (o *Email) HasFromName() bool {
	if o != nil && !IsNil(o.FromName) {
		return true
	}

	return false
}

// SetFromName gets a reference to the given string and assigns it to the FromName field.
func (o *Email) SetFromName(v string) *Email {
	o.FromName = &v
	return o
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise.
func (o *Email) GetToEmail() string {
	if o == nil || IsNil(o.ToEmail) {
		var ret string
		return ret
	}
	return *o.ToEmail
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetToEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ToEmail) {
		return nil, false
	}
	return o.ToEmail, true
}

// HasToEmail returns a boolean if a field has been set.
func (o *Email) HasToEmail() bool {
	if o != nil && !IsNil(o.ToEmail) {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given string and assigns it to the ToEmail field.
func (o *Email) SetToEmail(v string) *Email {
	o.ToEmail = &v
	return o
}

// GetToName returns the ToName field value if set, zero value otherwise.
func (o *Email) GetToName() string {
	if o == nil || IsNil(o.ToName) {
		var ret string
		return ret
	}
	return *o.ToName
}

// GetToNameOk returns a tuple with the ToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetToNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToName) {
		return nil, false
	}
	return o.ToName, true
}

// HasToName returns a boolean if a field has been set.
func (o *Email) HasToName() bool {
	if o != nil && !IsNil(o.ToName) {
		return true
	}

	return false
}

// SetToName gets a reference to the given string and assigns it to the ToName field.
func (o *Email) SetToName(v string) *Email {
	o.ToName = &v
	return o
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Email) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Email) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Email) SetSubject(v string) *Email {
	o.Subject = &v
	return o
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Email) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Email) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Email) SetContent(v string) *Email {
	o.Content = &v
	return o
}

// GetCopyTo returns the CopyTo field value if set, zero value otherwise.
func (o *Email) GetCopyTo() string {
	if o == nil || IsNil(o.CopyTo) {
		var ret string
		return ret
	}
	return *o.CopyTo
}

// GetCopyToOk returns a tuple with the CopyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetCopyToOk() (*string, bool) {
	if o == nil || IsNil(o.CopyTo) {
		return nil, false
	}
	return o.CopyTo, true
}

// HasCopyTo returns a boolean if a field has been set.
func (o *Email) HasCopyTo() bool {
	if o != nil && !IsNil(o.CopyTo) {
		return true
	}

	return false
}

// SetCopyTo gets a reference to the given string and assigns it to the CopyTo field.
func (o *Email) SetCopyTo(v string) *Email {
	o.CopyTo = &v
	return o
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

// GetRecipientDate returns the RecipientDate field value if set, zero value otherwise.
func (o *Email) GetRecipientDate() time.Time {
	if o == nil || IsNil(o.RecipientDate) {
		var ret time.Time
		return ret
	}
	return *o.RecipientDate
}

// GetRecipientDateOk returns a tuple with the RecipientDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetRecipientDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecipientDate) {
		return nil, false
	}
	return o.RecipientDate, true
}

// HasRecipientDate returns a boolean if a field has been set.
func (o *Email) HasRecipientDate() bool {
	if o != nil && !IsNil(o.RecipientDate) {
		return true
	}

	return false
}

// SetRecipientDate gets a reference to the given time.Time and assigns it to the RecipientDate field.
func (o *Email) SetRecipientDate(v time.Time) *Email {
	o.RecipientDate = &v
	return o
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Email) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Email) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Email) SetKind(v string) *Email {
	o.Kind = &v
	return o
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Email) GetAttachments() []EmailAttachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []EmailAttachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Email) GetAttachmentsOk() ([]EmailAttachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Email) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SentDate) {
		toSerialize["sent_date"] = o.SentDate
	}
	if !IsNil(o.ErrorsCount) {
		toSerialize["errors_count"] = o.ErrorsCount
	}
	if !IsNil(o.ErrorLog) {
		toSerialize["error_log"] = o.ErrorLog
	}
	if !IsNil(o.FromEmail) {
		toSerialize["from_email"] = o.FromEmail
	}
	if !IsNil(o.FromName) {
		toSerialize["from_name"] = o.FromName
	}
	if !IsNil(o.ToEmail) {
		toSerialize["to_email"] = o.ToEmail
	}
	if !IsNil(o.ToName) {
		toSerialize["to_name"] = o.ToName
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CopyTo) {
		toSerialize["copy_to"] = o.CopyTo
	}
	if !IsNil(o.RecipientStatus) {
		toSerialize["recipient_status"] = o.RecipientStatus
	}
	if !IsNil(o.RecipientDate) {
		toSerialize["recipient_date"] = o.RecipientDate
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Attachments) {
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


