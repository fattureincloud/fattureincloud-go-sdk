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

// checks if the CompanyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyInfo{}

// CompanyInfo struct for CompanyInfo
type CompanyInfo struct {
	// Company id
	Id NullableInt32 `json:"id,omitempty"`
	// Company name
	Name NullableString `json:"name,omitempty"`
	// Company email
	Email NullableString `json:"email,omitempty"`
	Type *CompanyType `json:"type,omitempty"`
	AccessInfo NullableCompanyInfoAccessInfo `json:"access_info,omitempty"`
	FicLicenseExpire NullableString `json:"fic_license_expire,omitempty"`
	FicPlanName *FattureInCloudPlanType `json:"fic_plan_name,omitempty"`
	PlanInfo NullableCompanyInfoPlanInfo `json:"plan_info,omitempty"`
	// Company accountant id
	AccountantId NullableInt32 `json:"accountant_id,omitempty"`
	// Is the logged account an accountant.
	IsAccountant NullableBool `json:"is_accountant,omitempty"`
}

// NewCompanyInfo instantiates a new CompanyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyInfo() *CompanyInfo {
	this := CompanyInfo{}
	return &this
}

// NewCompanyInfoWithDefaults instantiates a new CompanyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyInfoWithDefaults() *CompanyInfo {
	this := CompanyInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CompanyInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *CompanyInfo) SetId(v int32) *CompanyInfo {
	o.Id.Set(&v)
		return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CompanyInfo) SetIdNil() *CompanyInfo {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CompanyInfo) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CompanyInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CompanyInfo) SetName(v string) *CompanyInfo {
	o.Name.Set(&v)
		return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CompanyInfo) SetNameNil() *CompanyInfo {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CompanyInfo) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CompanyInfo) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CompanyInfo) SetEmail(v string) *CompanyInfo {
	o.Email.Set(&v)
		return o
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CompanyInfo) SetEmailNil() *CompanyInfo {
	o.Email.Set(nil)
	return o
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CompanyInfo) UnsetEmail() {
	o.Email.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CompanyInfo) GetType() CompanyType {
	if o == nil || IsNil(o.Type) {
		var ret CompanyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfo) GetTypeOk() (*CompanyType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CompanyInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CompanyType and assigns it to the Type field.
func (o *CompanyInfo) SetType(v CompanyType) *CompanyInfo {
	o.Type = &v
		return o
}

// GetAccessInfo returns the AccessInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetAccessInfo() CompanyInfoAccessInfo {
	if o == nil || IsNil(o.AccessInfo.Get()) {
		var ret CompanyInfoAccessInfo
		return ret
	}
	return *o.AccessInfo.Get()
}

// GetAccessInfoOk returns a tuple with the AccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetAccessInfoOk() (*CompanyInfoAccessInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessInfo.Get(), o.AccessInfo.IsSet()
}

// HasAccessInfo returns a boolean if a field has been set.
func (o *CompanyInfo) HasAccessInfo() bool {
	if o != nil && o.AccessInfo.IsSet() {
		return true
	}

	return false
}

// SetAccessInfo gets a reference to the given NullableCompanyInfoAccessInfo and assigns it to the AccessInfo field.
func (o *CompanyInfo) SetAccessInfo(v CompanyInfoAccessInfo) *CompanyInfo {
	o.AccessInfo.Set(&v)
		return o
}
// SetAccessInfoNil sets the value for AccessInfo to be an explicit nil
func (o *CompanyInfo) SetAccessInfoNil() *CompanyInfo {
	o.AccessInfo.Set(nil)
	return o
}

// UnsetAccessInfo ensures that no value is present for AccessInfo, not even an explicit nil
func (o *CompanyInfo) UnsetAccessInfo() {
	o.AccessInfo.Unset()
}

// GetFicLicenseExpire returns the FicLicenseExpire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetFicLicenseExpire() string {
	if o == nil || IsNil(o.FicLicenseExpire.Get()) {
		var ret string
		return ret
	}
	return *o.FicLicenseExpire.Get()
}

// GetFicLicenseExpireOk returns a tuple with the FicLicenseExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetFicLicenseExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FicLicenseExpire.Get(), o.FicLicenseExpire.IsSet()
}

// HasFicLicenseExpire returns a boolean if a field has been set.
func (o *CompanyInfo) HasFicLicenseExpire() bool {
	if o != nil && o.FicLicenseExpire.IsSet() {
		return true
	}

	return false
}

// SetFicLicenseExpire gets a reference to the given NullableString and assigns it to the FicLicenseExpire field.
func (o *CompanyInfo) SetFicLicenseExpire(v string) *CompanyInfo {
	o.FicLicenseExpire.Set(&v)
		return o
}
// SetFicLicenseExpireNil sets the value for FicLicenseExpire to be an explicit nil
func (o *CompanyInfo) SetFicLicenseExpireNil() *CompanyInfo {
	o.FicLicenseExpire.Set(nil)
	return o
}

// UnsetFicLicenseExpire ensures that no value is present for FicLicenseExpire, not even an explicit nil
func (o *CompanyInfo) UnsetFicLicenseExpire() {
	o.FicLicenseExpire.Unset()
}

// GetFicPlanName returns the FicPlanName field value if set, zero value otherwise.
func (o *CompanyInfo) GetFicPlanName() FattureInCloudPlanType {
	if o == nil || IsNil(o.FicPlanName) {
		var ret FattureInCloudPlanType
		return ret
	}
	return *o.FicPlanName
}

// GetFicPlanNameOk returns a tuple with the FicPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyInfo) GetFicPlanNameOk() (*FattureInCloudPlanType, bool) {
	if o == nil || IsNil(o.FicPlanName) {
		return nil, false
	}
	return o.FicPlanName, true
}

// HasFicPlanName returns a boolean if a field has been set.
func (o *CompanyInfo) HasFicPlanName() bool {
	if o != nil && !IsNil(o.FicPlanName) {
		return true
	}

	return false
}

// SetFicPlanName gets a reference to the given FattureInCloudPlanType and assigns it to the FicPlanName field.
func (o *CompanyInfo) SetFicPlanName(v FattureInCloudPlanType) *CompanyInfo {
	o.FicPlanName = &v
		return o
}

// GetPlanInfo returns the PlanInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetPlanInfo() CompanyInfoPlanInfo {
	if o == nil || IsNil(o.PlanInfo.Get()) {
		var ret CompanyInfoPlanInfo
		return ret
	}
	return *o.PlanInfo.Get()
}

// GetPlanInfoOk returns a tuple with the PlanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetPlanInfoOk() (*CompanyInfoPlanInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanInfo.Get(), o.PlanInfo.IsSet()
}

// HasPlanInfo returns a boolean if a field has been set.
func (o *CompanyInfo) HasPlanInfo() bool {
	if o != nil && o.PlanInfo.IsSet() {
		return true
	}

	return false
}

// SetPlanInfo gets a reference to the given NullableCompanyInfoPlanInfo and assigns it to the PlanInfo field.
func (o *CompanyInfo) SetPlanInfo(v CompanyInfoPlanInfo) *CompanyInfo {
	o.PlanInfo.Set(&v)
		return o
}
// SetPlanInfoNil sets the value for PlanInfo to be an explicit nil
func (o *CompanyInfo) SetPlanInfoNil() *CompanyInfo {
	o.PlanInfo.Set(nil)
	return o
}

// UnsetPlanInfo ensures that no value is present for PlanInfo, not even an explicit nil
func (o *CompanyInfo) UnsetPlanInfo() {
	o.PlanInfo.Unset()
}

// GetAccountantId returns the AccountantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetAccountantId() int32 {
	if o == nil || IsNil(o.AccountantId.Get()) {
		var ret int32
		return ret
	}
	return *o.AccountantId.Get()
}

// GetAccountantIdOk returns a tuple with the AccountantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetAccountantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountantId.Get(), o.AccountantId.IsSet()
}

// HasAccountantId returns a boolean if a field has been set.
func (o *CompanyInfo) HasAccountantId() bool {
	if o != nil && o.AccountantId.IsSet() {
		return true
	}

	return false
}

// SetAccountantId gets a reference to the given NullableInt32 and assigns it to the AccountantId field.
func (o *CompanyInfo) SetAccountantId(v int32) *CompanyInfo {
	o.AccountantId.Set(&v)
		return o
}
// SetAccountantIdNil sets the value for AccountantId to be an explicit nil
func (o *CompanyInfo) SetAccountantIdNil() *CompanyInfo {
	o.AccountantId.Set(nil)
	return o
}

// UnsetAccountantId ensures that no value is present for AccountantId, not even an explicit nil
func (o *CompanyInfo) UnsetAccountantId() {
	o.AccountantId.Unset()
}

// GetIsAccountant returns the IsAccountant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfo) GetIsAccountant() bool {
	if o == nil || IsNil(o.IsAccountant.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAccountant.Get()
}

// GetIsAccountantOk returns a tuple with the IsAccountant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfo) GetIsAccountantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAccountant.Get(), o.IsAccountant.IsSet()
}

// HasIsAccountant returns a boolean if a field has been set.
func (o *CompanyInfo) HasIsAccountant() bool {
	if o != nil && o.IsAccountant.IsSet() {
		return true
	}

	return false
}

// SetIsAccountant gets a reference to the given NullableBool and assigns it to the IsAccountant field.
func (o *CompanyInfo) SetIsAccountant(v bool) *CompanyInfo {
	o.IsAccountant.Set(&v)
		return o
}
// SetIsAccountantNil sets the value for IsAccountant to be an explicit nil
func (o *CompanyInfo) SetIsAccountantNil() *CompanyInfo {
	o.IsAccountant.Set(nil)
	return o
}

// UnsetIsAccountant ensures that no value is present for IsAccountant, not even an explicit nil
func (o *CompanyInfo) UnsetIsAccountant() {
	o.IsAccountant.Unset()
}

func (o CompanyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.AccessInfo.IsSet() {
		toSerialize["access_info"] = o.AccessInfo.Get()
	}
	if o.FicLicenseExpire.IsSet() {
		toSerialize["fic_license_expire"] = o.FicLicenseExpire.Get()
	}
	if !IsNil(o.FicPlanName) {
		toSerialize["fic_plan_name"] = o.FicPlanName
	}
	if o.PlanInfo.IsSet() {
		toSerialize["plan_info"] = o.PlanInfo.Get()
	}
	if o.AccountantId.IsSet() {
		toSerialize["accountant_id"] = o.AccountantId.Get()
	}
	if o.IsAccountant.IsSet() {
		toSerialize["is_accountant"] = o.IsAccountant.Get()
	}
	return toSerialize, nil
}

type NullableCompanyInfo struct {
	value *CompanyInfo
	isSet bool
}

func (v NullableCompanyInfo) Get() *CompanyInfo {
	return v.value
}

func (v *NullableCompanyInfo) Set(val *CompanyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyInfo(val *CompanyInfo) *NullableCompanyInfo {
	return &NullableCompanyInfo{value: val, isSet: true}
}

func (v NullableCompanyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


