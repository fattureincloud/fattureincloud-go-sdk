/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// PermissionsFicIssuedDocumentsDetailed struct for PermissionsFicIssuedDocumentsDetailed
type PermissionsFicIssuedDocumentsDetailed struct {
	Quotes *PermissionLevel `json:"quotes,omitempty"`
	Proformas *PermissionLevel `json:"proformas,omitempty"`
	Invoices *PermissionLevel `json:"invoices,omitempty"`
	Receipts *PermissionLevel `json:"receipts,omitempty"`
	DeliveryNotes *PermissionLevel `json:"delivery_notes,omitempty"`
	CreditNotes *PermissionLevel `json:"credit_notes,omitempty"`
	Orders *PermissionLevel `json:"orders,omitempty"`
	WorkReports *PermissionLevel `json:"work_reports,omitempty"`
	SupplierOrders *PermissionLevel `json:"supplier_orders,omitempty"`
	SelfInvoices *PermissionLevel `json:"self_invoices,omitempty"`
}

// NewPermissionsFicIssuedDocumentsDetailed instantiates a new PermissionsFicIssuedDocumentsDetailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsFicIssuedDocumentsDetailed() *PermissionsFicIssuedDocumentsDetailed {
	this := PermissionsFicIssuedDocumentsDetailed{}
	return &this
}

// NewPermissionsFicIssuedDocumentsDetailedWithDefaults instantiates a new PermissionsFicIssuedDocumentsDetailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsFicIssuedDocumentsDetailedWithDefaults() *PermissionsFicIssuedDocumentsDetailed {
	this := PermissionsFicIssuedDocumentsDetailed{}
	return &this
}

// GetQuotes returns the Quotes field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetQuotes() PermissionLevel {
	if o == nil || o.Quotes == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.Quotes
}

// GetQuotesOk returns a tuple with the Quotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetQuotesOk() (*PermissionLevel, bool) {
	if o == nil || o.Quotes == nil {
		return nil, false
	}
	return o.Quotes, true
}

// HasQuotes returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasQuotes() bool {
	if o != nil && o.Quotes != nil {
		return true
	}

	return false
}

// SetQuotes gets a reference to the given PermissionLevel and assigns it to the Quotes field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetQuotes(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.Quotes = &v
	return o
}

// GetProformas returns the Proformas field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetProformas() PermissionLevel {
	if o == nil || o.Proformas == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.Proformas
}

// GetProformasOk returns a tuple with the Proformas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetProformasOk() (*PermissionLevel, bool) {
	if o == nil || o.Proformas == nil {
		return nil, false
	}
	return o.Proformas, true
}

// HasProformas returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasProformas() bool {
	if o != nil && o.Proformas != nil {
		return true
	}

	return false
}

// SetProformas gets a reference to the given PermissionLevel and assigns it to the Proformas field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetProformas(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.Proformas = &v
	return o
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetInvoices() PermissionLevel {
	if o == nil || o.Invoices == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetInvoicesOk() (*PermissionLevel, bool) {
	if o == nil || o.Invoices == nil {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasInvoices() bool {
	if o != nil && o.Invoices != nil {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given PermissionLevel and assigns it to the Invoices field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetInvoices(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.Invoices = &v
	return o
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetReceipts() PermissionLevel {
	if o == nil || o.Receipts == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetReceiptsOk() (*PermissionLevel, bool) {
	if o == nil || o.Receipts == nil {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasReceipts() bool {
	if o != nil && o.Receipts != nil {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given PermissionLevel and assigns it to the Receipts field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetReceipts(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.Receipts = &v
	return o
}

// GetDeliveryNotes returns the DeliveryNotes field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetDeliveryNotes() PermissionLevel {
	if o == nil || o.DeliveryNotes == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.DeliveryNotes
}

// GetDeliveryNotesOk returns a tuple with the DeliveryNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetDeliveryNotesOk() (*PermissionLevel, bool) {
	if o == nil || o.DeliveryNotes == nil {
		return nil, false
	}
	return o.DeliveryNotes, true
}

// HasDeliveryNotes returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasDeliveryNotes() bool {
	if o != nil && o.DeliveryNotes != nil {
		return true
	}

	return false
}

// SetDeliveryNotes gets a reference to the given PermissionLevel and assigns it to the DeliveryNotes field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetDeliveryNotes(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.DeliveryNotes = &v
	return o
}

// GetCreditNotes returns the CreditNotes field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetCreditNotes() PermissionLevel {
	if o == nil || o.CreditNotes == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.CreditNotes
}

// GetCreditNotesOk returns a tuple with the CreditNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetCreditNotesOk() (*PermissionLevel, bool) {
	if o == nil || o.CreditNotes == nil {
		return nil, false
	}
	return o.CreditNotes, true
}

// HasCreditNotes returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasCreditNotes() bool {
	if o != nil && o.CreditNotes != nil {
		return true
	}

	return false
}

// SetCreditNotes gets a reference to the given PermissionLevel and assigns it to the CreditNotes field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetCreditNotes(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.CreditNotes = &v
	return o
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetOrders() PermissionLevel {
	if o == nil || o.Orders == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetOrdersOk() (*PermissionLevel, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given PermissionLevel and assigns it to the Orders field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetOrders(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.Orders = &v
	return o
}

// GetWorkReports returns the WorkReports field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetWorkReports() PermissionLevel {
	if o == nil || o.WorkReports == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.WorkReports
}

// GetWorkReportsOk returns a tuple with the WorkReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetWorkReportsOk() (*PermissionLevel, bool) {
	if o == nil || o.WorkReports == nil {
		return nil, false
	}
	return o.WorkReports, true
}

// HasWorkReports returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasWorkReports() bool {
	if o != nil && o.WorkReports != nil {
		return true
	}

	return false
}

// SetWorkReports gets a reference to the given PermissionLevel and assigns it to the WorkReports field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetWorkReports(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.WorkReports = &v
	return o
}

// GetSupplierOrders returns the SupplierOrders field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetSupplierOrders() PermissionLevel {
	if o == nil || o.SupplierOrders == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.SupplierOrders
}

// GetSupplierOrdersOk returns a tuple with the SupplierOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetSupplierOrdersOk() (*PermissionLevel, bool) {
	if o == nil || o.SupplierOrders == nil {
		return nil, false
	}
	return o.SupplierOrders, true
}

// HasSupplierOrders returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasSupplierOrders() bool {
	if o != nil && o.SupplierOrders != nil {
		return true
	}

	return false
}

// SetSupplierOrders gets a reference to the given PermissionLevel and assigns it to the SupplierOrders field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetSupplierOrders(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.SupplierOrders = &v
	return o
}

// GetSelfInvoices returns the SelfInvoices field value if set, zero value otherwise.
func (o *PermissionsFicIssuedDocumentsDetailed) GetSelfInvoices() PermissionLevel {
	if o == nil || o.SelfInvoices == nil {
		var ret PermissionLevel
		return ret
	}
	return *o.SelfInvoices
}

// GetSelfInvoicesOk returns a tuple with the SelfInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) GetSelfInvoicesOk() (*PermissionLevel, bool) {
	if o == nil || o.SelfInvoices == nil {
		return nil, false
	}
	return o.SelfInvoices, true
}

// HasSelfInvoices returns a boolean if a field has been set.
func (o *PermissionsFicIssuedDocumentsDetailed) HasSelfInvoices() bool {
	if o != nil && o.SelfInvoices != nil {
		return true
	}

	return false
}

// SetSelfInvoices gets a reference to the given PermissionLevel and assigns it to the SelfInvoices field.
func (o *PermissionsFicIssuedDocumentsDetailed) SetSelfInvoices(v PermissionLevel) *PermissionsFicIssuedDocumentsDetailed {
	o.SelfInvoices = &v
	return o
}

func (o PermissionsFicIssuedDocumentsDetailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quotes != nil {
		toSerialize["quotes"] = o.Quotes
	}
	if o.Proformas != nil {
		toSerialize["proformas"] = o.Proformas
	}
	if o.Invoices != nil {
		toSerialize["invoices"] = o.Invoices
	}
	if o.Receipts != nil {
		toSerialize["receipts"] = o.Receipts
	}
	if o.DeliveryNotes != nil {
		toSerialize["delivery_notes"] = o.DeliveryNotes
	}
	if o.CreditNotes != nil {
		toSerialize["credit_notes"] = o.CreditNotes
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.WorkReports != nil {
		toSerialize["work_reports"] = o.WorkReports
	}
	if o.SupplierOrders != nil {
		toSerialize["supplier_orders"] = o.SupplierOrders
	}
	if o.SelfInvoices != nil {
		toSerialize["self_invoices"] = o.SelfInvoices
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionsFicIssuedDocumentsDetailed struct {
	value *PermissionsFicIssuedDocumentsDetailed
	isSet bool
}

func (v NullablePermissionsFicIssuedDocumentsDetailed) Get() *PermissionsFicIssuedDocumentsDetailed {
	return v.value
}

func (v *NullablePermissionsFicIssuedDocumentsDetailed) Set(val *PermissionsFicIssuedDocumentsDetailed) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsFicIssuedDocumentsDetailed) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsFicIssuedDocumentsDetailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsFicIssuedDocumentsDetailed(val *PermissionsFicIssuedDocumentsDetailed) *NullablePermissionsFicIssuedDocumentsDetailed {
	return &NullablePermissionsFicIssuedDocumentsDetailed{value: val, isSet: true}
}

func (v NullablePermissionsFicIssuedDocumentsDetailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsFicIssuedDocumentsDetailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


