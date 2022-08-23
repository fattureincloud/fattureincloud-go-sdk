# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique identifier | [optional] 
**Code** | Pointer to **NullableString** | Client code. | [optional] 
**Name** | Pointer to **NullableString** | Client name | [optional] 
**Type** | Pointer to [**NullableClientType**](ClientType.md) |  | [optional] 
**FirstName** | Pointer to **NullableString** | Client first name. | [optional] 
**LastName** | Pointer to **NullableString** | Client last name. | [optional] 
**ContactPerson** | Pointer to **NullableString** |  | [optional] 
**VatNumber** | Pointer to **NullableString** | Client vat number | [optional] 
**TaxCode** | Pointer to **NullableString** | Client tax code. | [optional] 
**AddressStreet** | Pointer to **NullableString** | Client street address. | [optional] 
**AddressPostalCode** | Pointer to **NullableString** | Client postal code. | [optional] 
**AddressCity** | Pointer to **NullableString** | Client city. | [optional] 
**AddressProvince** | Pointer to **NullableString** | Client province. | [optional] 
**AddressExtra** | Pointer to **NullableString** | Client address extra info. | [optional] 
**Country** | Pointer to **NullableString** | Client country | [optional] 
**Email** | Pointer to **NullableString** | Client email. | [optional] 
**CertifiedEmail** | Pointer to **NullableString** | Client certified email. | [optional] 
**Phone** | Pointer to **NullableString** | Client phone. | [optional] 
**Fax** | Pointer to **NullableString** | Client fax. | [optional] 
**Notes** | Pointer to **NullableString** | Extra notes. | [optional] 
**DefaultVat** | Pointer to [**NullableVatType**](VatType.md) |  | [optional] 
**DefaultPaymentTerms** | Pointer to **NullableInt32** |  | [optional] 
**DefaultPaymentTermsType** | Pointer to [**DefaultPaymentTermsType**](DefaultPaymentTermsType.md) |  | [optional] [default to STANDARD]
**DefaultPaymentMethod** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  | [optional] 
**BankName** | Pointer to **NullableString** | Client bank name. | [optional] 
**BankIban** | Pointer to **NullableString** | Client iban. | [optional] 
**BankSwiftCode** | Pointer to **NullableString** | Client bank swift code. | [optional] 
**ShippingAddress** | Pointer to **NullableString** | Client shipping address. | [optional] 
**EInvoice** | Pointer to **NullableBool** | Use e-invoices for this entity | [optional] 
**EiCode** | Pointer to **NullableString** | E-invoice code | [optional] 
**DiscountHighlight** | Pointer to **NullableBool** | Discount Highlight. | [optional] 
**DefaultDiscount** | Pointer to **NullableFloat32** | Default discount. | [optional] 
**HasIntentDeclaration** | Pointer to **NullableBool** | Has intent declaration. | [optional] 
**IntentDeclarationProtocolNumber** | Pointer to **NullableString** | Intent declaration protocol number. | [optional] 
**IntentDeclarationProtocolDate** | Pointer to **NullableString** | Intent declaration protocol date. | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Client) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Client) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Client) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Client) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Client) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Client) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *Client) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Client) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Client) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Client) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Client) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Client) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Client) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Client) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Client) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Client) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Client) GetType() ClientType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Client) GetTypeOk() (*ClientType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Client) SetType(v ClientType)`

SetType sets Type field to given value.

### HasType

`func (o *Client) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Client) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Client) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFirstName

`func (o *Client) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Client) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Client) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Client) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Client) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Client) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Client) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Client) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Client) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Client) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Client) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Client) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetContactPerson

`func (o *Client) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *Client) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *Client) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *Client) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### SetContactPersonNil

`func (o *Client) SetContactPersonNil(b bool)`

 SetContactPersonNil sets the value for ContactPerson to be an explicit nil

### UnsetContactPerson
`func (o *Client) UnsetContactPerson()`

UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
### GetVatNumber

`func (o *Client) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Client) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Client) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Client) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *Client) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *Client) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetTaxCode

`func (o *Client) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Client) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Client) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Client) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *Client) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *Client) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetAddressStreet

`func (o *Client) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *Client) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *Client) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *Client) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### SetAddressStreetNil

`func (o *Client) SetAddressStreetNil(b bool)`

 SetAddressStreetNil sets the value for AddressStreet to be an explicit nil

### UnsetAddressStreet
`func (o *Client) UnsetAddressStreet()`

UnsetAddressStreet ensures that no value is present for AddressStreet, not even an explicit nil
### GetAddressPostalCode

`func (o *Client) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *Client) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *Client) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *Client) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### SetAddressPostalCodeNil

`func (o *Client) SetAddressPostalCodeNil(b bool)`

 SetAddressPostalCodeNil sets the value for AddressPostalCode to be an explicit nil

### UnsetAddressPostalCode
`func (o *Client) UnsetAddressPostalCode()`

UnsetAddressPostalCode ensures that no value is present for AddressPostalCode, not even an explicit nil
### GetAddressCity

`func (o *Client) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Client) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Client) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Client) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### SetAddressCityNil

`func (o *Client) SetAddressCityNil(b bool)`

 SetAddressCityNil sets the value for AddressCity to be an explicit nil

### UnsetAddressCity
`func (o *Client) UnsetAddressCity()`

UnsetAddressCity ensures that no value is present for AddressCity, not even an explicit nil
### GetAddressProvince

`func (o *Client) GetAddressProvince() string`

GetAddressProvince returns the AddressProvince field if non-nil, zero value otherwise.

### GetAddressProvinceOk

`func (o *Client) GetAddressProvinceOk() (*string, bool)`

GetAddressProvinceOk returns a tuple with the AddressProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressProvince

`func (o *Client) SetAddressProvince(v string)`

SetAddressProvince sets AddressProvince field to given value.

### HasAddressProvince

`func (o *Client) HasAddressProvince() bool`

HasAddressProvince returns a boolean if a field has been set.

### SetAddressProvinceNil

`func (o *Client) SetAddressProvinceNil(b bool)`

 SetAddressProvinceNil sets the value for AddressProvince to be an explicit nil

### UnsetAddressProvince
`func (o *Client) UnsetAddressProvince()`

UnsetAddressProvince ensures that no value is present for AddressProvince, not even an explicit nil
### GetAddressExtra

`func (o *Client) GetAddressExtra() string`

GetAddressExtra returns the AddressExtra field if non-nil, zero value otherwise.

### GetAddressExtraOk

`func (o *Client) GetAddressExtraOk() (*string, bool)`

GetAddressExtraOk returns a tuple with the AddressExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressExtra

`func (o *Client) SetAddressExtra(v string)`

SetAddressExtra sets AddressExtra field to given value.

### HasAddressExtra

`func (o *Client) HasAddressExtra() bool`

HasAddressExtra returns a boolean if a field has been set.

### SetAddressExtraNil

`func (o *Client) SetAddressExtraNil(b bool)`

 SetAddressExtraNil sets the value for AddressExtra to be an explicit nil

### UnsetAddressExtra
`func (o *Client) UnsetAddressExtra()`

UnsetAddressExtra ensures that no value is present for AddressExtra, not even an explicit nil
### GetCountry

`func (o *Client) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Client) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Client) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Client) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Client) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Client) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetEmail

`func (o *Client) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Client) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Client) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Client) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Client) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Client) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCertifiedEmail

`func (o *Client) GetCertifiedEmail() string`

GetCertifiedEmail returns the CertifiedEmail field if non-nil, zero value otherwise.

### GetCertifiedEmailOk

`func (o *Client) GetCertifiedEmailOk() (*string, bool)`

GetCertifiedEmailOk returns a tuple with the CertifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedEmail

`func (o *Client) SetCertifiedEmail(v string)`

SetCertifiedEmail sets CertifiedEmail field to given value.

### HasCertifiedEmail

`func (o *Client) HasCertifiedEmail() bool`

HasCertifiedEmail returns a boolean if a field has been set.

### SetCertifiedEmailNil

`func (o *Client) SetCertifiedEmailNil(b bool)`

 SetCertifiedEmailNil sets the value for CertifiedEmail to be an explicit nil

### UnsetCertifiedEmail
`func (o *Client) UnsetCertifiedEmail()`

UnsetCertifiedEmail ensures that no value is present for CertifiedEmail, not even an explicit nil
### GetPhone

`func (o *Client) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Client) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Client) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Client) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Client) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Client) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *Client) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Client) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Client) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Client) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *Client) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *Client) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetNotes

`func (o *Client) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Client) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Client) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Client) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Client) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Client) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDefaultVat

`func (o *Client) GetDefaultVat() VatType`

GetDefaultVat returns the DefaultVat field if non-nil, zero value otherwise.

### GetDefaultVatOk

`func (o *Client) GetDefaultVatOk() (*VatType, bool)`

GetDefaultVatOk returns a tuple with the DefaultVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVat

`func (o *Client) SetDefaultVat(v VatType)`

SetDefaultVat sets DefaultVat field to given value.

### HasDefaultVat

`func (o *Client) HasDefaultVat() bool`

HasDefaultVat returns a boolean if a field has been set.

### SetDefaultVatNil

`func (o *Client) SetDefaultVatNil(b bool)`

 SetDefaultVatNil sets the value for DefaultVat to be an explicit nil

### UnsetDefaultVat
`func (o *Client) UnsetDefaultVat()`

UnsetDefaultVat ensures that no value is present for DefaultVat, not even an explicit nil
### GetDefaultPaymentTerms

`func (o *Client) GetDefaultPaymentTerms() int32`

GetDefaultPaymentTerms returns the DefaultPaymentTerms field if non-nil, zero value otherwise.

### GetDefaultPaymentTermsOk

`func (o *Client) GetDefaultPaymentTermsOk() (*int32, bool)`

GetDefaultPaymentTermsOk returns a tuple with the DefaultPaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentTerms

`func (o *Client) SetDefaultPaymentTerms(v int32)`

SetDefaultPaymentTerms sets DefaultPaymentTerms field to given value.

### HasDefaultPaymentTerms

`func (o *Client) HasDefaultPaymentTerms() bool`

HasDefaultPaymentTerms returns a boolean if a field has been set.

### SetDefaultPaymentTermsNil

`func (o *Client) SetDefaultPaymentTermsNil(b bool)`

 SetDefaultPaymentTermsNil sets the value for DefaultPaymentTerms to be an explicit nil

### UnsetDefaultPaymentTerms
`func (o *Client) UnsetDefaultPaymentTerms()`

UnsetDefaultPaymentTerms ensures that no value is present for DefaultPaymentTerms, not even an explicit nil
### GetDefaultPaymentTermsType

`func (o *Client) GetDefaultPaymentTermsType() DefaultPaymentTermsType`

GetDefaultPaymentTermsType returns the DefaultPaymentTermsType field if non-nil, zero value otherwise.

### GetDefaultPaymentTermsTypeOk

`func (o *Client) GetDefaultPaymentTermsTypeOk() (*DefaultPaymentTermsType, bool)`

GetDefaultPaymentTermsTypeOk returns a tuple with the DefaultPaymentTermsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentTermsType

`func (o *Client) SetDefaultPaymentTermsType(v DefaultPaymentTermsType)`

SetDefaultPaymentTermsType sets DefaultPaymentTermsType field to given value.

### HasDefaultPaymentTermsType

`func (o *Client) HasDefaultPaymentTermsType() bool`

HasDefaultPaymentTermsType returns a boolean if a field has been set.

### GetDefaultPaymentMethod

`func (o *Client) GetDefaultPaymentMethod() PaymentMethod`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *Client) GetDefaultPaymentMethodOk() (*PaymentMethod, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *Client) SetDefaultPaymentMethod(v PaymentMethod)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.

### HasDefaultPaymentMethod

`func (o *Client) HasDefaultPaymentMethod() bool`

HasDefaultPaymentMethod returns a boolean if a field has been set.

### GetBankName

`func (o *Client) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Client) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Client) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *Client) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *Client) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *Client) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankIban

`func (o *Client) GetBankIban() string`

GetBankIban returns the BankIban field if non-nil, zero value otherwise.

### GetBankIbanOk

`func (o *Client) GetBankIbanOk() (*string, bool)`

GetBankIbanOk returns a tuple with the BankIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIban

`func (o *Client) SetBankIban(v string)`

SetBankIban sets BankIban field to given value.

### HasBankIban

`func (o *Client) HasBankIban() bool`

HasBankIban returns a boolean if a field has been set.

### SetBankIbanNil

`func (o *Client) SetBankIbanNil(b bool)`

 SetBankIbanNil sets the value for BankIban to be an explicit nil

### UnsetBankIban
`func (o *Client) UnsetBankIban()`

UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
### GetBankSwiftCode

`func (o *Client) GetBankSwiftCode() string`

GetBankSwiftCode returns the BankSwiftCode field if non-nil, zero value otherwise.

### GetBankSwiftCodeOk

`func (o *Client) GetBankSwiftCodeOk() (*string, bool)`

GetBankSwiftCodeOk returns a tuple with the BankSwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankSwiftCode

`func (o *Client) SetBankSwiftCode(v string)`

SetBankSwiftCode sets BankSwiftCode field to given value.

### HasBankSwiftCode

`func (o *Client) HasBankSwiftCode() bool`

HasBankSwiftCode returns a boolean if a field has been set.

### SetBankSwiftCodeNil

`func (o *Client) SetBankSwiftCodeNil(b bool)`

 SetBankSwiftCodeNil sets the value for BankSwiftCode to be an explicit nil

### UnsetBankSwiftCode
`func (o *Client) UnsetBankSwiftCode()`

UnsetBankSwiftCode ensures that no value is present for BankSwiftCode, not even an explicit nil
### GetShippingAddress

`func (o *Client) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Client) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Client) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Client) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### SetShippingAddressNil

`func (o *Client) SetShippingAddressNil(b bool)`

 SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil

### UnsetShippingAddress
`func (o *Client) UnsetShippingAddress()`

UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil
### GetEInvoice

`func (o *Client) GetEInvoice() bool`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *Client) GetEInvoiceOk() (*bool, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *Client) SetEInvoice(v bool)`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *Client) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### SetEInvoiceNil

`func (o *Client) SetEInvoiceNil(b bool)`

 SetEInvoiceNil sets the value for EInvoice to be an explicit nil

### UnsetEInvoice
`func (o *Client) UnsetEInvoice()`

UnsetEInvoice ensures that no value is present for EInvoice, not even an explicit nil
### GetEiCode

`func (o *Client) GetEiCode() string`

GetEiCode returns the EiCode field if non-nil, zero value otherwise.

### GetEiCodeOk

`func (o *Client) GetEiCodeOk() (*string, bool)`

GetEiCodeOk returns a tuple with the EiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEiCode

`func (o *Client) SetEiCode(v string)`

SetEiCode sets EiCode field to given value.

### HasEiCode

`func (o *Client) HasEiCode() bool`

HasEiCode returns a boolean if a field has been set.

### SetEiCodeNil

`func (o *Client) SetEiCodeNil(b bool)`

 SetEiCodeNil sets the value for EiCode to be an explicit nil

### UnsetEiCode
`func (o *Client) UnsetEiCode()`

UnsetEiCode ensures that no value is present for EiCode, not even an explicit nil
### GetDiscountHighlight

`func (o *Client) GetDiscountHighlight() bool`

GetDiscountHighlight returns the DiscountHighlight field if non-nil, zero value otherwise.

### GetDiscountHighlightOk

`func (o *Client) GetDiscountHighlightOk() (*bool, bool)`

GetDiscountHighlightOk returns a tuple with the DiscountHighlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountHighlight

`func (o *Client) SetDiscountHighlight(v bool)`

SetDiscountHighlight sets DiscountHighlight field to given value.

### HasDiscountHighlight

`func (o *Client) HasDiscountHighlight() bool`

HasDiscountHighlight returns a boolean if a field has been set.

### SetDiscountHighlightNil

`func (o *Client) SetDiscountHighlightNil(b bool)`

 SetDiscountHighlightNil sets the value for DiscountHighlight to be an explicit nil

### UnsetDiscountHighlight
`func (o *Client) UnsetDiscountHighlight()`

UnsetDiscountHighlight ensures that no value is present for DiscountHighlight, not even an explicit nil
### GetDefaultDiscount

`func (o *Client) GetDefaultDiscount() float32`

GetDefaultDiscount returns the DefaultDiscount field if non-nil, zero value otherwise.

### GetDefaultDiscountOk

`func (o *Client) GetDefaultDiscountOk() (*float32, bool)`

GetDefaultDiscountOk returns a tuple with the DefaultDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscount

`func (o *Client) SetDefaultDiscount(v float32)`

SetDefaultDiscount sets DefaultDiscount field to given value.

### HasDefaultDiscount

`func (o *Client) HasDefaultDiscount() bool`

HasDefaultDiscount returns a boolean if a field has been set.

### SetDefaultDiscountNil

`func (o *Client) SetDefaultDiscountNil(b bool)`

 SetDefaultDiscountNil sets the value for DefaultDiscount to be an explicit nil

### UnsetDefaultDiscount
`func (o *Client) UnsetDefaultDiscount()`

UnsetDefaultDiscount ensures that no value is present for DefaultDiscount, not even an explicit nil
### GetHasIntentDeclaration

`func (o *Client) GetHasIntentDeclaration() bool`

GetHasIntentDeclaration returns the HasIntentDeclaration field if non-nil, zero value otherwise.

### GetHasIntentDeclarationOk

`func (o *Client) GetHasIntentDeclarationOk() (*bool, bool)`

GetHasIntentDeclarationOk returns a tuple with the HasIntentDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIntentDeclaration

`func (o *Client) SetHasIntentDeclaration(v bool)`

SetHasIntentDeclaration sets HasIntentDeclaration field to given value.

### HasHasIntentDeclaration

`func (o *Client) HasHasIntentDeclaration() bool`

HasHasIntentDeclaration returns a boolean if a field has been set.

### SetHasIntentDeclarationNil

`func (o *Client) SetHasIntentDeclarationNil(b bool)`

 SetHasIntentDeclarationNil sets the value for HasIntentDeclaration to be an explicit nil

### UnsetHasIntentDeclaration
`func (o *Client) UnsetHasIntentDeclaration()`

UnsetHasIntentDeclaration ensures that no value is present for HasIntentDeclaration, not even an explicit nil
### GetIntentDeclarationProtocolNumber

`func (o *Client) GetIntentDeclarationProtocolNumber() string`

GetIntentDeclarationProtocolNumber returns the IntentDeclarationProtocolNumber field if non-nil, zero value otherwise.

### GetIntentDeclarationProtocolNumberOk

`func (o *Client) GetIntentDeclarationProtocolNumberOk() (*string, bool)`

GetIntentDeclarationProtocolNumberOk returns a tuple with the IntentDeclarationProtocolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentDeclarationProtocolNumber

`func (o *Client) SetIntentDeclarationProtocolNumber(v string)`

SetIntentDeclarationProtocolNumber sets IntentDeclarationProtocolNumber field to given value.

### HasIntentDeclarationProtocolNumber

`func (o *Client) HasIntentDeclarationProtocolNumber() bool`

HasIntentDeclarationProtocolNumber returns a boolean if a field has been set.

### SetIntentDeclarationProtocolNumberNil

`func (o *Client) SetIntentDeclarationProtocolNumberNil(b bool)`

 SetIntentDeclarationProtocolNumberNil sets the value for IntentDeclarationProtocolNumber to be an explicit nil

### UnsetIntentDeclarationProtocolNumber
`func (o *Client) UnsetIntentDeclarationProtocolNumber()`

UnsetIntentDeclarationProtocolNumber ensures that no value is present for IntentDeclarationProtocolNumber, not even an explicit nil
### GetIntentDeclarationProtocolDate

`func (o *Client) GetIntentDeclarationProtocolDate() string`

GetIntentDeclarationProtocolDate returns the IntentDeclarationProtocolDate field if non-nil, zero value otherwise.

### GetIntentDeclarationProtocolDateOk

`func (o *Client) GetIntentDeclarationProtocolDateOk() (*string, bool)`

GetIntentDeclarationProtocolDateOk returns a tuple with the IntentDeclarationProtocolDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentDeclarationProtocolDate

`func (o *Client) SetIntentDeclarationProtocolDate(v string)`

SetIntentDeclarationProtocolDate sets IntentDeclarationProtocolDate field to given value.

### HasIntentDeclarationProtocolDate

`func (o *Client) HasIntentDeclarationProtocolDate() bool`

HasIntentDeclarationProtocolDate returns a boolean if a field has been set.

### SetIntentDeclarationProtocolDateNil

`func (o *Client) SetIntentDeclarationProtocolDateNil(b bool)`

 SetIntentDeclarationProtocolDateNil sets the value for IntentDeclarationProtocolDate to be an explicit nil

### UnsetIntentDeclarationProtocolDate
`func (o *Client) UnsetIntentDeclarationProtocolDate()`

UnsetIntentDeclarationProtocolDate ensures that no value is present for IntentDeclarationProtocolDate, not even an explicit nil
### GetCreatedAt

`func (o *Client) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Client) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Client) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Client) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Client) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Client) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Client) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Client) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Client) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Client) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Client) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Client) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


