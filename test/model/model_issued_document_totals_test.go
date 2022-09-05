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

var IssuedDocumentTotalsJsonStr string = "{\"amount_net\":10,\"amount_rivalsa\":10,\"amount_net_with_rivalsa\":10,\"amount_cassa\":10,\"taxable_amount\":10,\"not_taxable_amount\":10,\"amount_vat\":10,\"amount_gross\":10,\"taxable_amount_withholding_tax\":10,\"amount_withholding_tax\":10,\"taxable_amount_other_withholding_tax\":10,\"amount_other_withholding_tax\":10,\"stamp_duty\":10,\"amount_due\":10,\"is_enasarco_maximal_exceeded\":true,\"payments_sum\":10,\"vat_list\":{\"22\":{\"amount_net\":12}}}"

func TestIssuedDocumentTotals(t *testing.T) {
	obj := NewIssuedDocumentTotals()
	json.Unmarshal([]byte(IssuedDocumentTotalsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentTotalsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentTotals()
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountRivalsa(obj.GetAmountRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetAmountRivalsa(), newObj.GetAmountRivalsa()))
	newObj.SetAmountNetWithRivalsa(obj.GetAmountNetWithRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNetWithRivalsa(), newObj.GetAmountNetWithRivalsa()))
	newObj.SetAmountCassa(obj.GetAmountCassa())
	assert.True(t, reflect.DeepEqual(obj.GetAmountCassa(), newObj.GetAmountCassa()))
	newObj.SetTaxableAmount(obj.GetTaxableAmount())
	assert.True(t, reflect.DeepEqual(obj.GetTaxableAmount(), newObj.GetTaxableAmount()))
	newObj.SetNotTaxableAmount(obj.GetNotTaxableAmount())
	assert.True(t, reflect.DeepEqual(obj.GetNotTaxableAmount(), newObj.GetNotTaxableAmount()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetTaxableAmountWithholdingTax(obj.GetTaxableAmountWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetTaxableAmountWithholdingTax(), newObj.GetTaxableAmountWithholdingTax()))
	newObj.SetAmountWithholdingTax(obj.GetAmountWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountWithholdingTax(), newObj.GetAmountWithholdingTax()))
	newObj.SetTaxableAmountOtherWithholdingTax(obj.GetTaxableAmountOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetTaxableAmountOtherWithholdingTax(), newObj.GetTaxableAmountOtherWithholdingTax()))
	newObj.SetAmountOtherWithholdingTax(obj.GetAmountOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOtherWithholdingTax(), newObj.GetAmountOtherWithholdingTax()))
	newObj.SetStampDuty(obj.GetStampDuty())
	assert.True(t, reflect.DeepEqual(obj.GetStampDuty(), newObj.GetStampDuty()))
	newObj.SetAmountDue(obj.GetAmountDue())
	assert.True(t, reflect.DeepEqual(obj.GetAmountDue(), newObj.GetAmountDue()))
	newObj.SetIsEnasarcoMaximalExceeded(obj.GetIsEnasarcoMaximalExceeded())
	assert.True(t, reflect.DeepEqual(obj.GetIsEnasarcoMaximalExceeded(), newObj.GetIsEnasarcoMaximalExceeded()))
	newObj.SetPaymentsSum(obj.GetPaymentsSum())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentsSum(), newObj.GetPaymentsSum()))
	newObj.SetVatList(obj.GetVatList())
	assert.True(t, reflect.DeepEqual(obj.GetVatList(), newObj.GetVatList()))
}
