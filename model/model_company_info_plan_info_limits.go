/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.23
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// CompanyInfoPlanInfoLimits Limits for this company.
type CompanyInfoPlanInfoLimits struct {
	Clients NullableInt32 `json:"clients,omitempty"`
	Suppliers NullableInt32 `json:"suppliers,omitempty"`
	Products NullableInt32 `json:"products,omitempty"`
	Documents NullableInt32 `json:"documents,omitempty"`
}

// NewCompanyInfoPlanInfoLimits instantiates a new CompanyInfoPlanInfoLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyInfoPlanInfoLimits() *CompanyInfoPlanInfoLimits {
	this := CompanyInfoPlanInfoLimits{}
	return &this
}

// NewCompanyInfoPlanInfoLimitsWithDefaults instantiates a new CompanyInfoPlanInfoLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyInfoPlanInfoLimitsWithDefaults() *CompanyInfoPlanInfoLimits {
	this := CompanyInfoPlanInfoLimits{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfoLimits) GetClients() int32 {
	if o == nil || isNil(o.Clients.Get()) {
		var ret int32
		return ret
	}
	return *o.Clients.Get()
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfoLimits) GetClientsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clients.Get(), o.Clients.IsSet()
}

// HasClients returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoLimits) HasClients() bool {
	if o != nil && o.Clients.IsSet() {
		return true
	}

	return false
}

// SetClients gets a reference to the given NullableInt32 and assigns it to the Clients field.
func (o *CompanyInfoPlanInfoLimits) SetClients(v int32) *CompanyInfoPlanInfoLimits {
	o.Clients.Set(&v)
	return o
}
// SetClientsNil sets the value for Clients to be an explicit nil
func (o *CompanyInfoPlanInfoLimits) SetClientsNil() *CompanyInfoPlanInfoLimits {
	o.Clients.Set(nil)
	return o
}

// UnsetClients ensures that no value is present for Clients, not even an explicit nil
func (o *CompanyInfoPlanInfoLimits) UnsetClients() {
	o.Clients.Unset()
}

// GetSuppliers returns the Suppliers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfoLimits) GetSuppliers() int32 {
	if o == nil || isNil(o.Suppliers.Get()) {
		var ret int32
		return ret
	}
	return *o.Suppliers.Get()
}

// GetSuppliersOk returns a tuple with the Suppliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfoLimits) GetSuppliersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suppliers.Get(), o.Suppliers.IsSet()
}

// HasSuppliers returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoLimits) HasSuppliers() bool {
	if o != nil && o.Suppliers.IsSet() {
		return true
	}

	return false
}

// SetSuppliers gets a reference to the given NullableInt32 and assigns it to the Suppliers field.
func (o *CompanyInfoPlanInfoLimits) SetSuppliers(v int32) *CompanyInfoPlanInfoLimits {
	o.Suppliers.Set(&v)
	return o
}
// SetSuppliersNil sets the value for Suppliers to be an explicit nil
func (o *CompanyInfoPlanInfoLimits) SetSuppliersNil() *CompanyInfoPlanInfoLimits {
	o.Suppliers.Set(nil)
	return o
}

// UnsetSuppliers ensures that no value is present for Suppliers, not even an explicit nil
func (o *CompanyInfoPlanInfoLimits) UnsetSuppliers() {
	o.Suppliers.Unset()
}

// GetProducts returns the Products field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfoLimits) GetProducts() int32 {
	if o == nil || isNil(o.Products.Get()) {
		var ret int32
		return ret
	}
	return *o.Products.Get()
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfoLimits) GetProductsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products.Get(), o.Products.IsSet()
}

// HasProducts returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoLimits) HasProducts() bool {
	if o != nil && o.Products.IsSet() {
		return true
	}

	return false
}

// SetProducts gets a reference to the given NullableInt32 and assigns it to the Products field.
func (o *CompanyInfoPlanInfoLimits) SetProducts(v int32) *CompanyInfoPlanInfoLimits {
	o.Products.Set(&v)
	return o
}
// SetProductsNil sets the value for Products to be an explicit nil
func (o *CompanyInfoPlanInfoLimits) SetProductsNil() *CompanyInfoPlanInfoLimits {
	o.Products.Set(nil)
	return o
}

// UnsetProducts ensures that no value is present for Products, not even an explicit nil
func (o *CompanyInfoPlanInfoLimits) UnsetProducts() {
	o.Products.Unset()
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyInfoPlanInfoLimits) GetDocuments() int32 {
	if o == nil || isNil(o.Documents.Get()) {
		var ret int32
		return ret
	}
	return *o.Documents.Get()
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyInfoPlanInfoLimits) GetDocumentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents.Get(), o.Documents.IsSet()
}

// HasDocuments returns a boolean if a field has been set.
func (o *CompanyInfoPlanInfoLimits) HasDocuments() bool {
	if o != nil && o.Documents.IsSet() {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given NullableInt32 and assigns it to the Documents field.
func (o *CompanyInfoPlanInfoLimits) SetDocuments(v int32) *CompanyInfoPlanInfoLimits {
	o.Documents.Set(&v)
	return o
}
// SetDocumentsNil sets the value for Documents to be an explicit nil
func (o *CompanyInfoPlanInfoLimits) SetDocumentsNil() *CompanyInfoPlanInfoLimits {
	o.Documents.Set(nil)
	return o
}

// UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
func (o *CompanyInfoPlanInfoLimits) UnsetDocuments() {
	o.Documents.Unset()
}

func (o CompanyInfoPlanInfoLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients.IsSet() {
		toSerialize["clients"] = o.Clients.Get()
	}
	if o.Suppliers.IsSet() {
		toSerialize["suppliers"] = o.Suppliers.Get()
	}
	if o.Products.IsSet() {
		toSerialize["products"] = o.Products.Get()
	}
	if o.Documents.IsSet() {
		toSerialize["documents"] = o.Documents.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyInfoPlanInfoLimits struct {
	value *CompanyInfoPlanInfoLimits
	isSet bool
}

func (v NullableCompanyInfoPlanInfoLimits) Get() *CompanyInfoPlanInfoLimits {
	return v.value
}

func (v *NullableCompanyInfoPlanInfoLimits) Set(val *CompanyInfoPlanInfoLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyInfoPlanInfoLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyInfoPlanInfoLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyInfoPlanInfoLimits(val *CompanyInfoPlanInfoLimits) *NullableCompanyInfoPlanInfoLimits {
	return &NullableCompanyInfoPlanInfoLimits{value: val, isSet: true}
}

func (v NullableCompanyInfoPlanInfoLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyInfoPlanInfoLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


