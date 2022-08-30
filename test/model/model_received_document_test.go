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

var ReceivedDocumentJsonStr string = "{\"id\":1,\"type\":\"expense\",\"entity\":{\"id\":1,\"name\":\"neim\"},\"date\":\"2021-12-25\",\"category\":\"cat6\",\"description\":\"Ricarica\",\"amount_net\":10,\"amount_vat\":0,\"amount_withholding_tax\":0,\"amount_other_withholding_tax\":0,\"amortization\":10,\"rc_center\":\"bg\",\"invoice_number\":\"in\",\"is_marked\":false,\"is_detailed\":false,\"e_invoice\":false,\"currency\":{\"id\":\"EUR\",\"exchange_rate\":\"1\"},\"tax_deductibility\":50,\"vat_deductibility\":100,\"items_list\":[{\"id\":1}],\"payments_list\":[{\"amount\":592,\"due_date\":\"2021-12-25\",\"paid_date\":\"2021-12-25\",\"payment_terms\":{\"days\":0,\"type\":\"standard\"},\"status\":\"paid\",\"payment_account\":{\"id\":21,\"type\":\"standard\"}}],\"attachment_token\":\"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ\\u003d\\u003d\"}"

func TestReceivedDocument(t *testing.T) {
	obj := NewReceivedDocument()
	json.Unmarshal([]byte(ReceivedDocumentJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocument()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetEntity(obj.GetEntity())
	assert.True(t, reflect.DeepEqual(obj.GetEntity(), newObj.GetEntity()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
	newObj.SetAmountWithholdingTax(obj.GetAmountWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountWithholdingTax(), newObj.GetAmountWithholdingTax()))
	newObj.SetAmountOtherWithholdingTax(obj.GetAmountOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOtherWithholdingTax(), newObj.GetAmountOtherWithholdingTax()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetAmortization(obj.GetAmortization())
	assert.True(t, reflect.DeepEqual(obj.GetAmortization(), newObj.GetAmortization()))
	newObj.SetRcCenter(obj.GetRcCenter())
	assert.True(t, reflect.DeepEqual(obj.GetRcCenter(), newObj.GetRcCenter()))
	newObj.SetInvoiceNumber(obj.GetInvoiceNumber())
	assert.True(t, reflect.DeepEqual(obj.GetInvoiceNumber(), newObj.GetInvoiceNumber()))
	newObj.SetIsMarked(obj.GetIsMarked())
	assert.True(t, reflect.DeepEqual(obj.GetIsMarked(), newObj.GetIsMarked()))
	newObj.SetIsDetailed(obj.GetIsDetailed())
	assert.True(t, reflect.DeepEqual(obj.GetIsDetailed(), newObj.GetIsDetailed()))
	newObj.SetEInvoice(obj.GetEInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetEInvoice(), newObj.GetEInvoice()))
	newObj.SetNextDueDate(obj.GetNextDueDate())
	assert.True(t, reflect.DeepEqual(obj.GetNextDueDate(), newObj.GetNextDueDate()))
	newObj.SetCurrency(obj.GetCurrency())
	assert.True(t, reflect.DeepEqual(obj.GetCurrency(), newObj.GetCurrency()))
	newObj.SetTaxDeductibility(obj.GetTaxDeductibility())
	assert.True(t, reflect.DeepEqual(obj.GetTaxDeductibility(), newObj.GetTaxDeductibility()))
	newObj.SetVatDeductibility(obj.GetVatDeductibility())
	assert.True(t, reflect.DeepEqual(obj.GetVatDeductibility(), newObj.GetVatDeductibility()))
	newObj.SetItemsList(obj.GetItemsList())
	assert.True(t, reflect.DeepEqual(obj.GetItemsList(), newObj.GetItemsList()))
	newObj.SetPaymentsList(obj.GetPaymentsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentsList(), newObj.GetPaymentsList()))
	newObj.SetAttachmentUrl(obj.GetAttachmentUrl())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentUrl(), newObj.GetAttachmentUrl()))
	newObj.SetAttachmentPreviewUrl(obj.GetAttachmentPreviewUrl())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentPreviewUrl(), newObj.GetAttachmentPreviewUrl()))
	newObj.SetAttachmentToken(obj.GetAttachmentToken())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentToken(), newObj.GetAttachmentToken()))
}
