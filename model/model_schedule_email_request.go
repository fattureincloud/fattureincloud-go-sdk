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

// checks if the ScheduleEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleEmailRequest{}

// ScheduleEmailRequest 
type ScheduleEmailRequest struct {
	Data *EmailSchedule `json:"data,omitempty"`
}

// NewScheduleEmailRequest instantiates a new ScheduleEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleEmailRequest() *ScheduleEmailRequest {
	this := ScheduleEmailRequest{}
	return &this
}

// NewScheduleEmailRequestWithDefaults instantiates a new ScheduleEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleEmailRequestWithDefaults() *ScheduleEmailRequest {
	this := ScheduleEmailRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScheduleEmailRequest) GetData() EmailSchedule {
	if o == nil || IsNil(o.Data) {
		var ret EmailSchedule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleEmailRequest) GetDataOk() (*EmailSchedule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScheduleEmailRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailSchedule and assigns it to the Data field.
func (o *ScheduleEmailRequest) SetData(v EmailSchedule) *ScheduleEmailRequest {
	o.Data = &v
	return o
}

func (o ScheduleEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableScheduleEmailRequest struct {
	value *ScheduleEmailRequest
	isSet bool
}

func (v NullableScheduleEmailRequest) Get() *ScheduleEmailRequest {
	return v.value
}

func (v *NullableScheduleEmailRequest) Set(val *ScheduleEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleEmailRequest(val *ScheduleEmailRequest) *NullableScheduleEmailRequest {
	return &NullableScheduleEmailRequest{value: val, isSet: true}
}

func (v NullableScheduleEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


