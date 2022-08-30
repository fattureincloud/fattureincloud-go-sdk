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

var IssuedDocumentItemsListItemJsonStr string = "{\"product_id\":12345,\"code\":\"cod3\",\"name\":\"prodott1\",\"category\":\"cat5\",\"description\":\"primo\",\"qty\":99,\"measure\":\"big\",\"net_price\":10,\"gross_price\":10,\"vat\":{\"id\":1},\"not_taxable\":true,\"apply_withholding_taxes\":true,\"discount\":0,\"discount_highlight\":true,\"in_ddt\":true,\"stock\":true}"

func TestIssuedDocumentItemsListItem(t *testing.T) {
	obj := NewIssuedDocumentItemsListItem()
	json.Unmarshal([]byte(IssuedDocumentItemsListItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentItemsListItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentItemsListItem()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetProductId(obj.GetProductId())
	assert.True(t, reflect.DeepEqual(obj.GetProductId(), newObj.GetProductId()))
	newObj.SetCode(obj.GetCode())
	assert.True(t, reflect.DeepEqual(obj.GetCode(), newObj.GetCode()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetQty(obj.GetQty())
	assert.True(t, reflect.DeepEqual(obj.GetQty(), newObj.GetQty()))
	newObj.SetMeasure(obj.GetMeasure())
	assert.True(t, reflect.DeepEqual(obj.GetMeasure(), newObj.GetMeasure()))
	newObj.SetNetPrice(obj.GetNetPrice())
	assert.True(t, reflect.DeepEqual(obj.GetNetPrice(), newObj.GetNetPrice()))
	newObj.SetGrossPrice(obj.GetGrossPrice())
	assert.True(t, reflect.DeepEqual(obj.GetGrossPrice(), newObj.GetGrossPrice()))
	newObj.SetVat(obj.GetVat())
	assert.True(t, reflect.DeepEqual(obj.GetVat(), newObj.GetVat()))
	newObj.SetNotTaxable(obj.GetNotTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetNotTaxable(), newObj.GetNotTaxable()))
	newObj.SetApplyWithholdingTaxes(obj.GetApplyWithholdingTaxes())
	assert.True(t, reflect.DeepEqual(obj.GetApplyWithholdingTaxes(), newObj.GetApplyWithholdingTaxes()))
	newObj.SetDiscount(obj.GetDiscount())
	assert.True(t, reflect.DeepEqual(obj.GetDiscount(), newObj.GetDiscount()))
	newObj.SetDiscountHighlight(obj.GetDiscountHighlight())
	assert.True(t, reflect.DeepEqual(obj.GetDiscountHighlight(), newObj.GetDiscountHighlight()))
	newObj.SetInDdt(obj.GetInDdt())
	assert.True(t, reflect.DeepEqual(obj.GetInDdt(), newObj.GetInDdt()))
	newObj.SetStock(obj.GetStock())
	assert.True(t, reflect.DeepEqual(obj.GetStock(), newObj.GetStock()))
	newObj.SetEiRaw(obj.GetEiRaw())
	assert.True(t, reflect.DeepEqual(obj.GetEiRaw(), newObj.GetEiRaw()))
}
