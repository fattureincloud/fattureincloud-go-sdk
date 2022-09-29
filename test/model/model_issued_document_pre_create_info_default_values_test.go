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

var IssuedDocumentPreCreateInfoDefaultValuesJsonStr string = "{\"default_template\":{\"id\":1},\"dn_template\":{\"id\":1},\"ai_template\":{\"id\":1},\"notes\":\"notes\",\"rivalsa\":1,\"cassa\":1,\"withholding_tax\":1,\"withholding_tax_taxable\":1,\"other_withholding_tax\":1,\"use_gross_prices\":true,\"payment_method\":{\"id\":1,\"type\":\"standard\"}}"

func TestIssuedDocumentPreCreateInfoDefaultValues(t *testing.T) {
	obj := NewIssuedDocumentPreCreateInfoDefaultValues()
	json.Unmarshal([]byte(IssuedDocumentPreCreateInfoDefaultValuesJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentPreCreateInfoDefaultValuesJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentPreCreateInfoDefaultValues()
	newObj.SetDefaultTemplate(obj.GetDefaultTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultTemplate(), newObj.GetDefaultTemplate()))
	newObj.SetDnTemplate(obj.GetDnTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetDnTemplate(), newObj.GetDnTemplate()))
	newObj.SetAiTemplate(obj.GetAiTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetAiTemplate(), newObj.GetAiTemplate()))
	newObj.SetNotes(obj.GetNotes())
	assert.True(t, reflect.DeepEqual(obj.GetNotes(), newObj.GetNotes()))
	newObj.SetRivalsa(obj.GetRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetRivalsa(), newObj.GetRivalsa()))
	newObj.SetCassa(obj.GetCassa())
	assert.True(t, reflect.DeepEqual(obj.GetCassa(), newObj.GetCassa()))
	newObj.SetWithholdingTax(obj.GetWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetWithholdingTax(), newObj.GetWithholdingTax()))
	newObj.SetWithholdingTaxTaxable(obj.GetWithholdingTaxTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetWithholdingTaxTaxable(), newObj.GetWithholdingTaxTaxable()))
	newObj.SetOtherWithholdingTax(obj.GetOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetOtherWithholdingTax(), newObj.GetOtherWithholdingTax()))
	newObj.SetUseGrossPrices(obj.GetUseGrossPrices())
	assert.True(t, reflect.DeepEqual(obj.GetUseGrossPrices(), newObj.GetUseGrossPrices()))
	newObj.SetPaymentMethod(obj.GetPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentMethod(), newObj.GetPaymentMethod()))
}
