# Supplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Supplier id | [optional] 
**Code** | Pointer to **NullableString** | Supplier code | [optional] 
**Name** | Pointer to **NullableString** | Supplier name | [optional] 
**Type** | Pointer to [**NullableSupplierType**](SupplierType.md) |  | [optional] 
**FirstName** | Pointer to **NullableString** | Supplier first name | [optional] 
**LastName** | Pointer to **NullableString** | Supplier last name | [optional] 
**ContactPerson** | Pointer to **NullableString** | Supplier contact person | [optional] 
**VatNumber** | Pointer to **NullableString** | Supplier vat number | [optional] 
**TaxCode** | Pointer to **NullableString** | Supplier tax code | [optional] 
**AddressStreet** | Pointer to **NullableString** | Supplier street address | [optional] 
**AddressPostalCode** | Pointer to **NullableString** | Supplier postal code | [optional] 
**AddressCity** | Pointer to **NullableString** | Supplier city | [optional] 
**AddressProvince** | Pointer to **NullableString** | Supplier province | [optional] 
**AddressExtra** | Pointer to **NullableString** | Supplier address extra info | [optional] 
**Country** | Pointer to **NullableString** | Supplier country | [optional] 
**CountryIso** | Pointer to **NullableString** | Supplier country iso code | [optional] 
**Email** | Pointer to **NullableString** | Supplier email | [optional] 
**CertifiedEmail** | Pointer to **NullableString** | Supplier certified email | [optional] 
**Phone** | Pointer to **NullableString** | Supplier phone | [optional] 
**Fax** | Pointer to **NullableString** | Supplier fax | [optional] 
**Notes** | Pointer to **NullableString** | Supplier extra notes | [optional] 
**BankIban** | Pointer to **NullableString** | Supplier bank IBAN | [optional] 
**CreatedAt** | Pointer to **NullableString** | Supplier creation date | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Supplier last update date | [optional] 

## Methods

### NewSupplier

`func NewSupplier() *Supplier`

NewSupplier instantiates a new Supplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierWithDefaults

`func NewSupplierWithDefaults() *Supplier`

NewSupplierWithDefaults instantiates a new Supplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Supplier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Supplier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Supplier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Supplier) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Supplier) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Supplier) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *Supplier) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Supplier) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Supplier) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Supplier) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Supplier) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Supplier) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *Supplier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Supplier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Supplier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Supplier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Supplier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Supplier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Supplier) GetType() SupplierType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Supplier) GetTypeOk() (*SupplierType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Supplier) SetType(v SupplierType)`

SetType sets Type field to given value.

### HasType

`func (o *Supplier) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Supplier) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Supplier) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFirstName

`func (o *Supplier) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Supplier) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Supplier) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Supplier) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Supplier) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Supplier) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Supplier) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Supplier) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Supplier) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Supplier) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Supplier) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Supplier) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetContactPerson

`func (o *Supplier) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *Supplier) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *Supplier) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *Supplier) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### SetContactPersonNil

`func (o *Supplier) SetContactPersonNil(b bool)`

 SetContactPersonNil sets the value for ContactPerson to be an explicit nil

### UnsetContactPerson
`func (o *Supplier) UnsetContactPerson()`

UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
### GetVatNumber

`func (o *Supplier) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Supplier) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Supplier) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Supplier) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *Supplier) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *Supplier) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetTaxCode

`func (o *Supplier) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Supplier) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Supplier) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Supplier) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *Supplier) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *Supplier) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetAddressStreet

`func (o *Supplier) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *Supplier) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *Supplier) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *Supplier) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### SetAddressStreetNil

`func (o *Supplier) SetAddressStreetNil(b bool)`

 SetAddressStreetNil sets the value for AddressStreet to be an explicit nil

### UnsetAddressStreet
`func (o *Supplier) UnsetAddressStreet()`

UnsetAddressStreet ensures that no value is present for AddressStreet, not even an explicit nil
### GetAddressPostalCode

`func (o *Supplier) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *Supplier) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *Supplier) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *Supplier) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### SetAddressPostalCodeNil

`func (o *Supplier) SetAddressPostalCodeNil(b bool)`

 SetAddressPostalCodeNil sets the value for AddressPostalCode to be an explicit nil

### UnsetAddressPostalCode
`func (o *Supplier) UnsetAddressPostalCode()`

UnsetAddressPostalCode ensures that no value is present for AddressPostalCode, not even an explicit nil
### GetAddressCity

`func (o *Supplier) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Supplier) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Supplier) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Supplier) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *Supplier) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *Supplier) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressProvince

`func (o *Supplier) GetAddressProvince() string`

GetAddressProvince returns the AddressProvince field if non-nil, zero value otherwise.

### GetAddressProvinceOk

`func (o *Supplier) GetAddressProvinceOk() (*string, bool)`

GetAddressProvinceOk returns a tuple with the AddressProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressProvince

`func (o *Supplier) SetAddressProvince(v string)`

SetAddressProvince sets AddressProvince field to given value.

### HasAddressProvince

`func (o *Supplier) HasAddressProvince() bool`

HasAddressProvince returns a boolean if a field has been set.

### SetAddressProvinceNil

`func (o *Supplier) SetAddressProvinceNil(b bool)`

 SetAddressProvinceNil sets the value for AddressProvince to be an explicit nil

### UnsetAddressProvince
`func (o *Supplier) UnsetAddressProvince()`

UnsetAddressProvince ensures that no value is present for AddressProvince, not even an explicit nil
### GetAddressExtra

`func (o *Supplier) GetAddressExtra() string`

GetAddressExtra returns the AddressExtra field if non-nil, zero value otherwise.

### GetAddressExtraOk

`func (o *Supplier) GetAddressExtraOk() (*string, bool)`

GetAddressExtraOk returns a tuple with the AddressExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressExtra

`func (o *Supplier) SetAddressExtra(v string)`

SetAddressExtra sets AddressExtra field to given value.

### HasAddressExtra

`func (o *Supplier) HasAddressExtra() bool`

HasAddressExtra returns a boolean if a field has been set.

### SetAddressExtraNil

`func (o *Supplier) SetAddressExtraNil(b bool)`

 SetAddressExtraNil sets the value for AddressExtra to be an explicit nil

### UnsetAddressExtra
`func (o *Supplier) UnsetAddressExtra()`

UnsetAddressExtra ensures that no value is present for AddressExtra, not even an explicit nil
### GetCountry

`func (o *Supplier) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Supplier) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Supplier) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Supplier) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Supplier) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Supplier) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryIso

`func (o *Supplier) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *Supplier) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *Supplier) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *Supplier) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### SetCountryIsoNil

`func (o *Supplier) SetCountryIsoNil(b bool)`

 SetCountryIsoNil sets the value for CountryIso to be an explicit nil

### UnsetCountryIso
`func (o *Supplier) UnsetCountryIso()`

UnsetCountryIso ensures that no value is present for CountryIso, not even an explicit nil
### GetEmail

`func (o *Supplier) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Supplier) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Supplier) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Supplier) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Supplier) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Supplier) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCertifiedEmail

`func (o *Supplier) GetCertifiedEmail() string`

GetCertifiedEmail returns the CertifiedEmail field if non-nil, zero value otherwise.

### GetCertifiedEmailOk

`func (o *Supplier) GetCertifiedEmailOk() (*string, bool)`

GetCertifiedEmailOk returns a tuple with the CertifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedEmail

`func (o *Supplier) SetCertifiedEmail(v string)`

SetCertifiedEmail sets CertifiedEmail field to given value.

### HasCertifiedEmail

`func (o *Supplier) HasCertifiedEmail() bool`

HasCertifiedEmail returns a boolean if a field has been set.

### SetCertifiedEmailNil

`func (o *Supplier) SetCertifiedEmailNil(b bool)`

 SetCertifiedEmailNil sets the value for CertifiedEmail to be an explicit nil

### UnsetCertifiedEmail
`func (o *Supplier) UnsetCertifiedEmail()`

UnsetCertifiedEmail ensures that no value is present for CertifiedEmail, not even an explicit nil
### GetPhone

`func (o *Supplier) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Supplier) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Supplier) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Supplier) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Supplier) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Supplier) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *Supplier) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Supplier) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Supplier) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Supplier) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *Supplier) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *Supplier) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetNotes

`func (o *Supplier) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Supplier) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Supplier) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Supplier) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Supplier) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Supplier) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetBankIban

`func (o *Supplier) GetBankIban() string`

GetBankIban returns the BankIban field if non-nil, zero value otherwise.

### GetBankIbanOk

`func (o *Supplier) GetBankIbanOk() (*string, bool)`

GetBankIbanOk returns a tuple with the BankIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIban

`func (o *Supplier) SetBankIban(v string)`

SetBankIban sets BankIban field to given value.

### HasBankIban

`func (o *Supplier) HasBankIban() bool`

HasBankIban returns a boolean if a field has been set.

### SetBankIbanNil

`func (o *Supplier) SetBankIbanNil(b bool)`

 SetBankIbanNil sets the value for BankIban to be an explicit nil

### UnsetBankIban
`func (o *Supplier) UnsetBankIban()`

UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
### GetCreatedAt

`func (o *Supplier) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Supplier) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Supplier) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Supplier) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Supplier) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Supplier) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Supplier) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Supplier) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Supplier) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Supplier) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Supplier) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Supplier) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


