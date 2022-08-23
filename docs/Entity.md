# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier | [optional] 
**Code** | Pointer to **NullableString** | Code. | [optional] 
**Name** | Pointer to **NullableString** | Name | [optional] 
**Type** | Pointer to [**NullableEntityType**](EntityType.md) |  | [optional] 
**FirstName** | Pointer to **NullableString** | First name. | [optional] 
**LastName** | Pointer to **NullableString** | Last name. | [optional] 
**ContactPerson** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** | Vat number | [optional] 
**TaxCode** | Pointer to **NullableString** | Tax code. | [optional] 
**AddressStreet** | Pointer to **NullableString** | Street address. | [optional] 
**AddressPostalCode** | Pointer to **NullableString** | Postal code. | [optional] 
**AddressCity** | Pointer to **NullableString** | City. | [optional] 
**AddressProvince** | Pointer to **NullableString** | Province. | [optional] 
**AddressExtra** | Pointer to **NullableString** | Address extra info. | [optional] 
**Country** | Pointer to **NullableString** | Country | [optional] 
**Email** | Pointer to **NullableString** | Email. | [optional] 
**CertifiedEmail** | Pointer to **NullableString** | Certified email. | [optional] 
**Phone** | Pointer to **NullableString** | Phone. | [optional] 
**Fax** | Pointer to **NullableString** | Fax. | [optional] 
**Notes** | Pointer to **NullableString** | Extra notes. | [optional] 
**DefaultVat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 
**DefaultPaymentTerms** | Pointer to **NullableInt32** | [Only for client] Default payment terms. | [optional] 
**DefaultPaymentTermsType** | Pointer to [**DefaultPaymentTermsType**](DefaultPaymentTermsType.md) |  | [optional] [default to STANDARD]
**DefaultPaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 
**BankName** | Pointer to **NullableString** | [Only for client] Bank name. | [optional] 
**BankIban** | Pointer to **NullableString** | [Only for client] Iban. | [optional] 
**BankSwiftCode** | Pointer to **NullableString** | [Only for client] Bank swift code. | [optional] 
**ShippingAddress** | Pointer to **NullableString** | [Only for client] Shipping address. | [optional] 
**EInvoice** | Pointer to **NullableBool** | [Only for client] Use e-invoices. | [optional] 
**EiCode** | Pointer to **NullableString** | [Only for client] E-invoices code. | [optional] 
**HasIntentDeclaration** | Pointer to **NullableBool** | [Only for client] Has intent declaration. | [optional] 
**IntentDeclarationProtocolNumber** | Pointer to **NullableString** | [Only for client] Intent declaration protocol number. | [optional] 
**IntentDeclarationProtocolDate** | Pointer to **NullableString** | [Only for client] Intent declaration protocol date. | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Entity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Entity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Entity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *Entity) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Entity) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Entity) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Entity) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Entity) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Entity) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *Entity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Entity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Entity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Entity) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entity) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entity) SetType(v EntityType)`

SetType sets Type field to given value.

### HasType

`func (o *Entity) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Entity) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Entity) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFirstName

`func (o *Entity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Entity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Entity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Entity) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Entity) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Entity) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Entity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Entity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Entity) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Entity) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Entity) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Entity) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetContactPerson

`func (o *Entity) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *Entity) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *Entity) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *Entity) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### SetContactPersonNil

`func (o *Entity) SetContactPersonNil(b bool)`

 SetContactPersonNil sets the value for ContactPerson to be an explicit nil

### UnsetContactPerson
`func (o *Entity) UnsetContactPerson()`

UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
### GetVatNumber

`func (o *Entity) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Entity) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Entity) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Entity) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *Entity) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *Entity) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetTaxCode

`func (o *Entity) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Entity) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Entity) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Entity) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *Entity) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *Entity) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetAddressStreet

`func (o *Entity) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *Entity) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *Entity) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *Entity) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### SetAddressStreetNil

`func (o *Entity) SetAddressStreetNil(b bool)`

 SetAddressStreetNil sets the value for AddressStreet to be an explicit nil

### UnsetAddressStreet
`func (o *Entity) UnsetAddressStreet()`

UnsetAddressStreet ensures that no value is present for AddressStreet, not even an explicit nil
### GetAddressPostalCode

`func (o *Entity) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *Entity) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *Entity) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *Entity) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### SetAddressPostalCodeNil

`func (o *Entity) SetAddressPostalCodeNil(b bool)`

 SetAddressPostalCodeNil sets the value for AddressPostalCode to be an explicit nil

### UnsetAddressPostalCode
`func (o *Entity) UnsetAddressPostalCode()`

UnsetAddressPostalCode ensures that no value is present for AddressPostalCode, not even an explicit nil
### GetAddressCity

`func (o *Entity) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Entity) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Entity) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Entity) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *Entity) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *Entity) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressProvince

`func (o *Entity) GetAddressProvince() string`

GetAddressProvince returns the AddressProvince field if non-nil, zero value otherwise.

### GetAddressProvinceOk

`func (o *Entity) GetAddressProvinceOk() (*string, bool)`

GetAddressProvinceOk returns a tuple with the AddressProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressProvince

`func (o *Entity) SetAddressProvince(v string)`

SetAddressProvince sets AddressProvince field to given value.

### HasAddressProvince

`func (o *Entity) HasAddressProvince() bool`

HasAddressProvince returns a boolean if a field has been set.

### SetAddressProvinceNil

`func (o *Entity) SetAddressProvinceNil(b bool)`

 SetAddressProvinceNil sets the value for AddressProvince to be an explicit nil

### UnsetAddressProvince
`func (o *Entity) UnsetAddressProvince()`

UnsetAddressProvince ensures that no value is present for AddressProvince, not even an explicit nil
### GetAddressExtra

`func (o *Entity) GetAddressExtra() string`

GetAddressExtra returns the AddressExtra field if non-nil, zero value otherwise.

### GetAddressExtraOk

`func (o *Entity) GetAddressExtraOk() (*string, bool)`

GetAddressExtraOk returns a tuple with the AddressExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressExtra

`func (o *Entity) SetAddressExtra(v string)`

SetAddressExtra sets AddressExtra field to given value.

### HasAddressExtra

`func (o *Entity) HasAddressExtra() bool`

HasAddressExtra returns a boolean if a field has been set.

### SetAddressExtraNil

`func (o *Entity) SetAddressExtraNil(b bool)`

 SetAddressExtraNil sets the value for AddressExtra to be an explicit nil

### UnsetAddressExtra
`func (o *Entity) UnsetAddressExtra()`

UnsetAddressExtra ensures that no value is present for AddressExtra, not even an explicit nil
### GetCountry

`func (o *Entity) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Entity) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Entity) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Entity) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Entity) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Entity) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetEmail

`func (o *Entity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Entity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Entity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Entity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Entity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Entity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCertifiedEmail

`func (o *Entity) GetCertifiedEmail() string`

GetCertifiedEmail returns the CertifiedEmail field if non-nil, zero value otherwise.

### GetCertifiedEmailOk

`func (o *Entity) GetCertifiedEmailOk() (*string, bool)`

GetCertifiedEmailOk returns a tuple with the CertifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedEmail

`func (o *Entity) SetCertifiedEmail(v string)`

SetCertifiedEmail sets CertifiedEmail field to given value.

### HasCertifiedEmail

`func (o *Entity) HasCertifiedEmail() bool`

HasCertifiedEmail returns a boolean if a field has been set.

### SetCertifiedEmailNil

`func (o *Entity) SetCertifiedEmailNil(b bool)`

 SetCertifiedEmailNil sets the value for CertifiedEmail to be an explicit nil

### UnsetCertifiedEmail
`func (o *Entity) UnsetCertifiedEmail()`

UnsetCertifiedEmail ensures that no value is present for CertifiedEmail, not even an explicit nil
### GetPhone

`func (o *Entity) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Entity) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Entity) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Entity) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Entity) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Entity) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *Entity) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Entity) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Entity) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Entity) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *Entity) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *Entity) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetNotes

`func (o *Entity) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Entity) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Entity) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Entity) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Entity) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Entity) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDefaultVat

`func (o *Entity) GetDefaultVat() VatType`

GetDefaultVat returns the DefaultVat field if non-nil, zero value otherwise.

### GetDefaultVatOk

`func (o *Entity) GetDefaultVatOk() (*VatType, bool)`

GetDefaultVatOk returns a tuple with the DefaultVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVat

`func (o *Entity) SetDefaultVat(v VatType)`

SetDefaultVat sets DefaultVat field to given value.

### HasDefaultVat

`func (o *Entity) HasDefaultVat() bool`

HasDefaultVat returns a boolean if a field has been set.

### SetDefaultVatNil

`func (o *Entity) SetDefaultVatNil(b bool)`

 SetDefaultVatNil sets the value for DefaultVat to be an explicit nil

### UnsetDefaultVat
`func (o *Entity) UnsetDefaultVat()`

UnsetDefaultVat ensures that no value is present for DefaultVat, not even an explicit nil
### GetDefaultPaymentTerms

`func (o *Entity) GetDefaultPaymentTerms() int32`

GetDefaultPaymentTerms returns the DefaultPaymentTerms field if non-nil, zero value otherwise.

### GetDefaultPaymentTermsOk

`func (o *Entity) GetDefaultPaymentTermsOk() (*int32, bool)`

GetDefaultPaymentTermsOk returns a tuple with the DefaultPaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentTerms

`func (o *Entity) SetDefaultPaymentTerms(v int32)`

SetDefaultPaymentTerms sets DefaultPaymentTerms field to given value.

### HasDefaultPaymentTerms

`func (o *Entity) HasDefaultPaymentTerms() bool`

HasDefaultPaymentTerms returns a boolean if a field has been set.

### SetDefaultPaymentTermsNil

`func (o *Entity) SetDefaultPaymentTermsNil(b bool)`

 SetDefaultPaymentTermsNil sets the value for DefaultPaymentTerms to be an explicit nil

### UnsetDefaultPaymentTerms
`func (o *Entity) UnsetDefaultPaymentTerms()`

UnsetDefaultPaymentTerms ensures that no value is present for DefaultPaymentTerms, not even an explicit nil
### GetDefaultPaymentTermsType

`func (o *Entity) GetDefaultPaymentTermsType() DefaultPaymentTermsType`

GetDefaultPaymentTermsType returns the DefaultPaymentTermsType field if non-nil, zero value otherwise.

### GetDefaultPaymentTermsTypeOk

`func (o *Entity) GetDefaultPaymentTermsTypeOk() (*DefaultPaymentTermsType, bool)`

GetDefaultPaymentTermsTypeOk returns a tuple with the DefaultPaymentTermsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentTermsType

`func (o *Entity) SetDefaultPaymentTermsType(v DefaultPaymentTermsType)`

SetDefaultPaymentTermsType sets DefaultPaymentTermsType field to given value.

### HasDefaultPaymentTermsType

`func (o *Entity) HasDefaultPaymentTermsType() bool`

HasDefaultPaymentTermsType returns a boolean if a field has been set.

### GetDefaultPaymentMethod

`func (o *Entity) GetDefaultPaymentMethod() PaymentMethod`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *Entity) GetDefaultPaymentMethodOk() (*PaymentMethod, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *Entity) SetDefaultPaymentMethod(v PaymentMethod)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.

### HasDefaultPaymentMethod

`func (o *Entity) HasDefaultPaymentMethod() bool`

HasDefaultPaymentMethod returns a boolean if a field has been set.

### GetBankName

`func (o *Entity) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Entity) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Entity) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *Entity) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *Entity) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *Entity) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankIban

`func (o *Entity) GetBankIban() string`

GetBankIban returns the BankIban field if non-nil, zero value otherwise.

### GetBankIbanOk

`func (o *Entity) GetBankIbanOk() (*string, bool)`

GetBankIbanOk returns a tuple with the BankIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIban

`func (o *Entity) SetBankIban(v string)`

SetBankIban sets BankIban field to given value.

### HasBankIban

`func (o *Entity) HasBankIban() bool`

HasBankIban returns a boolean if a field has been set.

### SetBankIbanNil

`func (o *Entity) SetBankIbanNil(b bool)`

 SetBankIbanNil sets the value for BankIban to be an explicit nil

### UnsetBankIban
`func (o *Entity) UnsetBankIban()`

UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
### GetBankSwiftCode

`func (o *Entity) GetBankSwiftCode() string`

GetBankSwiftCode returns the BankSwiftCode field if non-nil, zero value otherwise.

### GetBankSwiftCodeOk

`func (o *Entity) GetBankSwiftCodeOk() (*string, bool)`

GetBankSwiftCodeOk returns a tuple with the BankSwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankSwiftCode

`func (o *Entity) SetBankSwiftCode(v string)`

SetBankSwiftCode sets BankSwiftCode field to given value.

### HasBankSwiftCode

`func (o *Entity) HasBankSwiftCode() bool`

HasBankSwiftCode returns a boolean if a field has been set.

### SetBankSwiftCodeNil

`func (o *Entity) SetBankSwiftCodeNil(b bool)`

 SetBankSwiftCodeNil sets the value for BankSwiftCode to be an explicit nil

### UnsetBankSwiftCode
`func (o *Entity) UnsetBankSwiftCode()`

UnsetBankSwiftCode ensures that no value is present for BankSwiftCode, not even an explicit nil
### GetShippingAddress

`func (o *Entity) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Entity) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Entity) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Entity) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### SetShippingAddressNil

`func (o *Entity) SetShippingAddressNil(b bool)`

 SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil

### UnsetShippingAddress
`func (o *Entity) UnsetShippingAddress()`

UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil
### GetEInvoice

`func (o *Entity) GetEInvoice() bool`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *Entity) GetEInvoiceOk() (*bool, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *Entity) SetEInvoice(v bool)`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *Entity) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### SetEInvoiceNil

`func (o *Entity) SetEInvoiceNil(b bool)`

 SetEInvoiceNil sets the value for EInvoice to be an explicit nil

### UnsetEInvoice
`func (o *Entity) UnsetEInvoice()`

UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
### GetEiCode

`func (o *Entity) GetEiCode() string`

GetEiCode returns the EiCode field if non-nil, zero value otherwise.

### GetEiCodeOk

`func (o *Entity) GetEiCodeOk() (*string, bool)`

GetEiCodeOk returns a tuple with the EiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiCode

`func (o *Entity) SetEiCode(v string)`

SetEiCode sets EiCode field to given value.

### HasEiCode

`func (o *Entity) HasEiCode() bool`

HasEiCode returns a boolean if a field has been set.

### SetEiCodeNil

`func (o *Entity) SetEiCodeNil(b bool)`

 SetEiCodeNil sets the value for EiCode to be an explicit nil

### UnsetEiCode
`func (o *Entity) UnsetEiCode()`

UnsetEiCode ensures that no value is present for EiCode, not even an explicit nil
### GetHasIntentDeclaration

`func (o *Entity) GetHasIntentDeclaration() bool`

GetHasIntentDeclaration returns the HasIntentDeclaration field if non-nil, zero value otherwise.

### GetHasIntentDeclarationOk

`func (o *Entity) GetHasIntentDeclarationOk() (*bool, bool)`

GetHasIntentDeclarationOk returns a tuple with the HasIntentDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIntentDeclaration

`func (o *Entity) SetHasIntentDeclaration(v bool)`

SetHasIntentDeclaration sets HasIntentDeclaration field to given value.

### HasHasIntentDeclaration

`func (o *Entity) HasHasIntentDeclaration() bool`

HasHasIntentDeclaration returns a boolean if a field has been set.

### SetHasIntentDeclarationNil

`func (o *Entity) SetHasIntentDeclarationNil(b bool)`

 SetHasIntentDeclarationNil sets the value for HasIntentDeclaration to be an explicit nil

### UnsetHasIntentDeclaration
`func (o *Entity) UnsetHasIntentDeclaration()`

UnsetHasIntentDeclaration ensures that no value is present for HasIntentDeclaration, not even an explicit nil
### GetIntentDeclarationProtocolNumber

`func (o *Entity) GetIntentDeclarationProtocolNumber() string`

GetIntentDeclarationProtocolNumber returns the IntentDeclarationProtocolNumber field if non-nil, zero value otherwise.

### GetIntentDeclarationProtocolNumberOk

`func (o *Entity) GetIntentDeclarationProtocolNumberOk() (*string, bool)`

GetIntentDeclarationProtocolNumberOk returns a tuple with the IntentDeclarationProtocolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentDeclarationProtocolNumber

`func (o *Entity) SetIntentDeclarationProtocolNumber(v string)`

SetIntentDeclarationProtocolNumber sets IntentDeclarationProtocolNumber field to given value.

### HasIntentDeclarationProtocolNumber

`func (o *Entity) HasIntentDeclarationProtocolNumber() bool`

HasIntentDeclarationProtocolNumber returns a boolean if a field has been set.

### SetIntentDeclarationProtocolNumberNil

`func (o *Entity) SetIntentDeclarationProtocolNumberNil(b bool)`

 SetIntentDeclarationProtocolNumberNil sets the value for IntentDeclarationProtocolNumber to be an explicit nil

### UnsetIntentDeclarationProtocolNumber
`func (o *Entity) UnsetIntentDeclarationProtocolNumber()`

UnsetIntentDeclarationProtocolNumber ensures that no value is present for IntentDeclarationProtocolNumber, not even an explicit nil
### GetIntentDeclarationProtocolDate

`func (o *Entity) GetIntentDeclarationProtocolDate() string`

GetIntentDeclarationProtocolDate returns the IntentDeclarationProtocolDate field if non-nil, zero value otherwise.

### GetIntentDeclarationProtocolDateOk

`func (o *Entity) GetIntentDeclarationProtocolDateOk() (*string, bool)`

GetIntentDeclarationProtocolDateOk returns a tuple with the IntentDeclarationProtocolDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentDeclarationProtocolDate

`func (o *Entity) SetIntentDeclarationProtocolDate(v string)`

SetIntentDeclarationProtocolDate sets IntentDeclarationProtocolDate field to given value.

### HasIntentDeclarationProtocolDate

`func (o *Entity) HasIntentDeclarationProtocolDate() bool`

HasIntentDeclarationProtocolDate returns a boolean if a field has been set.

### SetIntentDeclarationProtocolDateNil

`func (o *Entity) SetIntentDeclarationProtocolDateNil(b bool)`

 SetIntentDeclarationProtocolDateNil sets the value for IntentDeclarationProtocolDate to be an explicit nil

### UnsetIntentDeclarationProtocolDate
`func (o *Entity) UnsetIntentDeclarationProtocolDate()`

UnsetIntentDeclarationProtocolDate ensures that no value is present for IntentDeclarationProtocolDate, not even an explicit nil
### GetCreatedAt

`func (o *Entity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Entity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Entity) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Entity) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Entity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Entity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Entity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Entity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Entity) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Entity) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


