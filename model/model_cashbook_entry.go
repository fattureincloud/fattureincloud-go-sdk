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

// checks if the CashbookEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashbookEntry{}

// CashbookEntry struct for CashbookEntry
type CashbookEntry struct {
	// Cashbook id
	Id NullableString `json:"id,omitempty"`
	// Cashbook date
	Date NullableString `json:"date,omitempty"`
	// Cashbook description
	Description NullableString `json:"description,omitempty"`
	Kind *CashbookEntryKind `json:"kind,omitempty"`
	Type NullableCashbookEntryType `json:"type,omitempty"`
	// Cashbook entity name
	EntityName NullableString `json:"entity_name,omitempty"`
	Document NullableCashbookEntryDocument `json:"document,omitempty"`
	// [Only for cashbook entry in] Cashbook total amount in
	AmountIn NullableFloat32 `json:"amount_in,omitempty"`
	PaymentAccountIn NullablePaymentAccount `json:"payment_account_in,omitempty"`
	// [Only for cashbook entry out] Cashbook total amount out
	AmountOut NullableFloat32 `json:"amount_out,omitempty"`
	PaymentAccountOut NullablePaymentAccount `json:"payment_account_out,omitempty"`
}

// NewCashbookEntry instantiates a new CashbookEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashbookEntry() *CashbookEntry {
	this := CashbookEntry{}
	return &this
}

// NewCashbookEntryWithDefaults instantiates a new CashbookEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashbookEntryWithDefaults() *CashbookEntry {
	this := CashbookEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CashbookEntry) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CashbookEntry) SetId(v string) *CashbookEntry {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CashbookEntry) SetIdNil() *CashbookEntry {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CashbookEntry) UnsetId() {
	o.Id.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *CashbookEntry) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *CashbookEntry) SetDate(v string) *CashbookEntry {
	o.Date.Set(&v)
	return o
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *CashbookEntry) SetDateNil() *CashbookEntry {
	o.Date.Set(nil)
	return o
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *CashbookEntry) UnsetDate() {
	o.Date.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CashbookEntry) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CashbookEntry) SetDescription(v string) *CashbookEntry {
	o.Description.Set(&v)
	return o
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CashbookEntry) SetDescriptionNil() *CashbookEntry {
	o.Description.Set(nil)
	return o
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CashbookEntry) UnsetDescription() {
	o.Description.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CashbookEntry) GetKind() CashbookEntryKind {
	if o == nil || IsNil(o.Kind) {
		var ret CashbookEntryKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashbookEntry) GetKindOk() (*CashbookEntryKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CashbookEntry) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CashbookEntryKind and assigns it to the Kind field.
func (o *CashbookEntry) SetKind(v CashbookEntryKind) *CashbookEntry {
	o.Kind = &v
	return o
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetType() CashbookEntryType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret CashbookEntryType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetTypeOk() (*CashbookEntryType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CashbookEntry) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableCashbookEntryType and assigns it to the Type field.
func (o *CashbookEntry) SetType(v CashbookEntryType) *CashbookEntry {
	o.Type.Set(&v)
	return o
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CashbookEntry) SetTypeNil() *CashbookEntry {
	o.Type.Set(nil)
	return o
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CashbookEntry) UnsetType() {
	o.Type.Unset()
}

// GetEntityName returns the EntityName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetEntityName() string {
	if o == nil || IsNil(o.EntityName.Get()) {
		var ret string
		return ret
	}
	return *o.EntityName.Get()
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetEntityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityName.Get(), o.EntityName.IsSet()
}

// HasEntityName returns a boolean if a field has been set.
func (o *CashbookEntry) HasEntityName() bool {
	if o != nil && o.EntityName.IsSet() {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given NullableString and assigns it to the EntityName field.
func (o *CashbookEntry) SetEntityName(v string) *CashbookEntry {
	o.EntityName.Set(&v)
	return o
}
// SetEntityNameNil sets the value for EntityName to be an explicit nil
func (o *CashbookEntry) SetEntityNameNil() *CashbookEntry {
	o.EntityName.Set(nil)
	return o
}

// UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
func (o *CashbookEntry) UnsetEntityName() {
	o.EntityName.Unset()
}

// GetDocument returns the Document field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetDocument() CashbookEntryDocument {
	if o == nil || IsNil(o.Document.Get()) {
		var ret CashbookEntryDocument
		return ret
	}
	return *o.Document.Get()
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetDocumentOk() (*CashbookEntryDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Document.Get(), o.Document.IsSet()
}

// HasDocument returns a boolean if a field has been set.
func (o *CashbookEntry) HasDocument() bool {
	if o != nil && o.Document.IsSet() {
		return true
	}

	return false
}

// SetDocument gets a reference to the given NullableCashbookEntryDocument and assigns it to the Document field.
func (o *CashbookEntry) SetDocument(v CashbookEntryDocument) *CashbookEntry {
	o.Document.Set(&v)
	return o
}
// SetDocumentNil sets the value for Document to be an explicit nil
func (o *CashbookEntry) SetDocumentNil() *CashbookEntry {
	o.Document.Set(nil)
	return o
}

// UnsetDocument ensures that no value is present for Document, not even an explicit nil
func (o *CashbookEntry) UnsetDocument() {
	o.Document.Unset()
}

// GetAmountIn returns the AmountIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetAmountIn() float32 {
	if o == nil || IsNil(o.AmountIn.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountIn.Get()
}

// GetAmountInOk returns a tuple with the AmountIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetAmountInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountIn.Get(), o.AmountIn.IsSet()
}

// HasAmountIn returns a boolean if a field has been set.
func (o *CashbookEntry) HasAmountIn() bool {
	if o != nil && o.AmountIn.IsSet() {
		return true
	}

	return false
}

// SetAmountIn gets a reference to the given NullableFloat32 and assigns it to the AmountIn field.
func (o *CashbookEntry) SetAmountIn(v float32) *CashbookEntry {
	o.AmountIn.Set(&v)
	return o
}
// SetAmountInNil sets the value for AmountIn to be an explicit nil
func (o *CashbookEntry) SetAmountInNil() *CashbookEntry {
	o.AmountIn.Set(nil)
	return o
}

// UnsetAmountIn ensures that no value is present for AmountIn, not even an explicit nil
func (o *CashbookEntry) UnsetAmountIn() {
	o.AmountIn.Unset()
}

// GetPaymentAccountIn returns the PaymentAccountIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetPaymentAccountIn() PaymentAccount {
	if o == nil || IsNil(o.PaymentAccountIn.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.PaymentAccountIn.Get()
}

// GetPaymentAccountInOk returns a tuple with the PaymentAccountIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetPaymentAccountInOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAccountIn.Get(), o.PaymentAccountIn.IsSet()
}

// HasPaymentAccountIn returns a boolean if a field has been set.
func (o *CashbookEntry) HasPaymentAccountIn() bool {
	if o != nil && o.PaymentAccountIn.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccountIn gets a reference to the given NullablePaymentAccount and assigns it to the PaymentAccountIn field.
func (o *CashbookEntry) SetPaymentAccountIn(v PaymentAccount) *CashbookEntry {
	o.PaymentAccountIn.Set(&v)
	return o
}
// SetPaymentAccountInNil sets the value for PaymentAccountIn to be an explicit nil
func (o *CashbookEntry) SetPaymentAccountInNil() *CashbookEntry {
	o.PaymentAccountIn.Set(nil)
	return o
}

// UnsetPaymentAccountIn ensures that no value is present for PaymentAccountIn, not even an explicit nil
func (o *CashbookEntry) UnsetPaymentAccountIn() {
	o.PaymentAccountIn.Unset()
}

// GetAmountOut returns the AmountOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetAmountOut() float32 {
	if o == nil || IsNil(o.AmountOut.Get()) {
		var ret float32
		return ret
	}
	return *o.AmountOut.Get()
}

// GetAmountOutOk returns a tuple with the AmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetAmountOutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountOut.Get(), o.AmountOut.IsSet()
}

// HasAmountOut returns a boolean if a field has been set.
func (o *CashbookEntry) HasAmountOut() bool {
	if o != nil && o.AmountOut.IsSet() {
		return true
	}

	return false
}

// SetAmountOut gets a reference to the given NullableFloat32 and assigns it to the AmountOut field.
func (o *CashbookEntry) SetAmountOut(v float32) *CashbookEntry {
	o.AmountOut.Set(&v)
	return o
}
// SetAmountOutNil sets the value for AmountOut to be an explicit nil
func (o *CashbookEntry) SetAmountOutNil() *CashbookEntry {
	o.AmountOut.Set(nil)
	return o
}

// UnsetAmountOut ensures that no value is present for AmountOut, not even an explicit nil
func (o *CashbookEntry) UnsetAmountOut() {
	o.AmountOut.Unset()
}

// GetPaymentAccountOut returns the PaymentAccountOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CashbookEntry) GetPaymentAccountOut() PaymentAccount {
	if o == nil || IsNil(o.PaymentAccountOut.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.PaymentAccountOut.Get()
}

// GetPaymentAccountOutOk returns a tuple with the PaymentAccountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashbookEntry) GetPaymentAccountOutOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentAccountOut.Get(), o.PaymentAccountOut.IsSet()
}

// HasPaymentAccountOut returns a boolean if a field has been set.
func (o *CashbookEntry) HasPaymentAccountOut() bool {
	if o != nil && o.PaymentAccountOut.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccountOut gets a reference to the given NullablePaymentAccount and assigns it to the PaymentAccountOut field.
func (o *CashbookEntry) SetPaymentAccountOut(v PaymentAccount) *CashbookEntry {
	o.PaymentAccountOut.Set(&v)
	return o
}
// SetPaymentAccountOutNil sets the value for PaymentAccountOut to be an explicit nil
func (o *CashbookEntry) SetPaymentAccountOutNil() *CashbookEntry {
	o.PaymentAccountOut.Set(nil)
	return o
}

// UnsetPaymentAccountOut ensures that no value is present for PaymentAccountOut, not even an explicit nil
func (o *CashbookEntry) UnsetPaymentAccountOut() {
	o.PaymentAccountOut.Unset()
}

func (o CashbookEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashbookEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.EntityName.IsSet() {
		toSerialize["entity_name"] = o.EntityName.Get()
	}
	if o.Document.IsSet() {
		toSerialize["document"] = o.Document.Get()
	}
	if o.AmountIn.IsSet() {
		toSerialize["amount_in"] = o.AmountIn.Get()
	}
	if o.PaymentAccountIn.IsSet() {
		toSerialize["payment_account_in"] = o.PaymentAccountIn.Get()
	}
	if o.AmountOut.IsSet() {
		toSerialize["amount_out"] = o.AmountOut.Get()
	}
	if o.PaymentAccountOut.IsSet() {
		toSerialize["payment_account_out"] = o.PaymentAccountOut.Get()
	}
	return toSerialize, nil
}

type NullableCashbookEntry struct {
	value *CashbookEntry
	isSet bool
}

func (v NullableCashbookEntry) Get() *CashbookEntry {
	return v.value
}

func (v *NullableCashbookEntry) Set(val *CashbookEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCashbookEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCashbookEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashbookEntry(val *CashbookEntry) *NullableCashbookEntry {
	return &NullableCashbookEntry{value: val, isSet: true}
}

func (v NullableCashbookEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashbookEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


