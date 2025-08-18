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

var PriceListJsonStr string = "{\"id\":\"10\",\"name\":\"Listino 1\",\"prices_type\":\"net\",\"is_default\":true,\"valid_from\":\"2023-01-01\",\"valid_to\":\"2023-12-31\",\"type\":\"sell\"}"

func TestPriceList(t *testing.T) {
	obj := NewPriceList()
	json.Unmarshal([]byte(PriceListJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PriceListJsonStr, string(jsonStr2), "they should be equal")
	newObj := NewPriceList()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetPricesType(obj.GetPricesType())
	assert.True(t, reflect.DeepEqual(obj.GetPricesType(), newObj.GetPricesType()))
	newObj.SetIsDefault(obj.GetIsDefault())
	assert.True(t, reflect.DeepEqual(obj.GetIsDefault(), newObj.GetIsDefault()))
	newObj.SetValidFrom(obj.GetValidFrom())
	assert.True(t, reflect.DeepEqual(obj.GetValidFrom(), newObj.GetValidFrom()))
	newObj.SetValidTo(obj.GetValidTo())
	assert.True(t, reflect.DeepEqual(obj.GetValidTo(), newObj.GetValidTo()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
}
