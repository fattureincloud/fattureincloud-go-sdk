/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.1.0
Contact: info@fattureincloud.it
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
)

// EventType Webhooks event type
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
	ISSUED_DOCUMENTS_PROFORMAS_DELETE EventType
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
	ARCHIVE_DOCUMENTS_CREATE EventType
	ARCHIVE_DOCUMENTS_UPDATE EventType
	ARCHIVE_DOCUMENTS_DELETE EventType
	CASHBOOK_CREATE EventType
	CASHBOOK_UPDATE EventType
	CASHBOOK_DELETE EventType
	PRODUCTS_CREATE EventType
	PRODUCTS_UPDATE EventType
	PRODUCTS_DELETE EventType
	PRODUCTS_STOCK_UPDATE EventType
	ENTITIES_CLIENTS_CREATE EventType
	ENTITIES_CLIENTS_UPDATE EventType
	ENTITIES_CLIENTS_DELETE EventType
	ENTITIES_SUPPLIERS_CREATE EventType
	ENTITIES_SUPPLIERS_UPDATE EventType
	ENTITIES_SUPPLIERS_DELETE EventType
	ENTITIES_ALL_CREATE EventType
	ENTITIES_ALL_UPDATE EventType
	ENTITIES_ALL_DELETE EventType
	ISSUED_DOCUMENTS_E_INVOICES_STATUS_UPDATE EventType
	RECEIVED_DOCUMENTS_E_INVOICES_RECEIVE EventType
} {
	ISSUED_DOCUMENTS_INVOICES_CREATE: "it.fattureincloud.webhooks.issued_documents.invoices.create",
	ISSUED_DOCUMENTS_INVOICES_UPDATE: "it.fattureincloud.webhooks.issued_documents.invoices.update",
	ISSUED_DOCUMENTS_INVOICES_DELETE: "it.fattureincloud.webhooks.issued_documents.invoices.delete",
	ISSUED_DOCUMENTS_QUOTES_CREATE: "it.fattureincloud.webhooks.issued_documents.quotes.create",
	ISSUED_DOCUMENTS_QUOTES_UPDATE: "it.fattureincloud.webhooks.issued_documents.quotes.update",
	ISSUED_DOCUMENTS_QUOTES_DELETE: "it.fattureincloud.webhooks.issued_documents.quotes.delete",
	ISSUED_DOCUMENTS_PROFORMAS_CREATE: "it.fattureincloud.webhooks.issued_documents.proformas.create",
	ISSUED_DOCUMENTS_PROFORMAS_UPDATE: "it.fattureincloud.webhooks.issued_documents.proformas.update",
	ISSUED_DOCUMENTS_PROFORMAS_DELETE: "it.fattureincloud.webhooks.issued_documents.proformas.delete",
	ISSUED_DOCUMENTS_RECEIPTS_CREATE: "it.fattureincloud.webhooks.issued_documents.receipts.create",
	ISSUED_DOCUMENTS_RECEIPTS_UPDATE: "it.fattureincloud.webhooks.issued_documents.receipts.update",
	ISSUED_DOCUMENTS_RECEIPTS_DELETE: "it.fattureincloud.webhooks.issued_documents.receipts.delete",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_CREATE: "it.fattureincloud.webhooks.issued_documents.delivery_notes.create",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_UPDATE: "it.fattureincloud.webhooks.issued_documents.delivery_notes.update",
	ISSUED_DOCUMENTS_DELIVERY_NOTES_DELETE: "it.fattureincloud.webhooks.issued_documents.delivery_notes.delete",
	ISSUED_DOCUMENTS_CREDIT_NOTES_CREATE: "it.fattureincloud.webhooks.issued_documents.credit_notes.create",
	ISSUED_DOCUMENTS_CREDIT_NOTES_UPDATE: "it.fattureincloud.webhooks.issued_documents.credit_notes.update",
	ISSUED_DOCUMENTS_CREDIT_NOTES_DELETE: "it.fattureincloud.webhooks.issued_documents.credit_notes.delete",
	ISSUED_DOCUMENTS_ORDERS_CREATE: "it.fattureincloud.webhooks.issued_documents.orders.create",
	ISSUED_DOCUMENTS_ORDERS_UPDATE: "it.fattureincloud.webhooks.issued_documents.orders.update",
	ISSUED_DOCUMENTS_ORDERS_DELETE: "it.fattureincloud.webhooks.issued_documents.orders.delete",
	ISSUED_DOCUMENTS_WORK_REPORTS_CREATE: "it.fattureincloud.webhooks.issued_documents.work_reports.create",
	ISSUED_DOCUMENTS_WORK_REPORTS_UPDATE: "it.fattureincloud.webhooks.issued_documents.work_reports.update",
	ISSUED_DOCUMENTS_WORK_REPORTS_DELETE: "it.fattureincloud.webhooks.issued_documents.work_reports.delete",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_CREATE: "it.fattureincloud.webhooks.issued_documents.supplier_orders.create",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_UPDATE: "it.fattureincloud.webhooks.issued_documents.supplier_orders.update",
	ISSUED_DOCUMENTS_SUPPLIER_ORDERS_DELETE: "it.fattureincloud.webhooks.issued_documents.supplier_orders.delete",
	ISSUED_DOCUMENTS_SELF_INVOICES_CREATE: "it.fattureincloud.webhooks.issued_documents.self_invoices.create",
	ISSUED_DOCUMENTS_SELF_INVOICES_UPDATE: "it.fattureincloud.webhooks.issued_documents.self_invoices.update",
	ISSUED_DOCUMENTS_SELF_INVOICES_DELETE: "it.fattureincloud.webhooks.issued_documents.self_invoices.delete",
	ISSUED_DOCUMENTS_ALL_CREATE: "it.fattureincloud.webhooks.issued_documents.all.create",
	ISSUED_DOCUMENTS_ALL_UPDATE: "it.fattureincloud.webhooks.issued_documents.all.update",
	ISSUED_DOCUMENTS_ALL_DELETE: "it.fattureincloud.webhooks.issued_documents.all.delete",
	RECEIVED_DOCUMENTS_CREATE: "it.fattureincloud.webhooks.received_documents.create",
	RECEIVED_DOCUMENTS_UPDATE: "it.fattureincloud.webhooks.received_documents.update",
	RECEIVED_DOCUMENTS_DELETE: "it.fattureincloud.webhooks.received_documents.delete",
	RECEIPTS_CREATE: "it.fattureincloud.webhooks.receipts.create",
	RECEIPTS_UPDATE: "it.fattureincloud.webhooks.receipts.update",
	RECEIPTS_DELETE: "it.fattureincloud.webhooks.receipts.delete",
	TAXES_CREATE: "it.fattureincloud.webhooks.taxes.create",
	TAXES_UPDATE: "it.fattureincloud.webhooks.taxes.update",
	TAXES_DELETE: "it.fattureincloud.webhooks.taxes.delete",
	ARCHIVE_DOCUMENTS_CREATE: "it.fattureincloud.webhooks.archive_documents.create",
	ARCHIVE_DOCUMENTS_UPDATE: "it.fattureincloud.webhooks.archive_documents.update",
	ARCHIVE_DOCUMENTS_DELETE: "it.fattureincloud.webhooks.archive_documents.delete",
	CASHBOOK_CREATE: "it.fattureincloud.webhooks.cashbook.create",
	CASHBOOK_UPDATE: "it.fattureincloud.webhooks.cashbook.update",
	CASHBOOK_DELETE: "it.fattureincloud.webhooks.cashbook.delete",
	PRODUCTS_CREATE: "it.fattureincloud.webhooks.products.create",
	PRODUCTS_UPDATE: "it.fattureincloud.webhooks.products.update",
	PRODUCTS_DELETE: "it.fattureincloud.webhooks.products.delete",
	PRODUCTS_STOCK_UPDATE: "it.fattureincloud.webhooks.products.stock_update",
	ENTITIES_CLIENTS_CREATE: "it.fattureincloud.webhooks.entities.clients.create",
	ENTITIES_CLIENTS_UPDATE: "it.fattureincloud.webhooks.entities.clients.update",
	ENTITIES_CLIENTS_DELETE: "it.fattureincloud.webhooks.entities.clients.delete",
	ENTITIES_SUPPLIERS_CREATE: "it.fattureincloud.webhooks.entities.suppliers.create",
	ENTITIES_SUPPLIERS_UPDATE: "it.fattureincloud.webhooks.entities.suppliers.update",
	ENTITIES_SUPPLIERS_DELETE: "it.fattureincloud.webhooks.entities.suppliers.delete",
	ENTITIES_ALL_CREATE: "it.fattureincloud.webhooks.entities.all.create",
	ENTITIES_ALL_UPDATE: "it.fattureincloud.webhooks.entities.all.update",
	ENTITIES_ALL_DELETE: "it.fattureincloud.webhooks.entities.all.delete",
	ISSUED_DOCUMENTS_E_INVOICES_STATUS_UPDATE: "it.fattureincloud.webhooks.issued_documents.e_invoices.status_update",
	RECEIVED_DOCUMENTS_E_INVOICES_RECEIVE: "it.fattureincloud.webhooks.received_documents.e_invoices.receive",
}

// All allowed values of EventType enum
var AllowedEventTypeEnumValues = []EventType{
	"it.fattureincloud.webhooks.issued_documents.invoices.create",
	"it.fattureincloud.webhooks.issued_documents.invoices.update",
	"it.fattureincloud.webhooks.issued_documents.invoices.delete",
	"it.fattureincloud.webhooks.issued_documents.quotes.create",
	"it.fattureincloud.webhooks.issued_documents.quotes.update",
	"it.fattureincloud.webhooks.issued_documents.quotes.delete",
	"it.fattureincloud.webhooks.issued_documents.proformas.create",
	"it.fattureincloud.webhooks.issued_documents.proformas.update",
	"it.fattureincloud.webhooks.issued_documents.proformas.delete",
	"it.fattureincloud.webhooks.issued_documents.receipts.create",
	"it.fattureincloud.webhooks.issued_documents.receipts.update",
	"it.fattureincloud.webhooks.issued_documents.receipts.delete",
	"it.fattureincloud.webhooks.issued_documents.delivery_notes.create",
	"it.fattureincloud.webhooks.issued_documents.delivery_notes.update",
	"it.fattureincloud.webhooks.issued_documents.delivery_notes.delete",
	"it.fattureincloud.webhooks.issued_documents.credit_notes.create",
	"it.fattureincloud.webhooks.issued_documents.credit_notes.update",
	"it.fattureincloud.webhooks.issued_documents.credit_notes.delete",
	"it.fattureincloud.webhooks.issued_documents.orders.create",
	"it.fattureincloud.webhooks.issued_documents.orders.update",
	"it.fattureincloud.webhooks.issued_documents.orders.delete",
	"it.fattureincloud.webhooks.issued_documents.work_reports.create",
	"it.fattureincloud.webhooks.issued_documents.work_reports.update",
	"it.fattureincloud.webhooks.issued_documents.work_reports.delete",
	"it.fattureincloud.webhooks.issued_documents.supplier_orders.create",
	"it.fattureincloud.webhooks.issued_documents.supplier_orders.update",
	"it.fattureincloud.webhooks.issued_documents.supplier_orders.delete",
	"it.fattureincloud.webhooks.issued_documents.self_invoices.create",
	"it.fattureincloud.webhooks.issued_documents.self_invoices.update",
	"it.fattureincloud.webhooks.issued_documents.self_invoices.delete",
	"it.fattureincloud.webhooks.issued_documents.all.create",
	"it.fattureincloud.webhooks.issued_documents.all.update",
	"it.fattureincloud.webhooks.issued_documents.all.delete",
	"it.fattureincloud.webhooks.received_documents.create",
	"it.fattureincloud.webhooks.received_documents.update",
	"it.fattureincloud.webhooks.received_documents.delete",
	"it.fattureincloud.webhooks.receipts.create",
	"it.fattureincloud.webhooks.receipts.update",
	"it.fattureincloud.webhooks.receipts.delete",
	"it.fattureincloud.webhooks.taxes.create",
	"it.fattureincloud.webhooks.taxes.update",
	"it.fattureincloud.webhooks.taxes.delete",
	"it.fattureincloud.webhooks.archive_documents.create",
	"it.fattureincloud.webhooks.archive_documents.update",
	"it.fattureincloud.webhooks.archive_documents.delete",
	"it.fattureincloud.webhooks.cashbook.create",
	"it.fattureincloud.webhooks.cashbook.update",
	"it.fattureincloud.webhooks.cashbook.delete",
	"it.fattureincloud.webhooks.products.create",
	"it.fattureincloud.webhooks.products.update",
	"it.fattureincloud.webhooks.products.delete",
	"it.fattureincloud.webhooks.products.stock_update",
	"it.fattureincloud.webhooks.entities.clients.create",
	"it.fattureincloud.webhooks.entities.clients.update",
	"it.fattureincloud.webhooks.entities.clients.delete",
	"it.fattureincloud.webhooks.entities.suppliers.create",
	"it.fattureincloud.webhooks.entities.suppliers.update",
	"it.fattureincloud.webhooks.entities.suppliers.delete",
	"it.fattureincloud.webhooks.entities.all.create",
	"it.fattureincloud.webhooks.entities.all.update",
	"it.fattureincloud.webhooks.entities.all.delete",
	"it.fattureincloud.webhooks.issued_documents.e_invoices.status_update",
	"it.fattureincloud.webhooks.received_documents.e_invoices.receive",
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

