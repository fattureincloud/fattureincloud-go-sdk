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
	"time"
)

// checks if the EInvoiceRejectionReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EInvoiceRejectionReason{}

// EInvoiceRejectionReason struct for EInvoiceRejectionReason
type EInvoiceRejectionReason struct {
	// E-invoice rejection reason
Reason NullableString `json:"reason,omitempty"`
	// E-invoice status
EiStatus NullableString `json:"ei_status,omitempty"`
	// Error solution.
Solution NullableString `json:"solution,omitempty"`
	// E-invoice rejection error code
Code NullableString `json:"code,omitempty"`
	// E-invoice rejection date
Date NullableTime `json:"date,omitempty"`
}

// NewEInvoiceRejectionReason instantiates a new EInvoiceRejectionReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEInvoiceRejectionReason() *EInvoiceRejectionReason {
	this := EInvoiceRejectionReason{}
	return &this
}

// NewEInvoiceRejectionReasonWithDefaults instantiates a new EInvoiceRejectionReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEInvoiceRejectionReasonWithDefaults() *EInvoiceRejectionReason {
	this := EInvoiceRejectionReason{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EInvoiceRejectionReason) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EInvoiceRejectionReason) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *EInvoiceRejectionReason) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *EInvoiceRejectionReason) SetReason(v string) *EInvoiceRejectionReason {
	o.Reason.Set(&v)
		return o
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *EInvoiceRejectionReason) SetReasonNil() *EInvoiceRejectionReason {
	o.Reason.Set(nil)
	return o
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *EInvoiceRejectionReason) UnsetReason() {
	o.Reason.Unset()
}

// GetEiStatus returns the EiStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EInvoiceRejectionReason) GetEiStatus() string {
	if o == nil || IsNil(o.EiStatus.Get()) {
		var ret string
		return ret
	}
	return *o.EiStatus.Get()
}

// GetEiStatusOk returns a tuple with the EiStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EInvoiceRejectionReason) GetEiStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EiStatus.Get(), o.EiStatus.IsSet()
}

// HasEiStatus returns a boolean if a field has been set.
func (o *EInvoiceRejectionReason) HasEiStatus() bool {
	if o != nil && o.EiStatus.IsSet() {
		return true
	}

	return false
}

// SetEiStatus gets a reference to the given NullableString and assigns it to the EiStatus field.
func (o *EInvoiceRejectionReason) SetEiStatus(v string) *EInvoiceRejectionReason {
	o.EiStatus.Set(&v)
		return o
}
// SetEiStatusNil sets the value for EiStatus to be an explicit nil
func (o *EInvoiceRejectionReason) SetEiStatusNil() *EInvoiceRejectionReason {
	o.EiStatus.Set(nil)
	return o
}

// UnsetEiStatus ensures that no value is present for EiStatus, not even an explicit nil
func (o *EInvoiceRejectionReason) UnsetEiStatus() {
	o.EiStatus.Unset()
}

// GetSolution returns the Solution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EInvoiceRejectionReason) GetSolution() string {
	if o == nil || IsNil(o.Solution.Get()) {
		var ret string
		return ret
	}
	return *o.Solution.Get()
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EInvoiceRejectionReason) GetSolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Solution.Get(), o.Solution.IsSet()
}

// HasSolution returns a boolean if a field has been set.
func (o *EInvoiceRejectionReason) HasSolution() bool {
	if o != nil && o.Solution.IsSet() {
		return true
	}

	return false
}

// SetSolution gets a reference to the given NullableString and assigns it to the Solution field.
func (o *EInvoiceRejectionReason) SetSolution(v string) *EInvoiceRejectionReason {
	o.Solution.Set(&v)
		return o
}
// SetSolutionNil sets the value for Solution to be an explicit nil
func (o *EInvoiceRejectionReason) SetSolutionNil() *EInvoiceRejectionReason {
	o.Solution.Set(nil)
	return o
}

// UnsetSolution ensures that no value is present for Solution, not even an explicit nil
func (o *EInvoiceRejectionReason) UnsetSolution() {
	o.Solution.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EInvoiceRejectionReason) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EInvoiceRejectionReason) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *EInvoiceRejectionReason) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *EInvoiceRejectionReason) SetCode(v string) *EInvoiceRejectionReason {
	o.Code.Set(&v)
		return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *EInvoiceRejectionReason) SetCodeNil() *EInvoiceRejectionReason {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *EInvoiceRejectionReason) UnsetCode() {
	o.Code.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EInvoiceRejectionReason) GetDate() time.Time {
	if o == nil || IsNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EInvoiceRejectionReason) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *EInvoiceRejectionReason) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *EInvoiceRejectionReason) SetDate(v time.Time) *EInvoiceRejectionReason {
	o.Date.Set(&v)
		return o
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *EInvoiceRejectionReason) SetDateNil() *EInvoiceRejectionReason {
	o.Date.Set(nil)
	return o
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *EInvoiceRejectionReason) UnsetDate() {
	o.Date.Unset()
}

func (o EInvoiceRejectionReason) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EInvoiceRejectionReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.EiStatus.IsSet() {
		toSerialize["ei_status"] = o.EiStatus.Get()
	}
	if o.Solution.IsSet() {
		toSerialize["solution"] = o.Solution.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	return toSerialize, nil
}

type NullableEInvoiceRejectionReason struct {
	value *EInvoiceRejectionReason
	isSet bool
}

func (v NullableEInvoiceRejectionReason) Get() *EInvoiceRejectionReason {
	return v.value
}

func (v *NullableEInvoiceRejectionReason) Set(val *EInvoiceRejectionReason) {
	v.value = val
	v.isSet = true
}

func (v NullableEInvoiceRejectionReason) IsSet() bool {
	return v.isSet
}

func (v *NullableEInvoiceRejectionReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEInvoiceRejectionReason(val *EInvoiceRejectionReason) *NullableEInvoiceRejectionReason {
	return &NullableEInvoiceRejectionReason{value: val, isSet: true}
}

func (v NullableEInvoiceRejectionReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEInvoiceRejectionReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


