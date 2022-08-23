# IssuedDocumentEiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatKind** | Pointer to [**NullableVatKind**](VatKind.md) |  | [optional] 
**OriginalDocumentType** | Pointer to [**NullableOriginalDocumentType**](OriginalDocumentType.md) |  | [optional] [default to ORDINE]
**OdNumber** | Pointer to **NullableString** | Original document number. | [optional] 
**OdDate** | Pointer to **NullableString** | Original document date. | [optional] 
**Cig** | Pointer to **NullableString** | CIG. | [optional] 
**Cup** | Pointer to **NullableString** | CUP. | [optional] 
**PaymentMethod** | Pointer to **NullableString** | Payment method (see https://www.fatturapa.gov.it/export/documenti/fatturapa/v1.2.1/Rappresentazione-tabellare-fattura-ordinaria.pdf for the accepted values of ModalitaPagamento). | [optional] 
**BankName** | Pointer to **NullableString** | Bank name. | [optional] 
**BankIban** | Pointer to **NullableString** | IBAN. | [optional] 
**BankBeneficiary** | Pointer to **NullableString** | Bank beneficiary. | [optional] 
**InvoiceNumber** | Pointer to **NullableString** | Invoice number. | [optional] 
**InvoiceDate** | Pointer to **NullableString** | Invoice date. | [optional] 

## Methods

### NewIssuedDocumentEiData

`func NewIssuedDocumentEiData() *IssuedDocumentEiData`

NewIssuedDocumentEiData instantiates a new IssuedDocumentEiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentEiDataWithDefaults

`func NewIssuedDocumentEiDataWithDefaults() *IssuedDocumentEiData`

NewIssuedDocumentEiDataWithDefaults instantiates a new IssuedDocumentEiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatKind

`func (o *IssuedDocumentEiData) GetVatKind() VatKind`

GetVatKind returns the VatKind field if non-nil, zero value otherwise.

### GetVatKindOk

`func (o *IssuedDocumentEiData) GetVatKindOk() (*VatKind, bool)`

GetVatKindOk returns a tuple with the VatKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatKind

`func (o *IssuedDocumentEiData) SetVatKind(v VatKind)`

SetVatKind sets VatKind field to given value.

### HasVatKind

`func (o *IssuedDocumentEiData) HasVatKind() bool`

HasVatKind returns a boolean if a field has been set.

### SetVatKindNil

`func (o *IssuedDocumentEiData) SetVatKindNil(b bool)`

 SetVatKindNil sets the value for VatKind to be an explicit nil

### UnsetVatKind
`func (o *IssuedDocumentEiData) UnsetVatKind()`

UnsetVatKind ensures that no value is present for VatKind, not even an explicit nil
### GetOriginalDocumentType

`func (o *IssuedDocumentEiData) GetOriginalDocumentType() OriginalDocumentType`

GetOriginalDocumentType returns the OriginalDocumentType field if non-nil, zero value otherwise.

### GetOriginalDocumentTypeOk

`func (o *IssuedDocumentEiData) GetOriginalDocumentTypeOk() (*OriginalDocumentType, bool)`

GetOriginalDocumentTypeOk returns a tuple with the OriginalDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDocumentType

`func (o *IssuedDocumentEiData) SetOriginalDocumentType(v OriginalDocumentType)`

SetOriginalDocumentType sets OriginalDocumentType field to given value.

### HasOriginalDocumentType

`func (o *IssuedDocumentEiData) HasOriginalDocumentType() bool`

HasOriginalDocumentType returns a boolean if a field has been set.

### SetOriginalDocumentTypeNil

`func (o *IssuedDocumentEiData) SetOriginalDocumentTypeNil(b bool)`

 SetOriginalDocumentTypeNil sets the value for OriginalDocumentType to be an explicit nil

### UnsetOriginalDocumentType
`func (o *IssuedDocumentEiData) UnsetOriginalDocumentType()`

UnsetOriginalDocumentType ensures that no value is present for OriginalDocumentType, not even an explicit nil
### GetOdNumber

`func (o *IssuedDocumentEiData) GetOdNumber() string`

GetOdNumber returns the OdNumber field if non-nil, zero value otherwise.

### GetOdNumberOk

`func (o *IssuedDocumentEiData) GetOdNumberOk() (*string, bool)`

GetOdNumberOk returns a tuple with the OdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdNumber

`func (o *IssuedDocumentEiData) SetOdNumber(v string)`

SetOdNumber sets OdNumber field to given value.

### HasOdNumber

`func (o *IssuedDocumentEiData) HasOdNumber() bool`

HasOdNumber returns a boolean if a field has been set.

### SetOdNumberNil

`func (o *IssuedDocumentEiData) SetOdNumberNil(b bool)`

 SetOdNumberNil sets the value for OdNumber to be an explicit nil

### UnsetOdNumber
`func (o *IssuedDocumentEiData) UnsetOdNumber()`

UnsetOdNumber ensures that no value is present for OdNumber, not even an explicit nil
### GetOdDate

`func (o *IssuedDocumentEiData) GetOdDate() string`

GetOdDate returns the OdDate field if non-nil, zero value otherwise.

### GetOdDateOk

`func (o *IssuedDocumentEiData) GetOdDateOk() (*string, bool)`

GetOdDateOk returns a tuple with the OdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdDate

`func (o *IssuedDocumentEiData) SetOdDate(v string)`

SetOdDate sets OdDate field to given value.

### HasOdDate

`func (o *IssuedDocumentEiData) HasOdDate() bool`

HasOdDate returns a boolean if a field has been set.

### SetOdDateNil

`func (o *IssuedDocumentEiData) SetOdDateNil(b bool)`

 SetOdDateNil sets the value for OdDate to be an explicit nil

### UnsetOdDate
`func (o *IssuedDocumentEiData) UnsetOdDate()`

UnsetOdDate ensures that no value is present for OdDate, not even an explicit nil
### GetCig

`func (o *IssuedDocumentEiData) GetCig() string`

GetCig returns the Cig field if non-nil, zero value otherwise.

### GetCigOk

`func (o *IssuedDocumentEiData) GetCigOk() (*string, bool)`

GetCigOk returns a tuple with the Cig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCig

`func (o *IssuedDocumentEiData) SetCig(v string)`

SetCig sets Cig field to given value.

### HasCig

`func (o *IssuedDocumentEiData) HasCig() bool`

HasCig returns a boolean if a field has been set.

### SetCigNil

`func (o *IssuedDocumentEiData) SetCigNil(b bool)`

 SetCigNil sets the value for Cig to be an explicit nil

### UnsetCig
`func (o *IssuedDocumentEiData) UnsetCig()`

UnsetCig ensures that no value is present for Cig, not even an explicit nil
### GetCup

`func (o *IssuedDocumentEiData) GetCup() string`

GetCup returns the Cup field if non-nil, zero value otherwise.

### GetCupOk

`func (o *IssuedDocumentEiData) GetCupOk() (*string, bool)`

GetCupOk returns a tuple with the Cup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCup

`func (o *IssuedDocumentEiData) SetCup(v string)`

SetCup sets Cup field to given value.

### HasCup

`func (o *IssuedDocumentEiData) HasCup() bool`

HasCup returns a boolean if a field has been set.

### SetCupNil

`func (o *IssuedDocumentEiData) SetCupNil(b bool)`

 SetCupNil sets the value for Cup to be an explicit nil

### UnsetCup
`func (o *IssuedDocumentEiData) UnsetCup()`

UnsetCup ensures that no value is present for Cup, not even an explicit nil
### GetPaymentMethod

`func (o *IssuedDocumentEiData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *IssuedDocumentEiData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *IssuedDocumentEiData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *IssuedDocumentEiData) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *IssuedDocumentEiData) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *IssuedDocumentEiData) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetBankName

`func (o *IssuedDocumentEiData) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *IssuedDocumentEiData) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *IssuedDocumentEiData) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *IssuedDocumentEiData) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *IssuedDocumentEiData) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *IssuedDocumentEiData) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankIban

`func (o *IssuedDocumentEiData) GetBankIban() string`

GetBankIban returns the BankIban field if non-nil, zero value otherwise.

### GetBankIbanOk

`func (o *IssuedDocumentEiData) GetBankIbanOk() (*string, bool)`

GetBankIbanOk returns a tuple with the BankIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIban

`func (o *IssuedDocumentEiData) SetBankIban(v string)`

SetBankIban sets BankIban field to given value.

### HasBankIban

`func (o *IssuedDocumentEiData) HasBankIban() bool`

HasBankIban returns a boolean if a field has been set.

### SetBankIbanNil

`func (o *IssuedDocumentEiData) SetBankIbanNil(b bool)`

 SetBankIbanNil sets the value for BankIban to be an explicit nil

### UnsetBankIban
`func (o *IssuedDocumentEiData) UnsetBankIban()`

UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
### GetBankBeneficiary

`func (o *IssuedDocumentEiData) GetBankBeneficiary() string`

GetBankBeneficiary returns the BankBeneficiary field if non-nil, zero value otherwise.

### GetBankBeneficiaryOk

`func (o *IssuedDocumentEiData) GetBankBeneficiaryOk() (*string, bool)`

GetBankBeneficiaryOk returns a tuple with the BankBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBeneficiary

`func (o *IssuedDocumentEiData) SetBankBeneficiary(v string)`

SetBankBeneficiary sets BankBeneficiary field to given value.

### HasBankBeneficiary

`func (o *IssuedDocumentEiData) HasBankBeneficiary() bool`

HasBankBeneficiary returns a boolean if a field has been set.

### SetBankBeneficiaryNil

`func (o *IssuedDocumentEiData) SetBankBeneficiaryNil(b bool)`

 SetBankBeneficiaryNil sets the value for BankBeneficiary to be an explicit nil

### UnsetBankBeneficiary
`func (o *IssuedDocumentEiData) UnsetBankBeneficiary()`

UnsetBankBeneficiary ensures that no value is present for BankBeneficiary, not even an explicit nil
### GetInvoiceNumber

`func (o *IssuedDocumentEiData) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *IssuedDocumentEiData) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *IssuedDocumentEiData) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *IssuedDocumentEiData) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### SetInvoiceNumberNil

`func (o *IssuedDocumentEiData) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *IssuedDocumentEiData) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetInvoiceDate

`func (o *IssuedDocumentEiData) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *IssuedDocumentEiData) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *IssuedDocumentEiData) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *IssuedDocumentEiData) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### SetInvoiceDateNil

`func (o *IssuedDocumentEiData) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *IssuedDocumentEiData) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


