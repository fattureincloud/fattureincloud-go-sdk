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

// checks if the ControlledCompany type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlledCompany{}

// ControlledCompany struct for ControlledCompany
type ControlledCompany struct {
	// Controlled company id
	Id NullableInt32 `json:"id,omitempty"`
	// Controlled company id
	Name NullableString `json:"name,omitempty"`
	Type *CompanyType `json:"type,omitempty"`
	// Controlled company access token Only if type=company]
	AccessToken NullableString `json:"access_token,omitempty"`
	// Controlled company connection id
	ConnectionId NullableFloat32 `json:"connection_id,omitempty"`
	// Controlled company tax code
	TaxCode NullableString `json:"tax_code,omitempty"`
}

// NewControlledCompany instantiates a new ControlledCompany object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlledCompany() *ControlledCompany {
	this := ControlledCompany{}
	return &this
}

// NewControlledCompanyWithDefaults instantiates a new ControlledCompany object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlledCompanyWithDefaults() *ControlledCompany {
	this := ControlledCompany{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlledCompany) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlledCompany) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ControlledCompany) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ControlledCompany) SetId(v int32) *ControlledCompany {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ControlledCompany) SetIdNil() *ControlledCompany {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ControlledCompany) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlledCompany) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlledCompany) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ControlledCompany) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ControlledCompany) SetName(v string) *ControlledCompany {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ControlledCompany) SetNameNil() *ControlledCompany {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ControlledCompany) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ControlledCompany) GetType() CompanyType {
	if o == nil || IsNil(o.Type) {
		var ret CompanyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlledCompany) GetTypeOk() (*CompanyType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ControlledCompany) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CompanyType and assigns it to the Type field.
func (o *ControlledCompany) SetType(v CompanyType) *ControlledCompany {
	o.Type = &v
	return o
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlledCompany) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlledCompany) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *ControlledCompany) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *ControlledCompany) SetAccessToken(v string) *ControlledCompany {
	o.AccessToken.Set(&v)
	return o
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *ControlledCompany) SetAccessTokenNil() *ControlledCompany {
	o.AccessToken.Set(nil)
	return o
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *ControlledCompany) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlledCompany) GetConnectionId() float32 {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret float32
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlledCompany) GetConnectionIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *ControlledCompany) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableFloat32 and assigns it to the ConnectionId field.
func (o *ControlledCompany) SetConnectionId(v float32) *ControlledCompany {
	o.ConnectionId.Set(&v)
	return o
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *ControlledCompany) SetConnectionIdNil() *ControlledCompany {
	o.ConnectionId.Set(nil)
	return o
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *ControlledCompany) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlledCompany) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode.Get()) {
		var ret string
		return ret
	}
	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlledCompany) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// HasTaxCode returns a boolean if a field has been set.
func (o *ControlledCompany) HasTaxCode() bool {
	if o != nil && o.TaxCode.IsSet() {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given NullableString and assigns it to the TaxCode field.
func (o *ControlledCompany) SetTaxCode(v string) *ControlledCompany {
	o.TaxCode.Set(&v)
	return o
}
// SetTaxCodeNil sets the value for TaxCode to be an explicit nil
func (o *ControlledCompany) SetTaxCodeNil() *ControlledCompany {
	o.TaxCode.Set(nil)
	return o
}

// UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
func (o *ControlledCompany) UnsetTaxCode() {
	o.TaxCode.Unset()
}

func (o ControlledCompany) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlledCompany) ToMap() (map[string]interface{}, error) {
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
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connection_id"] = o.ConnectionId.Get()
	}
	if o.TaxCode.IsSet() {
		toSerialize["tax_code"] = o.TaxCode.Get()
	}
	return toSerialize, nil
}

type NullableControlledCompany struct {
	value *ControlledCompany
	isSet bool
}

func (v NullableControlledCompany) Get() *ControlledCompany {
	return v.value
}

func (v *NullableControlledCompany) Set(val *ControlledCompany) {
	v.value = val
	v.isSet = true
}

func (v NullableControlledCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableControlledCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlledCompany(val *ControlledCompany) *NullableControlledCompany {
	return &NullableControlledCompany{value: val, isSet: true}
}

func (v NullableControlledCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlledCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


