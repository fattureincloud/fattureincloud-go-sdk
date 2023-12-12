/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.31
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// checks if the IssuedDocumentEiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuedDocumentEiData{}

// IssuedDocumentEiData Issued document e-invoice data [Only if e_invoice=true]
type IssuedDocumentEiData struct {
	VatKind NullableVatKind `json:"vat_kind,omitempty"`
	OriginalDocumentType NullableOriginalDocumentType `json:"original_document_type,omitempty"`
	// E-invoice original document number
	OdNumber NullableString `json:"od_number,omitempty"`
	// E-invoice original document date
	OdDate NullableString `json:"od_date,omitempty"`
	// E-invoice CIG
	Cig NullableString `json:"cig,omitempty"`
	// E-invoice CUP
	Cup NullableString `json:"cup,omitempty"`
	// E-invoice payment method [required for e-invoices] (see [here](https://www.fatturapa.gov.it/export/documenti/fatturapa/v1.2.2/Rappresentazione_Tabellare_FattOrdinaria_V1.2.2.pdf) for the accepted values of ModalitaPagamento)
	PaymentMethod NullableString `json:"payment_method,omitempty"`
	// E-invoice bank name
	BankName NullableString `json:"bank_name,omitempty"`
	// E-invoice bank IBAN
	BankIban NullableString `json:"bank_iban,omitempty"`
	// E-invoice bank beneficiary
	BankBeneficiary NullableString `json:"bank_beneficiary,omitempty"`
	// E-invoice invoice number
	InvoiceNumber NullableString `json:"invoice_number,omitempty"`
	// E-invoice invoice date
	InvoiceDate NullableString `json:"invoice_date,omitempty"`
}

// NewIssuedDocumentEiData instantiates a new IssuedDocumentEiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocumentEiData() *IssuedDocumentEiData {
	this := IssuedDocumentEiData{}
	return &this
}

// NewIssuedDocumentEiDataWithDefaults instantiates a new IssuedDocumentEiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentEiDataWithDefaults() *IssuedDocumentEiData {
	this := IssuedDocumentEiData{}
	var originalDocumentType OriginalDocumentType = OriginalDocumentTypes.ORDINE
	this.OriginalDocumentType = *NewNullableOriginalDocumentType(&originalDocumentType)
	return &this
}

// GetVatKind returns the VatKind field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetVatKind() VatKind {
	if o == nil || IsNil(o.VatKind.Get()) {
		var ret VatKind
		return ret
	}
	return *o.VatKind.Get()
}

// GetVatKindOk returns a tuple with the VatKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetVatKindOk() (*VatKind, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatKind.Get(), o.VatKind.IsSet()
}

// HasVatKind returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasVatKind() bool {
	if o != nil && o.VatKind.IsSet() {
		return true
	}

	return false
}

// SetVatKind gets a reference to the given NullableVatKind and assigns it to the VatKind field.
func (o *IssuedDocumentEiData) SetVatKind(v VatKind) *IssuedDocumentEiData {
	o.VatKind.Set(&v)
	return o
}
// SetVatKindNil sets the value for VatKind to be an explicit nil
func (o *IssuedDocumentEiData) SetVatKindNil() *IssuedDocumentEiData {
	o.VatKind.Set(nil)
	return o
}

// UnsetVatKind ensures that no value is present for VatKind, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetVatKind() {
	o.VatKind.Unset()
}

// GetOriginalDocumentType returns the OriginalDocumentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetOriginalDocumentType() OriginalDocumentType {
	if o == nil || IsNil(o.OriginalDocumentType.Get()) {
		var ret OriginalDocumentType
		return ret
	}
	return *o.OriginalDocumentType.Get()
}

// GetOriginalDocumentTypeOk returns a tuple with the OriginalDocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetOriginalDocumentTypeOk() (*OriginalDocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalDocumentType.Get(), o.OriginalDocumentType.IsSet()
}

// HasOriginalDocumentType returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasOriginalDocumentType() bool {
	if o != nil && o.OriginalDocumentType.IsSet() {
		return true
	}

	return false
}

// SetOriginalDocumentType gets a reference to the given NullableOriginalDocumentType and assigns it to the OriginalDocumentType field.
func (o *IssuedDocumentEiData) SetOriginalDocumentType(v OriginalDocumentType) *IssuedDocumentEiData {
	o.OriginalDocumentType.Set(&v)
	return o
}
// SetOriginalDocumentTypeNil sets the value for OriginalDocumentType to be an explicit nil
func (o *IssuedDocumentEiData) SetOriginalDocumentTypeNil() *IssuedDocumentEiData {
	o.OriginalDocumentType.Set(nil)
	return o
}

// UnsetOriginalDocumentType ensures that no value is present for OriginalDocumentType, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetOriginalDocumentType() {
	o.OriginalDocumentType.Unset()
}

// GetOdNumber returns the OdNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetOdNumber() string {
	if o == nil || IsNil(o.OdNumber.Get()) {
		var ret string
		return ret
	}
	return *o.OdNumber.Get()
}

// GetOdNumberOk returns a tuple with the OdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetOdNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OdNumber.Get(), o.OdNumber.IsSet()
}

// HasOdNumber returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasOdNumber() bool {
	if o != nil && o.OdNumber.IsSet() {
		return true
	}

	return false
}

// SetOdNumber gets a reference to the given NullableString and assigns it to the OdNumber field.
func (o *IssuedDocumentEiData) SetOdNumber(v string) *IssuedDocumentEiData {
	o.OdNumber.Set(&v)
	return o
}
// SetOdNumberNil sets the value for OdNumber to be an explicit nil
func (o *IssuedDocumentEiData) SetOdNumberNil() *IssuedDocumentEiData {
	o.OdNumber.Set(nil)
	return o
}

// UnsetOdNumber ensures that no value is present for OdNumber, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetOdNumber() {
	o.OdNumber.Unset()
}

// GetOdDate returns the OdDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetOdDate() string {
	if o == nil || IsNil(o.OdDate.Get()) {
		var ret string
		return ret
	}
	return *o.OdDate.Get()
}

// GetOdDateOk returns a tuple with the OdDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetOdDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OdDate.Get(), o.OdDate.IsSet()
}

// HasOdDate returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasOdDate() bool {
	if o != nil && o.OdDate.IsSet() {
		return true
	}

	return false
}

// SetOdDate gets a reference to the given NullableString and assigns it to the OdDate field.
func (o *IssuedDocumentEiData) SetOdDate(v string) *IssuedDocumentEiData {
	o.OdDate.Set(&v)
	return o
}
// SetOdDateNil sets the value for OdDate to be an explicit nil
func (o *IssuedDocumentEiData) SetOdDateNil() *IssuedDocumentEiData {
	o.OdDate.Set(nil)
	return o
}

// UnsetOdDate ensures that no value is present for OdDate, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetOdDate() {
	o.OdDate.Unset()
}

// GetCig returns the Cig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetCig() string {
	if o == nil || IsNil(o.Cig.Get()) {
		var ret string
		return ret
	}
	return *o.Cig.Get()
}

// GetCigOk returns a tuple with the Cig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetCigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cig.Get(), o.Cig.IsSet()
}

// HasCig returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasCig() bool {
	if o != nil && o.Cig.IsSet() {
		return true
	}

	return false
}

// SetCig gets a reference to the given NullableString and assigns it to the Cig field.
func (o *IssuedDocumentEiData) SetCig(v string) *IssuedDocumentEiData {
	o.Cig.Set(&v)
	return o
}
// SetCigNil sets the value for Cig to be an explicit nil
func (o *IssuedDocumentEiData) SetCigNil() *IssuedDocumentEiData {
	o.Cig.Set(nil)
	return o
}

// UnsetCig ensures that no value is present for Cig, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetCig() {
	o.Cig.Unset()
}

// GetCup returns the Cup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetCup() string {
	if o == nil || IsNil(o.Cup.Get()) {
		var ret string
		return ret
	}
	return *o.Cup.Get()
}

// GetCupOk returns a tuple with the Cup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetCupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cup.Get(), o.Cup.IsSet()
}

// HasCup returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasCup() bool {
	if o != nil && o.Cup.IsSet() {
		return true
	}

	return false
}

// SetCup gets a reference to the given NullableString and assigns it to the Cup field.
func (o *IssuedDocumentEiData) SetCup(v string) *IssuedDocumentEiData {
	o.Cup.Set(&v)
	return o
}
// SetCupNil sets the value for Cup to be an explicit nil
func (o *IssuedDocumentEiData) SetCupNil() *IssuedDocumentEiData {
	o.Cup.Set(nil)
	return o
}

// UnsetCup ensures that no value is present for Cup, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetCup() {
	o.Cup.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given NullableString and assigns it to the PaymentMethod field.
func (o *IssuedDocumentEiData) SetPaymentMethod(v string) *IssuedDocumentEiData {
	o.PaymentMethod.Set(&v)
	return o
}
// SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil
func (o *IssuedDocumentEiData) SetPaymentMethodNil() *IssuedDocumentEiData {
	o.PaymentMethod.Set(nil)
	return o
}

// UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetPaymentMethod() {
	o.PaymentMethod.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetBankName() string {
	if o == nil || IsNil(o.BankName.Get()) {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *IssuedDocumentEiData) SetBankName(v string) *IssuedDocumentEiData {
	o.BankName.Set(&v)
	return o
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *IssuedDocumentEiData) SetBankNameNil() *IssuedDocumentEiData {
	o.BankName.Set(nil)
	return o
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetBankName() {
	o.BankName.Unset()
}

// GetBankIban returns the BankIban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetBankIban() string {
	if o == nil || IsNil(o.BankIban.Get()) {
		var ret string
		return ret
	}
	return *o.BankIban.Get()
}

// GetBankIbanOk returns a tuple with the BankIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetBankIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankIban.Get(), o.BankIban.IsSet()
}

// HasBankIban returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasBankIban() bool {
	if o != nil && o.BankIban.IsSet() {
		return true
	}

	return false
}

// SetBankIban gets a reference to the given NullableString and assigns it to the BankIban field.
func (o *IssuedDocumentEiData) SetBankIban(v string) *IssuedDocumentEiData {
	o.BankIban.Set(&v)
	return o
}
// SetBankIbanNil sets the value for BankIban to be an explicit nil
func (o *IssuedDocumentEiData) SetBankIbanNil() *IssuedDocumentEiData {
	o.BankIban.Set(nil)
	return o
}

// UnsetBankIban ensures that no value is present for BankIban, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetBankIban() {
	o.BankIban.Unset()
}

// GetBankBeneficiary returns the BankBeneficiary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetBankBeneficiary() string {
	if o == nil || IsNil(o.BankBeneficiary.Get()) {
		var ret string
		return ret
	}
	return *o.BankBeneficiary.Get()
}

// GetBankBeneficiaryOk returns a tuple with the BankBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetBankBeneficiaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankBeneficiary.Get(), o.BankBeneficiary.IsSet()
}

// HasBankBeneficiary returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasBankBeneficiary() bool {
	if o != nil && o.BankBeneficiary.IsSet() {
		return true
	}

	return false
}

// SetBankBeneficiary gets a reference to the given NullableString and assigns it to the BankBeneficiary field.
func (o *IssuedDocumentEiData) SetBankBeneficiary(v string) *IssuedDocumentEiData {
	o.BankBeneficiary.Set(&v)
	return o
}
// SetBankBeneficiaryNil sets the value for BankBeneficiary to be an explicit nil
func (o *IssuedDocumentEiData) SetBankBeneficiaryNil() *IssuedDocumentEiData {
	o.BankBeneficiary.Set(nil)
	return o
}

// UnsetBankBeneficiary ensures that no value is present for BankBeneficiary, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetBankBeneficiary() {
	o.BankBeneficiary.Unset()
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber.IsSet() {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given NullableString and assigns it to the InvoiceNumber field.
func (o *IssuedDocumentEiData) SetInvoiceNumber(v string) *IssuedDocumentEiData {
	o.InvoiceNumber.Set(&v)
	return o
}
// SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil
func (o *IssuedDocumentEiData) SetInvoiceNumberNil() *IssuedDocumentEiData {
	o.InvoiceNumber.Set(nil)
	return o
}

// UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetInvoiceNumber() {
	o.InvoiceNumber.Unset()
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocumentEiData) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceDate.Get()
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocumentEiData) GetInvoiceDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceDate.Get(), o.InvoiceDate.IsSet()
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *IssuedDocumentEiData) HasInvoiceDate() bool {
	if o != nil && o.InvoiceDate.IsSet() {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given NullableString and assigns it to the InvoiceDate field.
func (o *IssuedDocumentEiData) SetInvoiceDate(v string) *IssuedDocumentEiData {
	o.InvoiceDate.Set(&v)
	return o
}
// SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil
func (o *IssuedDocumentEiData) SetInvoiceDateNil() *IssuedDocumentEiData {
	o.InvoiceDate.Set(nil)
	return o
}

// UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil
func (o *IssuedDocumentEiData) UnsetInvoiceDate() {
	o.InvoiceDate.Unset()
}

func (o IssuedDocumentEiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuedDocumentEiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VatKind.IsSet() {
		toSerialize["vat_kind"] = o.VatKind.Get()
	}
	if o.OriginalDocumentType.IsSet() {
		toSerialize["original_document_type"] = o.OriginalDocumentType.Get()
	}
	if o.OdNumber.IsSet() {
		toSerialize["od_number"] = o.OdNumber.Get()
	}
	if o.OdDate.IsSet() {
		toSerialize["od_date"] = o.OdDate.Get()
	}
	if o.Cig.IsSet() {
		toSerialize["cig"] = o.Cig.Get()
	}
	if o.Cup.IsSet() {
		toSerialize["cup"] = o.Cup.Get()
	}
	if o.PaymentMethod.IsSet() {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if o.BankIban.IsSet() {
		toSerialize["bank_iban"] = o.BankIban.Get()
	}
	if o.BankBeneficiary.IsSet() {
		toSerialize["bank_beneficiary"] = o.BankBeneficiary.Get()
	}
	if o.InvoiceNumber.IsSet() {
		toSerialize["invoice_number"] = o.InvoiceNumber.Get()
	}
	if o.InvoiceDate.IsSet() {
		toSerialize["invoice_date"] = o.InvoiceDate.Get()
	}
	return toSerialize, nil
}

type NullableIssuedDocumentEiData struct {
	value *IssuedDocumentEiData
	isSet bool
}

func (v NullableIssuedDocumentEiData) Get() *IssuedDocumentEiData {
	return v.value
}

func (v *NullableIssuedDocumentEiData) Set(val *IssuedDocumentEiData) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocumentEiData) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocumentEiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocumentEiData(val *IssuedDocumentEiData) *NullableIssuedDocumentEiData {
	return &NullableIssuedDocumentEiData{value: val, isSet: true}
}

func (v NullableIssuedDocumentEiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocumentEiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


