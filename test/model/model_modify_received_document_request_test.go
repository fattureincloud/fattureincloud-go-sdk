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

var ModifyReceivedDocumentRequestJsonStr string = "{\"data\":{\"id\":1,\"type\":\"expense\",\"entity\":{\"id\":1,\"name\":\"neim\"},\"date\":\"2021-12-25\",\"category\":\"cat6\",\"description\":\"Ricarica\",\"amount_net\":10,\"amount_vat\":0,\"amount_withholding_tax\":0,\"amount_other_withholding_tax\":0,\"amortization\":10,\"rc_center\":\"bg\",\"invoice_number\":\"in\",\"is_marked\":false,\"is_detailed\":false,\"e_invoice\":false,\"currency\":{\"id\":\"EUR\",\"exchange_rate\":\"1\"},\"tax_deductibility\":50,\"vat_deductibility\":100,\"items_list\":[{\"id\":1}],\"payments_list\":[{\"amount\":592,\"due_date\":\"2021-12-25\",\"paid_date\":\"2021-12-25\",\"payment_terms\":{\"days\":0,\"type\":\"standard\"},\"status\":\"paid\",\"payment_account\":{\"id\":21,\"type\":\"standard\"}}],\"attachment_token\":\"dGdweHdjNjlieWFjY3BseGZ0cTZmbWN0Njhhb3R0cXQvZmlsZW5hbWVfZXhhbXBsZQ\\u003d\\u003d\"}}"

func TestModifyReceivedDocumentRequest(t *testing.T) {
	obj := NewModifyReceivedDocumentRequest()
	json.Unmarshal([]byte(ModifyReceivedDocumentRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ModifyReceivedDocumentRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewModifyReceivedDocumentRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
