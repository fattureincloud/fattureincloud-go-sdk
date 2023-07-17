/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.29
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the City type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &City{}

// City struct for City
type City struct {
	// City postal code
	PostalCode NullableString `json:"postal_code,omitempty"`
	// City name
	City NullableString `json:"city,omitempty"`
	// City province
	Province NullableString `json:"province,omitempty"`
}

// NewCity instantiates a new City object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCity() *City {
	this := City{}
	return &this
}

// NewCityWithDefaults instantiates a new City object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityWithDefaults() *City {
	this := City{}
	return &this
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *City) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *City) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *City) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *City) SetPostalCode(v string) *City {
	o.PostalCode.Set(&v)
	return o
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *City) SetPostalCodeNil() *City {
	o.PostalCode.Set(nil)
	return o
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *City) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *City) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *City) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *City) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *City) SetCity(v string) *City {
	o.City.Set(&v)
	return o
}
// SetCityNil sets the value for City to be an explicit nil
func (o *City) SetCityNil() *City {
	o.City.Set(nil)
	return o
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *City) UnsetCity() {
	o.City.Unset()
}

// GetProvince returns the Province field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *City) GetProvince() string {
	if o == nil || IsNil(o.Province.Get()) {
		var ret string
		return ret
	}
	return *o.Province.Get()
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *City) GetProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Province.Get(), o.Province.IsSet()
}

// HasProvince returns a boolean if a field has been set.
func (o *City) HasProvince() bool {
	if o != nil && o.Province.IsSet() {
		return true
	}

	return false
}

// SetProvince gets a reference to the given NullableString and assigns it to the Province field.
func (o *City) SetProvince(v string) *City {
	o.Province.Set(&v)
	return o
}
// SetProvinceNil sets the value for Province to be an explicit nil
func (o *City) SetProvinceNil() *City {
	o.Province.Set(nil)
	return o
}

// UnsetProvince ensures that no value is present for Province, not even an explicit nil
func (o *City) UnsetProvince() {
	o.Province.Unset()
}

func (o City) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o City) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Province.IsSet() {
		toSerialize["province"] = o.Province.Get()
	}
	return toSerialize, nil
}

type NullableCity struct {
	value *City
	isSet bool
}

func (v NullableCity) Get() *City {
	return v.value
}

func (v *NullableCity) Set(val *City) {
	v.value = val
	v.isSet = true
}

func (v NullableCity) IsSet() bool {
	return v.isSet
}

func (v *NullableCity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCity(val *City) *NullableCity {
	return &NullableCity{value: val, isSet: true}
}

func (v NullableCity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


