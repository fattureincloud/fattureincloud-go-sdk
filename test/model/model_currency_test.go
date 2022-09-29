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

var CurrencyJsonStr string = "{\"id\":\"EUR\",\"symbol\":\"€\",\"exchange_rate\":\"1.00000\",\"html_symbol\":\"\\u0026euro;\"}"

func TestCurrency(t *testing.T) {
	obj := NewCurrency()
	json.Unmarshal([]byte(CurrencyJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CurrencyJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCurrency()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetSymbol(obj.GetSymbol())
	assert.True(t, reflect.DeepEqual(obj.GetSymbol(), newObj.GetSymbol()))
	newObj.SetExchangeRate(obj.GetExchangeRate())
	assert.True(t, reflect.DeepEqual(obj.GetExchangeRate(), newObj.GetExchangeRate()))
	newObj.SetHtmlSymbol(obj.GetHtmlSymbol())
	assert.True(t, reflect.DeepEqual(obj.GetHtmlSymbol(), newObj.GetHtmlSymbol()))
}
