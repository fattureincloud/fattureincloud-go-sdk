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

// checks if the Company type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Company{}

// Company struct for Company
type Company struct {
	// Company id
	Id NullableInt32 `json:"id,omitempty"`
	// Company name
	Name NullableString `json:"name,omitempty"`
	Type *CompanyType `json:"type,omitempty"`
	// Company authentication token for this company. [Only if type=company]
	AccessToken NullableString `json:"access_token,omitempty"`
	// Company list of controlled companies [Only if type=accountant]
	ControlledCompanies []ControlledCompany `json:"controlled_companies,omitempty"`
	FicLicenseExpire NullableString `json:"fic_license_expire,omitempty"`
	FicPlan *FattureInCloudPlanType `json:"fic_plan,omitempty"`
	// Company connection id
	ConnectionId NullableInt32 `json:"connection_id,omitempty"`
	// Company tax code
	TaxCode NullableString `json:"tax_code,omitempty"`
	// Company vat number
	VatNumber NullableString `json:"vat_number,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Company) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Company) SetId(v int32) *Company {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Company) SetIdNil() *Company {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Company) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Company) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Company) SetName(v string) *Company {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Company) SetNameNil() *Company {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Company) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Company) GetType() CompanyType {
	if o == nil || IsNil(o.Type) {
		var ret CompanyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetTypeOk() (*CompanyType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Company) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CompanyType and assigns it to the Type field.
func (o *Company) SetType(v CompanyType) *Company {
	o.Type = &v
	return o
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *Company) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *Company) SetAccessToken(v string) *Company {
	o.AccessToken.Set(&v)
	return o
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *Company) SetAccessTokenNil() *Company {
	o.AccessToken.Set(nil)
	return o
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *Company) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetControlledCompanies returns the ControlledCompanies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetControlledCompanies() []ControlledCompany {
	if o == nil {
		var ret []ControlledCompany
		return ret
	}
	return o.ControlledCompanies
}

// GetControlledCompaniesOk returns a tuple with the ControlledCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetControlledCompaniesOk() ([]ControlledCompany, bool) {
	if o == nil || IsNil(o.ControlledCompanies) {
		return nil, false
	}
	return o.ControlledCompanies, true
}

// HasControlledCompanies returns a boolean if a field has been set.
func (o *Company) HasControlledCompanies() bool {
	if o != nil && !IsNil(o.ControlledCompanies) {
		return true
	}

	return false
}

// SetControlledCompanies gets a reference to the given []ControlledCompany and assigns it to the ControlledCompanies field.
func (o *Company) SetControlledCompanies(v []ControlledCompany) *Company {
	o.ControlledCompanies = v
	return o
}

// GetFicLicenseExpire returns the FicLicenseExpire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetFicLicenseExpire() string {
	if o == nil || IsNil(o.FicLicenseExpire.Get()) {
		var ret string
		return ret
	}
	return *o.FicLicenseExpire.Get()
}

// GetFicLicenseExpireOk returns a tuple with the FicLicenseExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetFicLicenseExpireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FicLicenseExpire.Get(), o.FicLicenseExpire.IsSet()
}

// HasFicLicenseExpire returns a boolean if a field has been set.
func (o *Company) HasFicLicenseExpire() bool {
	if o != nil && o.FicLicenseExpire.IsSet() {
		return true
	}

	return false
}

// SetFicLicenseExpire gets a reference to the given NullableString and assigns it to the FicLicenseExpire field.
func (o *Company) SetFicLicenseExpire(v string) *Company {
	o.FicLicenseExpire.Set(&v)
	return o
}
// SetFicLicenseExpireNil sets the value for FicLicenseExpire to be an explicit nil
func (o *Company) SetFicLicenseExpireNil() *Company {
	o.FicLicenseExpire.Set(nil)
	return o
}

// UnsetFicLicenseExpire ensures that no value is present for FicLicenseExpire, not even an explicit nil
func (o *Company) UnsetFicLicenseExpire() {
	o.FicLicenseExpire.Unset()
}

// GetFicPlan returns the FicPlan field value if set, zero value otherwise.
func (o *Company) GetFicPlan() FattureInCloudPlanType {
	if o == nil || IsNil(o.FicPlan) {
		var ret FattureInCloudPlanType
		return ret
	}
	return *o.FicPlan
}

// GetFicPlanOk returns a tuple with the FicPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetFicPlanOk() (*FattureInCloudPlanType, bool) {
	if o == nil || IsNil(o.FicPlan) {
		return nil, false
	}
	return o.FicPlan, true
}

// HasFicPlan returns a boolean if a field has been set.
func (o *Company) HasFicPlan() bool {
	if o != nil && !IsNil(o.FicPlan) {
		return true
	}

	return false
}

// SetFicPlan gets a reference to the given FattureInCloudPlanType and assigns it to the FicPlan field.
func (o *Company) SetFicPlan(v FattureInCloudPlanType) *Company {
	o.FicPlan = &v
	return o
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetConnectionId() int32 {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret int32
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetConnectionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *Company) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt32 and assigns it to the ConnectionId field.
func (o *Company) SetConnectionId(v int32) *Company {
	o.ConnectionId.Set(&v)
	return o
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *Company) SetConnectionIdNil() *Company {
	o.ConnectionId.Set(nil)
	return o
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *Company) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode.Get()) {
		var ret string
		return ret
	}
	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// HasTaxCode returns a boolean if a field has been set.
func (o *Company) HasTaxCode() bool {
	if o != nil && o.TaxCode.IsSet() {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given NullableString and assigns it to the TaxCode field.
func (o *Company) SetTaxCode(v string) *Company {
	o.TaxCode.Set(&v)
	return o
}
// SetTaxCodeNil sets the value for TaxCode to be an explicit nil
func (o *Company) SetTaxCodeNil() *Company {
	o.TaxCode.Set(nil)
	return o
}

// UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
func (o *Company) UnsetTaxCode() {
	o.TaxCode.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Company) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *Company) SetVatNumber(v string) *Company {
	o.VatNumber.Set(&v)
	return o
}
// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *Company) SetVatNumberNil() *Company {
	o.VatNumber.Set(nil)
	return o
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *Company) UnsetVatNumber() {
	o.VatNumber.Unset()
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Company) ToMap() (map[string]interface{}, error) {
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
	if o.ControlledCompanies != nil {
		toSerialize["controlled_companies"] = o.ControlledCompanies
	}
	if o.FicLicenseExpire.IsSet() {
		toSerialize["fic_license_expire"] = o.FicLicenseExpire.Get()
	}
	if !IsNil(o.FicPlan) {
		toSerialize["fic_plan"] = o.FicPlan
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connection_id"] = o.ConnectionId.Get()
	}
	if o.TaxCode.IsSet() {
		toSerialize["tax_code"] = o.TaxCode.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vat_number"] = o.VatNumber.Get()
	}
	return toSerialize, nil
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


