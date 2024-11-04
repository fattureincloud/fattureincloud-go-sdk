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

// checks if the F24 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &F24{}

// F24 struct for F24
type F24 struct {
	// F24 id
Id NullableInt32 `json:"id,omitempty"`
	// F24 due date
DueDate NullableString `json:"due_date,omitempty"`
Status *F24Status `json:"status,omitempty"`
PaymentAccount NullablePaymentAccount `json:"payment_account,omitempty"`
	// F24 amount
Amount NullableFloat32 `json:"amount,omitempty"`
	// [Temporary] [Read Only] F24 url of the attached file
AttachmentUrl NullableString `json:"attachment_url,omitempty"`
	// [Write Only]  F24 attachment token returned by POST /taxes/attachment
AttachmentToken NullableString `json:"attachment_token,omitempty"`
	// F24 description
Description NullableString `json:"description,omitempty"`
}

// NewF24 instantiates a new F24 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewF24() *F24 {
	this := F24{}
	return &this
}

// NewF24WithDefaults instantiates a new F24 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewF24WithDefaults() *F24 {
	this := F24{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *F24) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *F24) SetId(v int32) *F24 {
	o.Id.Set(&v)
		return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *F24) SetIdNil() *F24 {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *F24) UnsetId() {
	o.Id.Unset()
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetDueDate() string {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret string
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *F24) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableString and assigns it to the DueDate field.
func (o *F24) SetDueDate(v string) *F24 {
	o.DueDate.Set(&v)
		return o
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *F24) SetDueDateNil() *F24 {
	o.DueDate.Set(nil)
	return o
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *F24) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *F24) GetStatus() F24Status {
	if o == nil || IsNil(o.Status) {
		var ret F24Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *F24) GetStatusOk() (*F24Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *F24) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given F24Status and assigns it to the Status field.
func (o *F24) SetStatus(v F24Status) *F24 {
	o.Status = &v
		return o
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetPaymentAccount() PaymentAccount {
	if o == nil || IsNil(o.PaymentAccount.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.PaymentAccount.Get()
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetPaymentAccountOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAccount.Get(), o.PaymentAccount.IsSet()
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *F24) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given NullablePaymentAccount and assigns it to the PaymentAccount field.
func (o *F24) SetPaymentAccount(v PaymentAccount) *F24 {
	o.PaymentAccount.Set(&v)
		return o
}
// SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil
func (o *F24) SetPaymentAccountNil() *F24 {
	o.PaymentAccount.Set(nil)
	return o
}

// UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil
func (o *F24) UnsetPaymentAccount() {
	o.PaymentAccount.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *F24) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *F24) SetAmount(v float32) *F24 {
	o.Amount.Set(&v)
		return o
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *F24) SetAmountNil() *F24 {
	o.Amount.Set(nil)
	return o
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *F24) UnsetAmount() {
	o.Amount.Unset()
}

// GetAttachmentUrl returns the AttachmentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetAttachmentUrl() string {
	if o == nil || IsNil(o.AttachmentUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentUrl.Get()
}

// GetAttachmentUrlOk returns a tuple with the AttachmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetAttachmentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentUrl.Get(), o.AttachmentUrl.IsSet()
}

// HasAttachmentUrl returns a boolean if a field has been set.
func (o *F24) HasAttachmentUrl() bool {
	if o != nil && o.AttachmentUrl.IsSet() {
		return true
	}

	return false
}

// SetAttachmentUrl gets a reference to the given NullableString and assigns it to the AttachmentUrl field.
func (o *F24) SetAttachmentUrl(v string) *F24 {
	o.AttachmentUrl.Set(&v)
		return o
}
// SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil
func (o *F24) SetAttachmentUrlNil() *F24 {
	o.AttachmentUrl.Set(nil)
	return o
}

// UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
func (o *F24) UnsetAttachmentUrl() {
	o.AttachmentUrl.Unset()
}

// GetAttachmentToken returns the AttachmentToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetAttachmentToken() string {
	if o == nil || IsNil(o.AttachmentToken.Get()) {
		var ret string
		return ret
	}
	return *o.AttachmentToken.Get()
}

// GetAttachmentTokenOk returns a tuple with the AttachmentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetAttachmentTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentToken.Get(), o.AttachmentToken.IsSet()
}

// HasAttachmentToken returns a boolean if a field has been set.
func (o *F24) HasAttachmentToken() bool {
	if o != nil && o.AttachmentToken.IsSet() {
		return true
	}

	return false
}

// SetAttachmentToken gets a reference to the given NullableString and assigns it to the AttachmentToken field.
func (o *F24) SetAttachmentToken(v string) *F24 {
	o.AttachmentToken.Set(&v)
		return o
}
// SetAttachmentTokenNil sets the value for AttachmentToken to be an explicit nil
func (o *F24) SetAttachmentTokenNil() *F24 {
	o.AttachmentToken.Set(nil)
	return o
}

// UnsetAttachmentToken ensures that no value is present for AttachmentToken, not even an explicit nil
func (o *F24) UnsetAttachmentToken() {
	o.AttachmentToken.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *F24) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *F24) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *F24) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *F24) SetDescription(v string) *F24 {
	o.Description.Set(&v)
		return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *F24) SetDescriptionNil() *F24 {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *F24) UnsetDescription() {
	o.Description.Unset()
}

func (o F24) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o F24) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.PaymentAccount.IsSet() {
		toSerialize["payment_account"] = o.PaymentAccount.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.AttachmentUrl.IsSet() {
		toSerialize["attachment_url"] = o.AttachmentUrl.Get()
	}
	if o.AttachmentToken.IsSet() {
		toSerialize["attachment_token"] = o.AttachmentToken.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableF24 struct {
	value *F24
	isSet bool
}

func (v NullableF24) Get() *F24 {
	return v.value
}

func (v *NullableF24) Set(val *F24) {
	v.value = val
	v.isSet = true
}

func (v NullableF24) IsSet() bool {
	return v.isSet
}

func (v *NullableF24) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableF24(val *F24) *NullableF24 {
	return &NullableF24{value: val, isSet: true}
}

func (v NullableF24) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableF24) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


