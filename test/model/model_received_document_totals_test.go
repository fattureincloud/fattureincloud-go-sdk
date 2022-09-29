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

var ReceivedDocumentTotalsJsonStr string = "{\"amount_net\":10,\"amount_vat\":10,\"amount_gross\":10,\"amount_withholding_tax\":10,\"amount_other_withholding_tax\":10,\"amount_due\":10,\"payments_sum\":10}"

func TestReceivedDocumentTotals(t *testing.T) {
	obj := NewReceivedDocumentTotals()
	json.Unmarshal([]byte(ReceivedDocumentTotalsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentTotalsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocumentTotals()
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetAmountWithholdingTax(obj.GetAmountWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountWithholdingTax(), newObj.GetAmountWithholdingTax()))
	newObj.SetAmountOtherWithholdingTax(obj.GetAmountOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOtherWithholdingTax(), newObj.GetAmountOtherWithholdingTax()))
	newObj.SetAmountDue(obj.GetAmountDue())
	assert.True(t, reflect.DeepEqual(obj.GetAmountDue(), newObj.GetAmountDue()))
	newObj.SetPaymentsSum(obj.GetPaymentsSum())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentsSum(), newObj.GetPaymentsSum()))
}
