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

	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

var CompanyInfoPlanInfoFunctionsJsonStr string = "{\"archive\":true,\"cerved\":true,\"document_attachments\":true,\"e_invoice\":true,\"genius\":true,\"mail_tracking\":true,\"payment_notifications\":true,\"paypal\":true,\"receipts\":true,\"recurring\":true,\"smtp\":true,\"sofort\":false,\"stock\":true,\"subaccounts\":true,\"tessera_sanitaria\":true,\"ts_digital\":true,\"ts_invoice_trading\":true,\"ts_pay\":true}"

func TestCompanyInfoPlanInfoFunctions(t *testing.T) {
	obj := NewCompanyInfoPlanInfoFunctions()
	json.Unmarshal([]byte(CompanyInfoPlanInfoFunctionsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyInfoPlanInfoFunctionsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompanyInfoPlanInfoFunctions()
	newObj.SetArchive(obj.GetArchive())
	assert.True(t, reflect.DeepEqual(obj.GetArchive(), newObj.GetArchive()))
	newObj.SetCerved(obj.GetCerved())
	assert.True(t, reflect.DeepEqual(obj.GetCerved(), newObj.GetCerved()))
	newObj.SetDocumentAttachments(obj.GetDocumentAttachments())
	assert.True(t, reflect.DeepEqual(obj.GetDocumentAttachments(), newObj.GetDocumentAttachments()))
	newObj.SetEInvoice(obj.GetEInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetEInvoice(), newObj.GetEInvoice()))
	newObj.SetGenius(obj.GetGenius())
	assert.True(t, reflect.DeepEqual(obj.GetGenius(), newObj.GetGenius()))
	newObj.SetMailTracking(obj.GetMailTracking())
	assert.True(t, reflect.DeepEqual(obj.GetMailTracking(), newObj.GetMailTracking()))
	newObj.SetPaymentNotifications(obj.GetPaymentNotifications())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentNotifications(), newObj.GetPaymentNotifications()))
	newObj.SetPaypal(obj.GetPaypal())
	assert.True(t, reflect.DeepEqual(obj.GetPaypal(), newObj.GetPaypal()))
	newObj.SetReceipts(obj.GetReceipts())
	assert.True(t, reflect.DeepEqual(obj.GetReceipts(), newObj.GetReceipts()))
	newObj.SetRecurring(obj.GetRecurring())
	assert.True(t, reflect.DeepEqual(obj.GetRecurring(), newObj.GetRecurring()))
	newObj.SetSmtp(obj.GetSmtp())
	assert.True(t, reflect.DeepEqual(obj.GetSmtp(), newObj.GetSmtp()))
	newObj.SetSofort(obj.GetSofort())
	assert.True(t, reflect.DeepEqual(obj.GetSofort(), newObj.GetSofort()))
	newObj.SetStock(obj.GetStock())
	assert.True(t, reflect.DeepEqual(obj.GetStock(), newObj.GetStock()))
	newObj.SetSubaccounts(obj.GetSubaccounts())
	assert.True(t, reflect.DeepEqual(obj.GetSubaccounts(), newObj.GetSubaccounts()))
	newObj.SetTesseraSanitaria(obj.GetTesseraSanitaria())
	assert.True(t, reflect.DeepEqual(obj.GetTesseraSanitaria(), newObj.GetTesseraSanitaria()))
	newObj.SetTsDigital(obj.GetTsDigital())
	assert.True(t, reflect.DeepEqual(obj.GetTsDigital(), newObj.GetTsDigital()))
	newObj.SetTsInvoiceTrading(obj.GetTsInvoiceTrading())
	assert.True(t, reflect.DeepEqual(obj.GetTsInvoiceTrading(), newObj.GetTsInvoiceTrading()))
	newObj.SetTsPay(obj.GetTsPay())
	assert.True(t, reflect.DeepEqual(obj.GetTsPay(), newObj.GetTsPay()))
}
