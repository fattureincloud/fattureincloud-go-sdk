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

var ProductJsonStr string = "{\"id\":1,\"name\":\"neim\",\"code\":\"cod\",\"net_price\":10,\"gross_price\":10,\"use_gross_price\":true,\"default_vat\":{\"id\":1},\"net_cost\":10,\"measure\":\"big\",\"description\":\"desc\",\"category\":\"cat6\",\"notes\":\"nots\",\"in_stock\":true,\"stock_initial\":10,\"average_cost\":10,\"average_price\":10,\"created_at\":\"2021-10-10\",\"updated_at\":\"2021-10-10\"}"

func TestProduct(t *testing.T) {
	obj := NewProduct()
	json.Unmarshal([]byte(ProductJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ProductJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewProduct()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetCode(obj.GetCode())
	assert.True(t, reflect.DeepEqual(obj.GetCode(), newObj.GetCode()))
	newObj.SetNetPrice(obj.GetNetPrice())
	assert.True(t, reflect.DeepEqual(obj.GetNetPrice(), newObj.GetNetPrice()))
	newObj.SetGrossPrice(obj.GetGrossPrice())
	assert.True(t, reflect.DeepEqual(obj.GetGrossPrice(), newObj.GetGrossPrice()))
	newObj.SetUseGrossPrice(obj.GetUseGrossPrice())
	assert.True(t, reflect.DeepEqual(obj.GetUseGrossPrice(), newObj.GetUseGrossPrice()))
	newObj.SetDefaultVat(obj.GetDefaultVat())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultVat(), newObj.GetDefaultVat()))
	newObj.SetNetCost(obj.GetNetCost())
	assert.True(t, reflect.DeepEqual(obj.GetNetCost(), newObj.GetNetCost()))
	newObj.SetMeasure(obj.GetMeasure())
	assert.True(t, reflect.DeepEqual(obj.GetMeasure(), newObj.GetMeasure()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetNotes(obj.GetNotes())
	assert.True(t, reflect.DeepEqual(obj.GetNotes(), newObj.GetNotes()))
	newObj.SetInStock(obj.GetInStock())
	assert.True(t, reflect.DeepEqual(obj.GetInStock(), newObj.GetInStock()))
	newObj.SetStockInitial(obj.GetStockInitial())
	assert.True(t, reflect.DeepEqual(obj.GetStockInitial(), newObj.GetStockInitial()))
	newObj.SetStockCurrent(obj.GetStockCurrent())
	assert.True(t, reflect.DeepEqual(obj.GetStockCurrent(), newObj.GetStockCurrent()))
	newObj.SetAverageCost(obj.GetAverageCost())
	assert.True(t, reflect.DeepEqual(obj.GetAverageCost(), newObj.GetAverageCost()))
	newObj.SetAveragePrice(obj.GetAveragePrice())
	assert.True(t, reflect.DeepEqual(obj.GetAveragePrice(), newObj.GetAveragePrice()))
	newObj.SetCreatedAt(obj.GetCreatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetCreatedAt(), newObj.GetCreatedAt()))
	newObj.SetUpdatedAt(obj.GetUpdatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetUpdatedAt(), newObj.GetUpdatedAt()))
}
