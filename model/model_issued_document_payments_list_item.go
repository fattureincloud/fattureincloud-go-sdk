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

// checks if the IssuedDocumentPaymentsListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentPaymentsListItem{}

// IssuedDocumentPaymentsListItem struct for IssuedDocumentPaymentsListItem
type IssuedDocumentPaymentsListItem struct {
	// Issued document payment item id
	Id NullableInt32 `json:"id,omitempty"`
	// Issued document payment due date
	DueDate NullableString `json:"due_date,omitempty"`
	// Issued document payment amount
	Amount NullableFloat32 `json:"amount,omitempty"`
	Status *IssuedDocumentStatus `json:"status,omitempty"`
	PaymentAccount NullablePaymentAccount `json:"payment_account,omitempty"`
	// Issued document payment date [Only if status is paid]
	PaidDate NullableString `json:"paid_date,omitempty"`
	// Issued document payment advanced raw attributes for e-invoices
	EiRaw map[string]interface{} `json:"ei_raw,omitempty"`
	PaymentTerms *IssuedDocumentPaymentsListItemPaymentTerms `json:"payment_terms,omitempty"`
}

// NewIssuedDocumentPaymentsListItem instantiates a new IssuedDocumentPaymentsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentPaymentsListItem() *IssuedDocumentPaymentsListItem {
	this := IssuedDocumentPaymentsListItem{}
	return &this
}

// NewIssuedDocumentPaymentsListItemWithDefaults instantiates a new IssuedDocumentPaymentsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentPaymentsListItemWithDefaults() *IssuedDocumentPaymentsListItem {
	this := IssuedDocumentPaymentsListItem{}
	var status IssuedDocumentStatus = IssuedDocumentStatuses.NOT_PAID
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IssuedDocumentPaymentsListItem) SetId(v int32) *IssuedDocumentPaymentsListItem {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IssuedDocumentPaymentsListItem) SetIdNil() *IssuedDocumentPaymentsListItem {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IssuedDocumentPaymentsListItem) UnsetId() {
	o.Id.Unset()
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetDueDate() string {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret string
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableString and assigns it to the DueDate field.
func (o *IssuedDocumentPaymentsListItem) SetDueDate(v string) *IssuedDocumentPaymentsListItem {
	o.DueDate.Set(&v)
	return o
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *IssuedDocumentPaymentsListItem) SetDueDateNil() *IssuedDocumentPaymentsListItem {
	o.DueDate.Set(nil)
	return o
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *IssuedDocumentPaymentsListItem) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *IssuedDocumentPaymentsListItem) SetAmount(v float32) *IssuedDocumentPaymentsListItem {
	o.Amount.Set(&v)
	return o
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *IssuedDocumentPaymentsListItem) SetAmountNil() *IssuedDocumentPaymentsListItem {
	o.Amount.Set(nil)
	return o
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *IssuedDocumentPaymentsListItem) UnsetAmount() {
	o.Amount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssuedDocumentPaymentsListItem) GetStatus() IssuedDocumentStatus {
	if o == nil || IsNil(o.Status) {
		var ret IssuedDocumentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPaymentsListItem) GetStatusOk() (*IssuedDocumentStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IssuedDocumentStatus and assigns it to the Status field.
func (o *IssuedDocumentPaymentsListItem) SetStatus(v IssuedDocumentStatus) *IssuedDocumentPaymentsListItem {
	o.Status = &v
	return o
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetPaymentAccount() PaymentAccount {
	if o == nil || IsNil(o.PaymentAccount.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.PaymentAccount.Get()
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetPaymentAccountOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAccount.Get(), o.PaymentAccount.IsSet()
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given NullablePaymentAccount and assigns it to the PaymentAccount field.
func (o *IssuedDocumentPaymentsListItem) SetPaymentAccount(v PaymentAccount) *IssuedDocumentPaymentsListItem {
	o.PaymentAccount.Set(&v)
	return o
}
// SetPaymentAccountNil sets the value for PaymentAccount to be an explicit nil
func (o *IssuedDocumentPaymentsListItem) SetPaymentAccountNil() *IssuedDocumentPaymentsListItem {
	o.PaymentAccount.Set(nil)
	return o
}

// UnsetPaymentAccount ensures that no value is present for PaymentAccount, not even an explicit nil
func (o *IssuedDocumentPaymentsListItem) UnsetPaymentAccount() {
	o.PaymentAccount.Unset()
}

// GetPaidDate returns the PaidDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetPaidDate() string {
	if o == nil || IsNil(o.PaidDate.Get()) {
		var ret string
		return ret
	}
	return *o.PaidDate.Get()
}

// GetPaidDateOk returns a tuple with the PaidDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetPaidDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidDate.Get(), o.PaidDate.IsSet()
}

// HasPaidDate returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasPaidDate() bool {
	if o != nil && o.PaidDate.IsSet() {
		return true
	}

	return false
}

// SetPaidDate gets a reference to the given NullableString and assigns it to the PaidDate field.
func (o *IssuedDocumentPaymentsListItem) SetPaidDate(v string) *IssuedDocumentPaymentsListItem {
	o.PaidDate.Set(&v)
	return o
}
// SetPaidDateNil sets the value for PaidDate to be an explicit nil
func (o *IssuedDocumentPaymentsListItem) SetPaidDateNil() *IssuedDocumentPaymentsListItem {
	o.PaidDate.Set(nil)
	return o
}

// UnsetPaidDate ensures that no value is present for PaidDate, not even an explicit nil
func (o *IssuedDocumentPaymentsListItem) UnsetPaidDate() {
	o.PaidDate.Unset()
}

// GetEiRaw returns the EiRaw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentPaymentsListItem) GetEiRaw() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EiRaw
}

// GetEiRawOk returns a tuple with the EiRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentPaymentsListItem) GetEiRawOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EiRaw) {
		return map[string]interface{}{}, false
	}
	return o.EiRaw, true
}

// HasEiRaw returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasEiRaw() bool {
	if o != nil && IsNil(o.EiRaw) {
		return true
	}

	return false
}

// SetEiRaw gets a reference to the given map[string]interface{} and assigns it to the EiRaw field.
func (o *IssuedDocumentPaymentsListItem) SetEiRaw(v map[string]interface{}) *IssuedDocumentPaymentsListItem {
	o.EiRaw = v
	return o
}

// GetPaymentTerms returns the PaymentTerms field value if set, zero value otherwise.
func (o *IssuedDocumentPaymentsListItem) GetPaymentTerms() IssuedDocumentPaymentsListItemPaymentTerms {
	if o == nil || IsNil(o.PaymentTerms) {
		var ret IssuedDocumentPaymentsListItemPaymentTerms
		return ret
	}
	return *o.PaymentTerms
}

// GetPaymentTermsOk returns a tuple with the PaymentTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuedDocumentPaymentsListItem) GetPaymentTermsOk() (*IssuedDocumentPaymentsListItemPaymentTerms, bool) {
	if o == nil || IsNil(o.PaymentTerms) {
		return nil, false
	}
	return o.PaymentTerms, true
}

// HasPaymentTerms returns a boolean if a field has been set.
func (o *IssuedDocumentPaymentsListItem) HasPaymentTerms() bool {
	if o != nil && !IsNil(o.PaymentTerms) {
		return true
	}

	return false
}

// SetPaymentTerms gets a reference to the given IssuedDocumentPaymentsListItemPaymentTerms and assigns it to the PaymentTerms field.
func (o *IssuedDocumentPaymentsListItem) SetPaymentTerms(v IssuedDocumentPaymentsListItemPaymentTerms) *IssuedDocumentPaymentsListItem {
	o.PaymentTerms = &v
	return o
}

func (o IssuedDocumentPaymentsListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentPaymentsListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.PaymentAccount.IsSet() {
		toSerialize["payment_account"] = o.PaymentAccount.Get()
	}
	if o.PaidDate.IsSet() {
		toSerialize["paid_date"] = o.PaidDate.Get()
	}
	if o.EiRaw != nil {
		toSerialize["ei_raw"] = o.EiRaw
	}
	if !IsNil(o.PaymentTerms) {
		toSerialize["payment_terms"] = o.PaymentTerms
	}
	return toSerialize, nil
}

type NullableIssuedDocumentPaymentsListItem struct {
	value *IssuedDocumentPaymentsListItem
	isSet bool
}

func (v NullableIssuedDocumentPaymentsListItem) Get() *IssuedDocumentPaymentsListItem {
	return v.value
}

func (v *NullableIssuedDocumentPaymentsListItem) Set(val *IssuedDocumentPaymentsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentPaymentsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentPaymentsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentPaymentsListItem(val *IssuedDocumentPaymentsListItem) *NullableIssuedDocumentPaymentsListItem {
	return &NullableIssuedDocumentPaymentsListItem{value: val, isSet: true}
}

func (v NullableIssuedDocumentPaymentsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentPaymentsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


