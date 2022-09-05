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

var PermissionsFicIssuedDocumentsDetailedJsonStr string = "{\"quotes\":\"write\",\"proformas\":\"write\",\"invoices\":\"write\",\"receipts\":\"write\",\"delivery_notes\":\"write\",\"credit_notes\":\"write\",\"orders\":\"write\",\"work_reports\":\"write\",\"supplier_orders\":\"write\",\"self_invoices\":\"write\"}"

func TestPermissionsFicIssuedDocumentsDetailed(t *testing.T) {
	obj := NewPermissionsFicIssuedDocumentsDetailed()
	json.Unmarshal([]byte(PermissionsFicIssuedDocumentsDetailedJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PermissionsFicIssuedDocumentsDetailedJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewPermissionsFicIssuedDocumentsDetailed()
	newObj.SetQuotes(obj.GetQuotes())
	assert.True(t, reflect.DeepEqual(obj.GetQuotes(), newObj.GetQuotes()))
	newObj.SetProformas(obj.GetProformas())
	assert.True(t, reflect.DeepEqual(obj.GetProformas(), newObj.GetProformas()))
	newObj.SetInvoices(obj.GetInvoices())
	assert.True(t, reflect.DeepEqual(obj.GetInvoices(), newObj.GetInvoices()))
	newObj.SetReceipts(obj.GetReceipts())
	assert.True(t, reflect.DeepEqual(obj.GetReceipts(), newObj.GetReceipts()))
	newObj.SetDeliveryNotes(obj.GetDeliveryNotes())
	assert.True(t, reflect.DeepEqual(obj.GetDeliveryNotes(), newObj.GetDeliveryNotes()))
	newObj.SetCreditNotes(obj.GetCreditNotes())
	assert.True(t, reflect.DeepEqual(obj.GetCreditNotes(), newObj.GetCreditNotes()))
	newObj.SetOrders(obj.GetOrders())
	assert.True(t, reflect.DeepEqual(obj.GetOrders(), newObj.GetOrders()))
	newObj.SetWorkReports(obj.GetWorkReports())
	assert.True(t, reflect.DeepEqual(obj.GetWorkReports(), newObj.GetWorkReports()))
	newObj.SetSupplierOrders(obj.GetSupplierOrders())
	assert.True(t, reflect.DeepEqual(obj.GetSupplierOrders(), newObj.GetSupplierOrders()))
	newObj.SetSelfInvoices(obj.GetSelfInvoices())
	assert.True(t, reflect.DeepEqual(obj.GetSelfInvoices(), newObj.GetSelfInvoices()))
}
