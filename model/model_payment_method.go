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

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	// Unique identifier
	Id NullableInt32 `json:"id,omitempty"`
	// Name of the payment method
	Name NullableString `json:"name,omitempty"`
	Type *PaymentMethodType `json:"type,omitempty"`
	// Determines if this is the default payment method.
	IsDefault NullableBool `json:"is_default,omitempty"`
	DefaultPaymentAccount NullablePaymentAccount `json:"default_payment_account,omitempty"`
	// Method details rows
	Details []PaymentMethodDetails `json:"details,omitempty"`
	// Bank iban
	BankIban NullableString `json:"bank_iban,omitempty"`
	// Bank name
	BankName NullableString `json:"bank_name,omitempty"`
	// Bank beneficiary
	BankBeneficiary NullableString `json:"bank_beneficiary,omitempty"`
	// E-invoice payment method
	EiPaymentMethod NullableString `json:"ei_payment_method,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	var type_ PaymentMethodType = PaymentMethodTypes.STANDARD
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethod) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *PaymentMethod) SetId(v int32) *PaymentMethod {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PaymentMethod) SetIdNil() *PaymentMethod {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PaymentMethod) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethod) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PaymentMethod) SetName(v string) *PaymentMethod {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PaymentMethod) SetNameNil() *PaymentMethod {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PaymentMethod) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() PaymentMethodType {
	if o == nil || IsNil(o.Type) {
		var ret PaymentMethodType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentMethodType and assigns it to the Type field.
func (o *PaymentMethod) SetType(v PaymentMethodType) *PaymentMethod {
	o.Type = &v
	return o
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *PaymentMethod) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *PaymentMethod) SetIsDefault(v bool) *PaymentMethod {
	o.IsDefault.Set(&v)
	return o
}
// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *PaymentMethod) SetIsDefaultNil() *PaymentMethod {
	o.IsDefault.Set(nil)
	return o
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *PaymentMethod) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetDefaultPaymentAccount returns the DefaultPaymentAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetDefaultPaymentAccount() PaymentAccount {
	if o == nil || IsNil(o.DefaultPaymentAccount.Get()) {
		var ret PaymentAccount
		return ret
	}
	return *o.DefaultPaymentAccount.Get()
}

// GetDefaultPaymentAccountOk returns a tuple with the DefaultPaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetDefaultPaymentAccountOk() (*PaymentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPaymentAccount.Get(), o.DefaultPaymentAccount.IsSet()
}

// HasDefaultPaymentAccount returns a boolean if a field has been set.
func (o *PaymentMethod) HasDefaultPaymentAccount() bool {
	if o != nil && o.DefaultPaymentAccount.IsSet() {
		return true
	}

	return false
}

// SetDefaultPaymentAccount gets a reference to the given NullablePaymentAccount and assigns it to the DefaultPaymentAccount field.
func (o *PaymentMethod) SetDefaultPaymentAccount(v PaymentAccount) *PaymentMethod {
	o.DefaultPaymentAccount.Set(&v)
	return o
}
// SetDefaultPaymentAccountNil sets the value for DefaultPaymentAccount to be an explicit nil
func (o *PaymentMethod) SetDefaultPaymentAccountNil() *PaymentMethod {
	o.DefaultPaymentAccount.Set(nil)
	return o
}

// UnsetDefaultPaymentAccount ensures that no value is present for DefaultPaymentAccount, not even an explicit nil
func (o *PaymentMethod) UnsetDefaultPaymentAccount() {
	o.DefaultPaymentAccount.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetDetails() []PaymentMethodDetails {
	if o == nil {
		var ret []PaymentMethodDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetDetailsOk() ([]PaymentMethodDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PaymentMethod) HasDetails() bool {
	if o != nil && IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []PaymentMethodDetails and assigns it to the Details field.
func (o *PaymentMethod) SetDetails(v []PaymentMethodDetails) *PaymentMethod {
	o.Details = v
	return o
}

// GetBankIban returns the BankIban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetBankIban() string {
	if o == nil || IsNil(o.BankIban.Get()) {
		var ret string
		return ret
	}
	return *o.BankIban.Get()
}

// GetBankIbanOk returns a tuple with the BankIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetBankIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankIban.Get(), o.BankIban.IsSet()
}

// HasBankIban returns a boolean if a field has been set.
func (o *PaymentMethod) HasBankIban() bool {
	if o != nil && o.BankIban.IsSet() {
		return true
	}

	return false
}

// SetBankIban gets a reference to the given NullableString and assigns it to the BankIban field.
func (o *PaymentMethod) SetBankIban(v string) *PaymentMethod {
	o.BankIban.Set(&v)
	return o
}
// SetBankIbanNil sets the value for BankIban to be an explicit nil
func (o *PaymentMethod) SetBankIbanNil() *PaymentMethod {
	o.BankIban.Set(nil)
	return o
}

// UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
func (o *PaymentMethod) UnsetBankIban() {
	o.BankIban.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetBankName() string {
	if o == nil || IsNil(o.BankName.Get()) {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *PaymentMethod) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *PaymentMethod) SetBankName(v string) *PaymentMethod {
	o.BankName.Set(&v)
	return o
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *PaymentMethod) SetBankNameNil() *PaymentMethod {
	o.BankName.Set(nil)
	return o
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *PaymentMethod) UnsetBankName() {
	o.BankName.Unset()
}

// GetBankBeneficiary returns the BankBeneficiary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetBankBeneficiary() string {
	if o == nil || IsNil(o.BankBeneficiary.Get()) {
		var ret string
		return ret
	}
	return *o.BankBeneficiary.Get()
}

// GetBankBeneficiaryOk returns a tuple with the BankBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetBankBeneficiaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankBeneficiary.Get(), o.BankBeneficiary.IsSet()
}

// HasBankBeneficiary returns a boolean if a field has been set.
func (o *PaymentMethod) HasBankBeneficiary() bool {
	if o != nil && o.BankBeneficiary.IsSet() {
		return true
	}

	return false
}

// SetBankBeneficiary gets a reference to the given NullableString and assigns it to the BankBeneficiary field.
func (o *PaymentMethod) SetBankBeneficiary(v string) *PaymentMethod {
	o.BankBeneficiary.Set(&v)
	return o
}
// SetBankBeneficiaryNil sets the value for BankBeneficiary to be an explicit nil
func (o *PaymentMethod) SetBankBeneficiaryNil() *PaymentMethod {
	o.BankBeneficiary.Set(nil)
	return o
}

// UnsetBankBeneficiary ensures that no value is present for BankBeneficiary, not even an explicit nil
func (o *PaymentMethod) UnsetBankBeneficiary() {
	o.BankBeneficiary.Unset()
}

// GetEiPaymentMethod returns the EiPaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetEiPaymentMethod() string {
	if o == nil || IsNil(o.EiPaymentMethod.Get()) {
		var ret string
		return ret
	}
	return *o.EiPaymentMethod.Get()
}

// GetEiPaymentMethodOk returns a tuple with the EiPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetEiPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EiPaymentMethod.Get(), o.EiPaymentMethod.IsSet()
}

// HasEiPaymentMethod returns a boolean if a field has been set.
func (o *PaymentMethod) HasEiPaymentMethod() bool {
	if o != nil && o.EiPaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetEiPaymentMethod gets a reference to the given NullableString and assigns it to the EiPaymentMethod field.
func (o *PaymentMethod) SetEiPaymentMethod(v string) *PaymentMethod {
	o.EiPaymentMethod.Set(&v)
	return o
}
// SetEiPaymentMethodNil sets the value for EiPaymentMethod to be an explicit nil
func (o *PaymentMethod) SetEiPaymentMethodNil() *PaymentMethod {
	o.EiPaymentMethod.Set(nil)
	return o
}

// UnsetEiPaymentMethod ensures that no value is present for EiPaymentMethod, not even an explicit nil
func (o *PaymentMethod) UnsetEiPaymentMethod() {
	o.EiPaymentMethod.Unset()
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.IsDefault.IsSet() {
		toSerialize["is_default"] = o.IsDefault.Get()
	}
	if o.DefaultPaymentAccount.IsSet() {
		toSerialize["default_payment_account"] = o.DefaultPaymentAccount.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.BankIban.IsSet() {
		toSerialize["bank_iban"] = o.BankIban.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if o.BankBeneficiary.IsSet() {
		toSerialize["bank_beneficiary"] = o.BankBeneficiary.Get()
	}
	if o.EiPaymentMethod.IsSet() {
		toSerialize["ei_payment_method"] = o.EiPaymentMethod.Get()
	}
	return toSerialize, nil
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


