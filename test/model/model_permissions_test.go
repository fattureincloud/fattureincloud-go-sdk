/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

var PermissionsJsonStr string = "{\"fic_situation\":\"write\",\"fic_clients\":\"write\",\"fic_suppliers\":\"write\",\"fic_products\":\"write\",\"fic_issued_documents\":\"write\",\"fic_received_documents\":\"write\",\"fic_receipts\":\"write\",\"fic_calendar\":\"write\",\"fic_archive\":\"write\",\"fic_taxes\":\"write\",\"fic_stock\":\"write\",\"fic_cashbook\":\"write\",\"fic_settings\":\"write\",\"fic_emails\":\"write\",\"fic_export\":\"write\",\"fic_import_bankstatements\":\"write\",\"fic_import_clients_suppliers\":\"write\",\"fic_import_issued_documents\":\"write\",\"fic_import_products\":\"write\",\"fic_recurring\":\"write\",\"fic_riba\":\"write\",\"dic_employees\":\"write\",\"dic_settings\":\"write\",\"dic_timesheet\":\"write\",\"fic_issued_documents_detailed\":{\"quotes\":\"write\",\"proformas\":\"write\",\"invoices\":\"write\",\"receipts\":\"write\",\"delivery_notes\":\"write\",\"credit_notes\":\"write\",\"orders\":\"write\",\"work_reports\":\"write\",\"supplier_orders\":\"write\",\"self_invoices\":\"write\"}}"

func TestPermissions(t *testing.T) {
	obj := NewPermissions()
	json.Unmarshal([]byte(PermissionsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PermissionsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewPermissions()
	newObj.SetFicSituation(obj.GetFicSituation())
	assert.True(t, reflect.DeepEqual(obj.GetFicSituation(), newObj.GetFicSituation()))
	newObj.SetFicClients(obj.GetFicClients())
	assert.True(t, reflect.DeepEqual(obj.GetFicClients(), newObj.GetFicClients()))
	newObj.SetFicSuppliers(obj.GetFicSuppliers())
	assert.True(t, reflect.DeepEqual(obj.GetFicSuppliers(), newObj.GetFicSuppliers()))
	newObj.SetFicProducts(obj.GetFicProducts())
	assert.True(t, reflect.DeepEqual(obj.GetFicProducts(), newObj.GetFicProducts()))
	newObj.SetFicIssuedDocuments(obj.GetFicIssuedDocuments())
	assert.True(t, reflect.DeepEqual(obj.GetFicIssuedDocuments(), newObj.GetFicIssuedDocuments()))
	newObj.SetFicReceivedDocuments(obj.GetFicReceivedDocuments())
	assert.True(t, reflect.DeepEqual(obj.GetFicReceivedDocuments(), newObj.GetFicReceivedDocuments()))
	newObj.SetFicReceipts(obj.GetFicReceipts())
	assert.True(t, reflect.DeepEqual(obj.GetFicReceipts(), newObj.GetFicReceipts()))
	newObj.SetFicCalendar(obj.GetFicCalendar())
	assert.True(t, reflect.DeepEqual(obj.GetFicCalendar(), newObj.GetFicCalendar()))
	newObj.SetFicArchive(obj.GetFicArchive())
	assert.True(t, reflect.DeepEqual(obj.GetFicArchive(), newObj.GetFicArchive()))
	newObj.SetFicTaxes(obj.GetFicTaxes())
	assert.True(t, reflect.DeepEqual(obj.GetFicTaxes(), newObj.GetFicTaxes()))
	newObj.SetFicStock(obj.GetFicStock())
	assert.True(t, reflect.DeepEqual(obj.GetFicStock(), newObj.GetFicStock()))
	newObj.SetFicCashbook(obj.GetFicCashbook())
	assert.True(t, reflect.DeepEqual(obj.GetFicCashbook(), newObj.GetFicCashbook()))
	newObj.SetFicSettings(obj.GetFicSettings())
	assert.True(t, reflect.DeepEqual(obj.GetFicSettings(), newObj.GetFicSettings()))
	newObj.SetFicEmails(obj.GetFicEmails())
	assert.True(t, reflect.DeepEqual(obj.GetFicEmails(), newObj.GetFicEmails()))
	newObj.SetFicExport(obj.GetFicExport())
	assert.True(t, reflect.DeepEqual(obj.GetFicExport(), newObj.GetFicExport()))
	newObj.SetFicImportBankstatements(obj.GetFicImportBankstatements())
	assert.True(t, reflect.DeepEqual(obj.GetFicImportBankstatements(), newObj.GetFicImportBankstatements()))
	newObj.SetFicImportClientsSuppliers(obj.GetFicImportClientsSuppliers())
	assert.True(t, reflect.DeepEqual(obj.GetFicImportClientsSuppliers(), newObj.GetFicImportClientsSuppliers()))
	newObj.SetFicImportIssuedDocuments(obj.GetFicImportIssuedDocuments())
	assert.True(t, reflect.DeepEqual(obj.GetFicImportIssuedDocuments(), newObj.GetFicImportIssuedDocuments()))
	newObj.SetFicImportProducts(obj.GetFicImportProducts())
	assert.True(t, reflect.DeepEqual(obj.GetFicImportProducts(), newObj.GetFicImportProducts()))
	newObj.SetFicRecurring(obj.GetFicRecurring())
	assert.True(t, reflect.DeepEqual(obj.GetFicRecurring(), newObj.GetFicRecurring()))
	newObj.SetFicRiba(obj.GetFicRiba())
	assert.True(t, reflect.DeepEqual(obj.GetFicRiba(), newObj.GetFicRiba()))
	newObj.SetDicEmployees(obj.GetDicEmployees())
	assert.True(t, reflect.DeepEqual(obj.GetDicEmployees(), newObj.GetDicEmployees()))
	newObj.SetDicSettings(obj.GetDicSettings())
	assert.True(t, reflect.DeepEqual(obj.GetDicSettings(), newObj.GetDicSettings()))
	newObj.SetDicTimesheet(obj.GetDicTimesheet())
	assert.True(t, reflect.DeepEqual(obj.GetDicTimesheet(), newObj.GetDicTimesheet()))
	newObj.SetFicIssuedDocumentsDetailed(obj.GetFicIssuedDocumentsDetailed())
	assert.True(t, reflect.DeepEqual(obj.GetFicIssuedDocumentsDetailed(), newObj.GetFicIssuedDocumentsDetailed()))
}
