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

// checks if the Supplier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Supplier{}

// Supplier struct for Supplier
type Supplier struct {
	// Supplier id
	Id NullableInt32 `json:"id,omitempty"`
	// Supplier code
	Code NullableString `json:"code,omitempty"`
	// Supplier name
	Name NullableString `json:"name,omitempty"`
	Type NullableSupplierType `json:"type,omitempty"`
	// Supplier first name
	FirstName NullableString `json:"first_name,omitempty"`
	// Supplier last name
	LastName NullableString `json:"last_name,omitempty"`
	// Supplier contact person
	ContactPerson NullableString `json:"contact_person,omitempty"`
	// Supplier vat number
	VatNumber NullableString `json:"vat_number,omitempty"`
	// Supplier tax code
	TaxCode NullableString `json:"tax_code,omitempty"`
	// Supplier street address
	AddressStreet NullableString `json:"address_street,omitempty"`
	// Supplier postal code
	AddressPostalCode NullableString `json:"address_postal_code,omitempty"`
	// Supplier city
	AddressCity NullableString `json:"address_city,omitempty"`
	// Supplier province
	AddressProvince NullableString `json:"address_province,omitempty"`
	// Supplier address extra info
	AddressExtra NullableString `json:"address_extra,omitempty"`
	// Supplier country
	Country NullableString `json:"country,omitempty"`
	// Supplier email
	Email NullableString `json:"email,omitempty"`
	// Supplier certified email
	CertifiedEmail NullableString `json:"certified_email,omitempty"`
	// Supplier phone
	Phone NullableString `json:"phone,omitempty"`
	// Supplier fax
	Fax NullableString `json:"fax,omitempty"`
	// Supplier extra notes
	Notes NullableString `json:"notes,omitempty"`
	// Supplier bank IBAN
	BankIban NullableString `json:"bank_iban,omitempty"`
	// Supplier creation date
	CreatedAt NullableString `json:"created_at,omitempty"`
	// Supplier last update date
	UpdatedAt NullableString `json:"updated_at,omitempty"`
}

// NewSupplier instantiates a new Supplier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplier() *Supplier {
	this := Supplier{}
	return &this
}

// NewSupplierWithDefaults instantiates a new Supplier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplierWithDefaults() *Supplier {
	this := Supplier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Supplier) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Supplier) SetId(v int32) *Supplier {
	o.Id.Set(&v)
	return o
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Supplier) SetIdNil() *Supplier {
	o.Id.Set(nil)
	return o
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Supplier) UnsetId() {
	o.Id.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *Supplier) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *Supplier) SetCode(v string) *Supplier {
	o.Code.Set(&v)
	return o
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *Supplier) SetCodeNil() *Supplier {
	o.Code.Set(nil)
	return o
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *Supplier) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Supplier) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Supplier) SetName(v string) *Supplier {
	o.Name.Set(&v)
	return o
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Supplier) SetNameNil() *Supplier {
	o.Name.Set(nil)
	return o
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Supplier) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetType() SupplierType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret SupplierType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetTypeOk() (*SupplierType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Supplier) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableSupplierType and assigns it to the Type field.
func (o *Supplier) SetType(v SupplierType) *Supplier {
	o.Type.Set(&v)
	return o
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Supplier) SetTypeNil() *Supplier {
	o.Type.Set(nil)
	return o
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Supplier) UnsetType() {
	o.Type.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *Supplier) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *Supplier) SetFirstName(v string) *Supplier {
	o.FirstName.Set(&v)
	return o
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *Supplier) SetFirstNameNil() *Supplier {
	o.FirstName.Set(nil)
	return o
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *Supplier) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *Supplier) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *Supplier) SetLastName(v string) *Supplier {
	o.LastName.Set(&v)
	return o
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *Supplier) SetLastNameNil() *Supplier {
	o.LastName.Set(nil)
	return o
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *Supplier) UnsetLastName() {
	o.LastName.Unset()
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetContactPerson() string {
	if o == nil || IsNil(o.ContactPerson.Get()) {
		var ret string
		return ret
	}
	return *o.ContactPerson.Get()
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetContactPersonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactPerson.Get(), o.ContactPerson.IsSet()
}

// HasContactPerson returns a boolean if a field has been set.
func (o *Supplier) HasContactPerson() bool {
	if o != nil && o.ContactPerson.IsSet() {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given NullableString and assigns it to the ContactPerson field.
func (o *Supplier) SetContactPerson(v string) *Supplier {
	o.ContactPerson.Set(&v)
	return o
}
// SetContactPersonNil sets the value for ContactPerson to be an explicit nil
func (o *Supplier) SetContactPersonNil() *Supplier {
	o.ContactPerson.Set(nil)
	return o
}

// UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
func (o *Supplier) UnsetContactPerson() {
	o.ContactPerson.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Supplier) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *Supplier) SetVatNumber(v string) *Supplier {
	o.VatNumber.Set(&v)
	return o
}
// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *Supplier) SetVatNumberNil() *Supplier {
	o.VatNumber.Set(nil)
	return o
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *Supplier) UnsetVatNumber() {
	o.VatNumber.Unset()
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode.Get()) {
		var ret string
		return ret
	}
	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// HasTaxCode returns a boolean if a field has been set.
func (o *Supplier) HasTaxCode() bool {
	if o != nil && o.TaxCode.IsSet() {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given NullableString and assigns it to the TaxCode field.
func (o *Supplier) SetTaxCode(v string) *Supplier {
	o.TaxCode.Set(&v)
	return o
}
// SetTaxCodeNil sets the value for TaxCode to be an explicit nil
func (o *Supplier) SetTaxCodeNil() *Supplier {
	o.TaxCode.Set(nil)
	return o
}

// UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
func (o *Supplier) UnsetTaxCode() {
	o.TaxCode.Unset()
}

// GetAddressStreet returns the AddressStreet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetAddressStreet() string {
	if o == nil || IsNil(o.AddressStreet.Get()) {
		var ret string
		return ret
	}
	return *o.AddressStreet.Get()
}

// GetAddressStreetOk returns a tuple with the AddressStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetAddressStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressStreet.Get(), o.AddressStreet.IsSet()
}

// HasAddressStreet returns a boolean if a field has been set.
func (o *Supplier) HasAddressStreet() bool {
	if o != nil && o.AddressStreet.IsSet() {
		return true
	}

	return false
}

// SetAddressStreet gets a reference to the given NullableString and assigns it to the AddressStreet field.
func (o *Supplier) SetAddressStreet(v string) *Supplier {
	o.AddressStreet.Set(&v)
	return o
}
// SetAddressStreetNil sets the value for AddressStreet to be an explicit nil
func (o *Supplier) SetAddressStreetNil() *Supplier {
	o.AddressStreet.Set(nil)
	return o
}

// UnsetAddressStreet ensures that no value is present for AddressStreet, not even an explicit nil
func (o *Supplier) UnsetAddressStreet() {
	o.AddressStreet.Unset()
}

// GetAddressPostalCode returns the AddressPostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetAddressPostalCode() string {
	if o == nil || IsNil(o.AddressPostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.AddressPostalCode.Get()
}

// GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetAddressPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressPostalCode.Get(), o.AddressPostalCode.IsSet()
}

// HasAddressPostalCode returns a boolean if a field has been set.
func (o *Supplier) HasAddressPostalCode() bool {
	if o != nil && o.AddressPostalCode.IsSet() {
		return true
	}

	return false
}

// SetAddressPostalCode gets a reference to the given NullableString and assigns it to the AddressPostalCode field.
func (o *Supplier) SetAddressPostalCode(v string) *Supplier {
	o.AddressPostalCode.Set(&v)
	return o
}
// SetAddressPostalCodeNil sets the value for AddressPostalCode to be an explicit nil
func (o *Supplier) SetAddressPostalCodeNil() *Supplier {
	o.AddressPostalCode.Set(nil)
	return o
}

// UnsetAddressPostalCode ensures that no value is present for AddressPostalCode, not even an explicit nil
func (o *Supplier) UnsetAddressPostalCode() {
	o.AddressPostalCode.Unset()
}

// GetAddressCity returns the AddressCity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetAddressCity() string {
	if o == nil || IsNil(o.AddressCity.Get()) {
		var ret string
		return ret
	}
	return *o.AddressCity.Get()
}

// GetAddressCityOk returns a tuple with the AddressCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetAddressCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressCity.Get(), o.AddressCity.IsSet()
}

// HasAddressCity returns a boolean if a field has been set.
func (o *Supplier) HasAddressCity() bool {
	if o != nil && o.AddressCity.IsSet() {
		return true
	}

	return false
}

// SetAddressCity gets a reference to the given NullableString and assigns it to the AddressCity field.
func (o *Supplier) SetAddressCity(v string) *Supplier {
	o.AddressCity.Set(&v)
	return o
}
// SetAddressCityNil sets the value for AddressCity to be an explicit nil
func (o *Supplier) SetAddressCityNil() *Supplier {
	o.AddressCity.Set(nil)
	return o
}

// UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
func (o *Supplier) UnsetAddressCity() {
	o.AddressCity.Unset()
}

// GetAddressProvince returns the AddressProvince field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetAddressProvince() string {
	if o == nil || IsNil(o.AddressProvince.Get()) {
		var ret string
		return ret
	}
	return *o.AddressProvince.Get()
}

// GetAddressProvinceOk returns a tuple with the AddressProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetAddressProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressProvince.Get(), o.AddressProvince.IsSet()
}

// HasAddressProvince returns a boolean if a field has been set.
func (o *Supplier) HasAddressProvince() bool {
	if o != nil && o.AddressProvince.IsSet() {
		return true
	}

	return false
}

// SetAddressProvince gets a reference to the given NullableString and assigns it to the AddressProvince field.
func (o *Supplier) SetAddressProvince(v string) *Supplier {
	o.AddressProvince.Set(&v)
	return o
}
// SetAddressProvinceNil sets the value for AddressProvince to be an explicit nil
func (o *Supplier) SetAddressProvinceNil() *Supplier {
	o.AddressProvince.Set(nil)
	return o
}

// UnsetAddressProvince ensures that no value is present for AddressProvince, not even an explicit nil
func (o *Supplier) UnsetAddressProvince() {
	o.AddressProvince.Unset()
}

// GetAddressExtra returns the AddressExtra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetAddressExtra() string {
	if o == nil || IsNil(o.AddressExtra.Get()) {
		var ret string
		return ret
	}
	return *o.AddressExtra.Get()
}

// GetAddressExtraOk returns a tuple with the AddressExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetAddressExtraOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressExtra.Get(), o.AddressExtra.IsSet()
}

// HasAddressExtra returns a boolean if a field has been set.
func (o *Supplier) HasAddressExtra() bool {
	if o != nil && o.AddressExtra.IsSet() {
		return true
	}

	return false
}

// SetAddressExtra gets a reference to the given NullableString and assigns it to the AddressExtra field.
func (o *Supplier) SetAddressExtra(v string) *Supplier {
	o.AddressExtra.Set(&v)
	return o
}
// SetAddressExtraNil sets the value for AddressExtra to be an explicit nil
func (o *Supplier) SetAddressExtraNil() *Supplier {
	o.AddressExtra.Set(nil)
	return o
}

// UnsetAddressExtra ensures that no value is present for AddressExtra, not even an explicit nil
func (o *Supplier) UnsetAddressExtra() {
	o.AddressExtra.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *Supplier) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *Supplier) SetCountry(v string) *Supplier {
	o.Country.Set(&v)
	return o
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *Supplier) SetCountryNil() *Supplier {
	o.Country.Set(nil)
	return o
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *Supplier) UnsetCountry() {
	o.Country.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Supplier) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Supplier) SetEmail(v string) *Supplier {
	o.Email.Set(&v)
	return o
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Supplier) SetEmailNil() *Supplier {
	o.Email.Set(nil)
	return o
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Supplier) UnsetEmail() {
	o.Email.Unset()
}

// GetCertifiedEmail returns the CertifiedEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetCertifiedEmail() string {
	if o == nil || IsNil(o.CertifiedEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CertifiedEmail.Get()
}

// GetCertifiedEmailOk returns a tuple with the CertifiedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetCertifiedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertifiedEmail.Get(), o.CertifiedEmail.IsSet()
}

// HasCertifiedEmail returns a boolean if a field has been set.
func (o *Supplier) HasCertifiedEmail() bool {
	if o != nil && o.CertifiedEmail.IsSet() {
		return true
	}

	return false
}

// SetCertifiedEmail gets a reference to the given NullableString and assigns it to the CertifiedEmail field.
func (o *Supplier) SetCertifiedEmail(v string) *Supplier {
	o.CertifiedEmail.Set(&v)
	return o
}
// SetCertifiedEmailNil sets the value for CertifiedEmail to be an explicit nil
func (o *Supplier) SetCertifiedEmailNil() *Supplier {
	o.CertifiedEmail.Set(nil)
	return o
}

// UnsetCertifiedEmail ensures that no value is present for CertifiedEmail, not even an explicit nil
func (o *Supplier) UnsetCertifiedEmail() {
	o.CertifiedEmail.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *Supplier) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *Supplier) SetPhone(v string) *Supplier {
	o.Phone.Set(&v)
	return o
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *Supplier) SetPhoneNil() *Supplier {
	o.Phone.Set(nil)
	return o
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *Supplier) UnsetPhone() {
	o.Phone.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetFax() string {
	if o == nil || IsNil(o.Fax.Get()) {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetFaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *Supplier) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *Supplier) SetFax(v string) *Supplier {
	o.Fax.Set(&v)
	return o
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *Supplier) SetFaxNil() *Supplier {
	o.Fax.Set(nil)
	return o
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *Supplier) UnsetFax() {
	o.Fax.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *Supplier) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *Supplier) SetNotes(v string) *Supplier {
	o.Notes.Set(&v)
	return o
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *Supplier) SetNotesNil() *Supplier {
	o.Notes.Set(nil)
	return o
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *Supplier) UnsetNotes() {
	o.Notes.Unset()
}

// GetBankIban returns the BankIban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetBankIban() string {
	if o == nil || IsNil(o.BankIban.Get()) {
		var ret string
		return ret
	}
	return *o.BankIban.Get()
}

// GetBankIbanOk returns a tuple with the BankIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetBankIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankIban.Get(), o.BankIban.IsSet()
}

// HasBankIban returns a boolean if a field has been set.
func (o *Supplier) HasBankIban() bool {
	if o != nil && o.BankIban.IsSet() {
		return true
	}

	return false
}

// SetBankIban gets a reference to the given NullableString and assigns it to the BankIban field.
func (o *Supplier) SetBankIban(v string) *Supplier {
	o.BankIban.Set(&v)
	return o
}
// SetBankIbanNil sets the value for BankIban to be an explicit nil
func (o *Supplier) SetBankIbanNil() *Supplier {
	o.BankIban.Set(nil)
	return o
}

// UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
func (o *Supplier) UnsetBankIban() {
	o.BankIban.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Supplier) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Supplier) SetCreatedAt(v string) *Supplier {
	o.CreatedAt.Set(&v)
	return o
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Supplier) SetCreatedAtNil() *Supplier {
	o.CreatedAt.Set(nil)
	return o
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Supplier) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Supplier) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Supplier) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Supplier) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *Supplier) SetUpdatedAt(v string) *Supplier {
	o.UpdatedAt.Set(&v)
	return o
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Supplier) SetUpdatedAtNil() *Supplier {
	o.UpdatedAt.Set(nil)
	return o
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Supplier) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Supplier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Supplier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.ContactPerson.IsSet() {
		toSerialize["contact_person"] = o.ContactPerson.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vat_number"] = o.VatNumber.Get()
	}
	if o.TaxCode.IsSet() {
		toSerialize["tax_code"] = o.TaxCode.Get()
	}
	if o.AddressStreet.IsSet() {
		toSerialize["address_street"] = o.AddressStreet.Get()
	}
	if o.AddressPostalCode.IsSet() {
		toSerialize["address_postal_code"] = o.AddressPostalCode.Get()
	}
	if o.AddressCity.IsSet() {
		toSerialize["address_city"] = o.AddressCity.Get()
	}
	if o.AddressProvince.IsSet() {
		toSerialize["address_province"] = o.AddressProvince.Get()
	}
	if o.AddressExtra.IsSet() {
		toSerialize["address_extra"] = o.AddressExtra.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.CertifiedEmail.IsSet() {
		toSerialize["certified_email"] = o.CertifiedEmail.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Fax.IsSet() {
		toSerialize["fax"] = o.Fax.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.BankIban.IsSet() {
		toSerialize["bank_iban"] = o.BankIban.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableSupplier struct {
	value *Supplier
	isSet bool
}

func (v NullableSupplier) Get() *Supplier {
	return v.value
}

func (v *NullableSupplier) Set(val *Supplier) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplier) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplier(val *Supplier) *NullableSupplier {
	return &NullableSupplier{value: val, isSet: true}
}

func (v NullableSupplier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


