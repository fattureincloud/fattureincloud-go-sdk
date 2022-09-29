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

var ReceiptItemsListItemJsonStr string = "{\"id\":10,\"amount_net\":10,\"amount_gross\":10,\"category\":\"cat6\",\"vat\":{\"id\":1}}"

func TestReceiptItemsListItem(t *testing.T) {
	obj := NewReceiptItemsListItem()
	json.Unmarshal([]byte(ReceiptItemsListItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceiptItemsListItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceiptItemsListItem()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetVat(obj.GetVat())
	assert.True(t, reflect.DeepEqual(obj.GetVat(), newObj.GetVat()))
}
