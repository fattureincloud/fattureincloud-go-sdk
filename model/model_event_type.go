/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.27
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// EventType Event type.
type EventType string

// List of EventType
var EventTypes = struct {
	ISSUED_DOCUMENTS_INVOICES_CREATE EventType
	ISSUED_DOCUMENTS_INVOICES_UPDATE EventType
	ISSUED_DOCUMENTS_INVOICES_DELETE EventType
	ISSUED_DOCUMENTS_QUOTES_CREATE EventType
	ISSUED_DOCUMENTS_QUOTES_UPDATE EventType
	ISSUED_DOCUMENTS_QUOTES_DELETE EventType
	ISSUED_DOCUMENTS_PROFORMAS_CREATE EventType
	ISSUED_DOCUMENTS_PROFORMAS_UPDATE EventType
	ISSUED_DOCUMENTS_PROFORMAS_CREATE EventType
	ISSUED_DOCUMENTS_RECEIPTS_CREATE EventType
	ISSUED_DOCUMENTS_RECEIPTS_UPDATE EventType
	ISSUED_DOCUMENTS_RECEIPTS_DELETE EventType
	ISSUED_DOCUMENTS_DELIVERY_NOTES_CREATE EventType
	ISSUED_DOCUMENTS_DELIVERY_NOTES_UPDATE EventType
	ISSUED_DOCUMENTS_DELIVERY_NOTES_DELETE EventType
	ISSUED_DOCUMENTS_CREDIT_NOTES_CREATE EventType
	ISSUED_DOCUMENTS_CREDIT_NOTES_UPDATE EventType
	ISSUED_DOCUMENTS_CREDIT_NOTES_DELETE EventType
	ISSUED_DOCUMENTS_ORDERS_CREATE EventType
	ISSUED_DOCUMENTS_ORDERS_UPDATE EventType
	ISSUED_DOCUMENTS_ORDERS_DELETE EventType
	ISSUED_DOCUMENTS_WORK_REPORTS_CREATE EventType
	ISSUED_DOCUMENTS_WORK_REPORTS_UPDATE EventType
	ISSUED_DOCUMENTS_WORK_REPORTS_DELETE EventType
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_CREATE EventType
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_UPDATE EventType
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_DELETE EventType
	ISSUED_DOCUMENTS_SELF_INVOICES_CREATE EventType
	ISSUED_DOCUMENTS_SELF_INVOICES_UPDATE EventType
	ISSUED_DOCUMENTS_SELF_INVOICES_DELETE EventType
	ISSUED_DOCUMENTS_ALL_CREATE EventType
	ISSUED_DOCUMENTS_ALL_UPDATE EventType
	ISSUED_DOCUMENTS_ALL_DELETE EventType
	RECEIVED_DOCUMENTS_CREATE EventType
	RECEIVED_DOCUMENTS_UPDATE EventType
	RECEIVED_DOCUMENTS_DELETE EventType
	RECEIPTS_CREATE EventType
	RECEIPTS_UPDATE EventType
	RECEIPTS_DELETE EventType
	TAXES_CREATE EventType
	TAXES_UPDATE EventType
	TAXES_DELETE EventType
	CASHBOOK_CREATE EventType
	CASHBOOK_UPDATE EventType
	CASHBOOK_DELETE EventType
	PRODUCTS_CREATE EventType
	PRODUCTS_UPDATE EventType
	PRODUCTS_DELETE EventType
	ENTITIES_CLIENTS_CREATE EventType
	ENTITIES_CLIENTS_UPDATE EventType
	ENTITIES_CLIENTS_DELETE EventType
	ENTITIES_SUPPLIERS_CREATE EventType
	ENTITIES_SUPPLIERS_UPDATE EventType
	ENTITIES_SUPPLIERS_DELETE EventType
	ENTITIES_ALL_CREATE EventType
	ENTITIES_ALL_UPDATE EventType
	ENTITIES_ALL_DELETE EventType
	ISSUED_DOCUMENTS_E_INVOICES EventType
	RECEIVED_DOCUMENTS_E_INVOICES EventType
} {
	ISSUED_DOCUMENTS_INVOICES_CREATE: "it.fattureincloud.issued_documents.invoices.create",
	ISSUED_DOCUMENTS_INVOICES_UPDATE: "it.fattureincloud.issued_documents.invoices.update",
	ISSUED_DOCUMENTS_INVOICES_DELETE: "it.fattureincloud.issued_documents.invoices.delete",
	ISSUED_DOCUMENTS_QUOTES_CREATE: "it.fattureincloud.issued_documents.quotes.create",
	ISSUED_DOCUMENTS_QUOTES_UPDATE: "it.fattureincloud.issued_documents.quotes.update",
	ISSUED_DOCUMENTS_QUOTES_DELETE: "it.fattureincloud.issued_documents.quotes.delete",
	ISSUED_DOCUMENTS_PROFORMAS_CREATE: "it.fattureincloud.issued_documents.proformas.create",
	ISSUED_DOCUMENTS_PROFORMAS_UPDATE: "it.fattureincloud.issued_documents.proformas.update",
	ISSUED_DOCUMENTS_PROFORMAS_CREATE: "it.fattureincloud.issued_documents.proformas.create",
	ISSUED_DOCUMENTS_RECEIPTS_CREATE: "it.fattureincloud.issued_documents.receipts.create",
	ISSUED_DOCUMENTS_RECEIPTS_UPDATE: "it.fattureincloud.issued_documents.receipts.update",
	ISSUED_DOCUMENTS_RECEIPTS_DELETE: "it.fattureincloud.issued_documents.receipts.delete",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_CREATE: "it.fattureincloud.issued_documents.delivery_notes.create",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_UPDATE: "it.fattureincloud.issued_documents.delivery_notes.update",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_DELETE: "it.fattureincloud.issued_documents.delivery_notes.delete",
	ISSUED_DOCUMENTS_CREDIT_NOTES_CREATE: "it.fattureincloud.issued_documents.credit_notes.create",
	ISSUED_DOCUMENTS_CREDIT_NOTES_UPDATE: "it.fattureincloud.issued_documents.credit_notes.update",
	ISSUED_DOCUMENTS_CREDIT_NOTES_DELETE: "it.fattureincloud.issued_documents.credit_notes.delete",
	ISSUED_DOCUMENTS_ORDERS_CREATE: "it.fattureincloud.issued_documents.orders.create",
	ISSUED_DOCUMENTS_ORDERS_UPDATE: "it.fattureincloud.issued_documents.orders.update",
	ISSUED_DOCUMENTS_ORDERS_DELETE: "it.fattureincloud.issued_documents.orders.delete",
	ISSUED_DOCUMENTS_WORK_REPORTS_CREATE: "it.fattureincloud.issued_documents.work_reports.create",
	ISSUED_DOCUMENTS_WORK_REPORTS_UPDATE: "it.fattureincloud.issued_documents.work_reports.update",
	ISSUED_DOCUMENTS_WORK_REPORTS_DELETE: "it.fattureincloud.issued_documents.work_reports.delete",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_CREATE: "it.fattureincloud.issued_documents.supplier_orders.create",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_UPDATE: "it.fattureincloud.issued_documents.supplier_orders.update",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_DELETE: "it.fattureincloud.issued_documents.supplier_orders.delete",
	ISSUED_DOCUMENTS_SELF_INVOICES_CREATE: "it.fattureincloud.issued_documents.self_invoices.create",
	ISSUED_DOCUMENTS_SELF_INVOICES_UPDATE: "it.fattureincloud.issued_documents.self_invoices.update",
	ISSUED_DOCUMENTS_SELF_INVOICES_DELETE: "it.fattureincloud.issued_documents.self_invoices.delete",
	ISSUED_DOCUMENTS_ALL_CREATE: "it.fattureincloud.issued_documents.all.create",
	ISSUED_DOCUMENTS_ALL_UPDATE: "it.fattureincloud.issued_documents.all.update",
	ISSUED_DOCUMENTS_ALL_DELETE: "it.fattureincloud.issued_documents.all.delete",
	RECEIVED_DOCUMENTS_CREATE: "it.fattureincloud.received_documents.create",
	RECEIVED_DOCUMENTS_UPDATE: "it.fattureincloud.received_documents.update",
	RECEIVED_DOCUMENTS_DELETE: "it.fattureincloud.received_documents.delete",
	RECEIPTS_CREATE: "it.fattureincloud.receipts.create",
	RECEIPTS_UPDATE: "it.fattureincloud.receipts.update",
	RECEIPTS_DELETE: "it.fattureincloud.receipts.delete",
	TAXES_CREATE: "it.fattureincloud.taxes.create",
	TAXES_UPDATE: "it.fattureincloud.taxes.update",
	TAXES_DELETE: "it.fattureincloud.taxes.delete",
	CASHBOOK_CREATE: "it.fattureincloud.cashbook.create",
	CASHBOOK_UPDATE: "it.fattureincloud.cashbook.update",
	CASHBOOK_DELETE: "it.fattureincloud.cashbook.delete",
	PRODUCTS_CREATE: "it.fattureincloud.products.create",
	PRODUCTS_UPDATE: "it.fattureincloud.products.update",
	PRODUCTS_DELETE: "it.fattureincloud.products.delete",
	ENTITIES_CLIENTS_CREATE: "it.fattureincloud.entities.clients.create",
	ENTITIES_CLIENTS_UPDATE: "it.fattureincloud.entities.clients.update",
	ENTITIES_CLIENTS_DELETE: "it.fattureincloud.entities.clients.delete",
	ENTITIES_SUPPLIERS_CREATE: "it.fattureincloud.entities.suppliers.create",
	ENTITIES_SUPPLIERS_UPDATE: "it.fattureincloud.entities.suppliers.update",
	ENTITIES_SUPPLIERS_DELETE: "it.fattureincloud.entities.suppliers.delete",
	ENTITIES_ALL_CREATE: "it.fattureincloud.entities.all.create",
	ENTITIES_ALL_UPDATE: "it.fattureincloud.entities.all.update",
	ENTITIES_ALL_DELETE: "it.fattureincloud.entities.all.delete",
	ISSUED_DOCUMENTS_E_INVOICES: "it.fattureincloud.issued_documents.e_invoices",
	RECEIVED_DOCUMENTS_E_INVOICES: "it.fattureincloud.received_documents.e_invoices",
}

// All allowed values of EventType enum
var AllowedEventTypeEnumValues = []EventType{
	"it.fattureincloud.issued_documents.invoices.create",
	"it.fattureincloud.issued_documents.invoices.update",
	"it.fattureincloud.issued_documents.invoices.delete",
	"it.fattureincloud.issued_documents.quotes.create",
	"it.fattureincloud.issued_documents.quotes.update",
	"it.fattureincloud.issued_documents.quotes.delete",
	"it.fattureincloud.issued_documents.proformas.create",
	"it.fattureincloud.issued_documents.proformas.update",
	"it.fattureincloud.issued_documents.proformas.create",
	"it.fattureincloud.issued_documents.receipts.create",
	"it.fattureincloud.issued_documents.receipts.update",
	"it.fattureincloud.issued_documents.receipts.delete",
	"it.fattureincloud.issued_documents.delivery_notes.create",
	"it.fattureincloud.issued_documents.delivery_notes.update",
	"it.fattureincloud.issued_documents.delivery_notes.delete",
	"it.fattureincloud.issued_documents.credit_notes.create",
	"it.fattureincloud.issued_documents.credit_notes.update",
	"it.fattureincloud.issued_documents.credit_notes.delete",
	"it.fattureincloud.issued_documents.orders.create",
	"it.fattureincloud.issued_documents.orders.update",
	"it.fattureincloud.issued_documents.orders.delete",
	"it.fattureincloud.issued_documents.work_reports.create",
	"it.fattureincloud.issued_documents.work_reports.update",
	"it.fattureincloud.issued_documents.work_reports.delete",
	"it.fattureincloud.issued_documents.supplier_orders.create",
	"it.fattureincloud.issued_documents.supplier_orders.update",
	"it.fattureincloud.issued_documents.supplier_orders.delete",
	"it.fattureincloud.issued_documents.self_invoices.create",
	"it.fattureincloud.issued_documents.self_invoices.update",
	"it.fattureincloud.issued_documents.self_invoices.delete",
	"it.fattureincloud.issued_documents.all.create",
	"it.fattureincloud.issued_documents.all.update",
	"it.fattureincloud.issued_documents.all.delete",
	"it.fattureincloud.received_documents.create",
	"it.fattureincloud.received_documents.update",
	"it.fattureincloud.received_documents.delete",
	"it.fattureincloud.receipts.create",
	"it.fattureincloud.receipts.update",
	"it.fattureincloud.receipts.delete",
	"it.fattureincloud.taxes.create",
	"it.fattureincloud.taxes.update",
	"it.fattureincloud.taxes.delete",
	"it.fattureincloud.cashbook.create",
	"it.fattureincloud.cashbook.update",
	"it.fattureincloud.cashbook.delete",
	"it.fattureincloud.products.create",
	"it.fattureincloud.products.update",
	"it.fattureincloud.products.delete",
	"it.fattureincloud.entities.clients.create",
	"it.fattureincloud.entities.clients.update",
	"it.fattureincloud.entities.clients.delete",
	"it.fattureincloud.entities.suppliers.create",
	"it.fattureincloud.entities.suppliers.update",
	"it.fattureincloud.entities.suppliers.delete",
	"it.fattureincloud.entities.all.create",
	"it.fattureincloud.entities.all.update",
	"it.fattureincloud.entities.all.delete",
	"it.fattureincloud.issued_documents.e_invoices",
	"it.fattureincloud.received_documents.e_invoices",
}

func (v *EventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventType(value)
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventType", value)
}

// NewEventTypeFromValue returns a pointer to a valid EventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventType: valid values are %v", v, AllowedEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventType) IsValid() bool {
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventType value
func (v EventType) Ptr() *EventType {
	return &v
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

