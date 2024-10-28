/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.2
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the DetailedCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedCountry{}

// DetailedCountry struct for DetailedCountry
type DetailedCountry struct {
	// Country name
Name NullableString `json:"name,omitempty"`
	// Country settings name
SettingsName NullableString `json:"settings_name,omitempty"`
	// Country iso code
Iso NullableString `json:"iso,omitempty"`
	// Country fiscal iso code
FiscalIso NullableString `json:"fiscal_iso,omitempty"`
	// Country uic
Uic NullableString `json:"uic,omitempty"`
}

// NewDetailedCountry instantiates a new DetailedCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedCountry() *DetailedCountry {
	this := DetailedCountry{}
	return &this
}

// NewDetailedCountryWithDefaults instantiates a new DetailedCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedCountryWithDefaults() *DetailedCountry {
	this := DetailedCountry{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedCountry) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedCountry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DetailedCountry) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DetailedCountry) SetName(v string) *DetailedCountry {
	o.Name.Set(&v)
        return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DetailedCountry) SetNameNil() *DetailedCountry {
	o.Name.Set(nil)
    return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DetailedCountry) UnsetName() {
	o.Name.Unset()
}

// GetSettingsName returns the SettingsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedCountry) GetSettingsName() string {
	if o == nil || IsNil(o.SettingsName.Get()) {
		var ret string
		return ret
	}
	return *o.SettingsName.Get()
}

// GetSettingsNameOk returns a tuple with the SettingsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedCountry) GetSettingsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingsName.Get(), o.SettingsName.IsSet()
}

// HasSettingsName returns a boolean if a field has been set.
func (o *DetailedCountry) HasSettingsName() bool {
	if o != nil && o.SettingsName.IsSet() {
		return true
	}

	return false
}

// SetSettingsName gets a reference to the given NullableString and assigns it to the SettingsName field.
func (o *DetailedCountry) SetSettingsName(v string) *DetailedCountry {
	o.SettingsName.Set(&v)
        return o
}
// SetSettingsNameNil sets the value for SettingsName to be an explicit nil
func (o *DetailedCountry) SetSettingsNameNil() *DetailedCountry {
	o.SettingsName.Set(nil)
    return o
}

// UnsetSettingsName ensures that no value is present for SettingsName, not even an explicit nil
func (o *DetailedCountry) UnsetSettingsName() {
	o.SettingsName.Unset()
}

// GetIso returns the Iso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedCountry) GetIso() string {
	if o == nil || IsNil(o.Iso.Get()) {
		var ret string
		return ret
	}
	return *o.Iso.Get()
}

// GetIsoOk returns a tuple with the Iso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedCountry) GetIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iso.Get(), o.Iso.IsSet()
}

// HasIso returns a boolean if a field has been set.
func (o *DetailedCountry) HasIso() bool {
	if o != nil && o.Iso.IsSet() {
		return true
	}

	return false
}

// SetIso gets a reference to the given NullableString and assigns it to the Iso field.
func (o *DetailedCountry) SetIso(v string) *DetailedCountry {
	o.Iso.Set(&v)
        return o
}
// SetIsoNil sets the value for Iso to be an explicit nil
func (o *DetailedCountry) SetIsoNil() *DetailedCountry {
	o.Iso.Set(nil)
    return o
}

// UnsetIso ensures that no value is present for Iso, not even an explicit nil
func (o *DetailedCountry) UnsetIso() {
	o.Iso.Unset()
}

// GetFiscalIso returns the FiscalIso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedCountry) GetFiscalIso() string {
	if o == nil || IsNil(o.FiscalIso.Get()) {
		var ret string
		return ret
	}
	return *o.FiscalIso.Get()
}

// GetFiscalIsoOk returns a tuple with the FiscalIso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedCountry) GetFiscalIsoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FiscalIso.Get(), o.FiscalIso.IsSet()
}

// HasFiscalIso returns a boolean if a field has been set.
func (o *DetailedCountry) HasFiscalIso() bool {
	if o != nil && o.FiscalIso.IsSet() {
		return true
	}

	return false
}

// SetFiscalIso gets a reference to the given NullableString and assigns it to the FiscalIso field.
func (o *DetailedCountry) SetFiscalIso(v string) *DetailedCountry {
	o.FiscalIso.Set(&v)
        return o
}
// SetFiscalIsoNil sets the value for FiscalIso to be an explicit nil
func (o *DetailedCountry) SetFiscalIsoNil() *DetailedCountry {
	o.FiscalIso.Set(nil)
    return o
}

// UnsetFiscalIso ensures that no value is present for FiscalIso, not even an explicit nil
func (o *DetailedCountry) UnsetFiscalIso() {
	o.FiscalIso.Unset()
}

// GetUic returns the Uic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailedCountry) GetUic() string {
	if o == nil || IsNil(o.Uic.Get()) {
		var ret string
		return ret
	}
	return *o.Uic.Get()
}

// GetUicOk returns a tuple with the Uic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailedCountry) GetUicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uic.Get(), o.Uic.IsSet()
}

// HasUic returns a boolean if a field has been set.
func (o *DetailedCountry) HasUic() bool {
	if o != nil && o.Uic.IsSet() {
		return true
	}

	return false
}

// SetUic gets a reference to the given NullableString and assigns it to the Uic field.
func (o *DetailedCountry) SetUic(v string) *DetailedCountry {
	o.Uic.Set(&v)
        return o
}
// SetUicNil sets the value for Uic to be an explicit nil
func (o *DetailedCountry) SetUicNil() *DetailedCountry {
	o.Uic.Set(nil)
    return o
}

// UnsetUic ensures that no value is present for Uic, not even an explicit nil
func (o *DetailedCountry) UnsetUic() {
	o.Uic.Unset()
}

func (o DetailedCountry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SettingsName.IsSet() {
		toSerialize["settings_name"] = o.SettingsName.Get()
	}
	if o.Iso.IsSet() {
		toSerialize["iso"] = o.Iso.Get()
	}
	if o.FiscalIso.IsSet() {
		toSerialize["fiscal_iso"] = o.FiscalIso.Get()
	}
	if o.Uic.IsSet() {
		toSerialize["uic"] = o.Uic.Get()
	}
	return toSerialize, nil
}

type NullableDetailedCountry struct {
	value *DetailedCountry
	isSet bool
}

func (v NullableDetailedCountry) Get() *DetailedCountry {
	return v.value
}

func (v *NullableDetailedCountry) Set(val *DetailedCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedCountry(val *DetailedCountry) *NullableDetailedCountry {
	return &NullableDetailedCountry{value: val, isSet: true}
}

func (v NullableDetailedCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


