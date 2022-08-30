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

var ReceivedDocumentItemsListItemJsonStr string = "{\"id\":1,\"product_id\":1,\"code\":\"cod\",\"name\":\"neim\",\"measure\":\"s\",\"net_price\":10,\"category\":\"cat6\",\"qty\":10,\"vat\":{\"id\":1},\"stock\":10}"

func TestReceivedDocumentItemsListItem(t *testing.T) {
	obj := NewReceivedDocumentItemsListItem()
	json.Unmarshal([]byte(ReceivedDocumentItemsListItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentItemsListItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocumentItemsListItem()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetProductId(obj.GetProductId())
	assert.True(t, reflect.DeepEqual(obj.GetProductId(), newObj.GetProductId()))
	newObj.SetCode(obj.GetCode())
	assert.True(t, reflect.DeepEqual(obj.GetCode(), newObj.GetCode()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetMeasure(obj.GetMeasure())
	assert.True(t, reflect.DeepEqual(obj.GetMeasure(), newObj.GetMeasure()))
	newObj.SetNetPrice(obj.GetNetPrice())
	assert.True(t, reflect.DeepEqual(obj.GetNetPrice(), newObj.GetNetPrice()))
	newObj.SetCategory(obj.GetCategory())
	assert.True(t, reflect.DeepEqual(obj.GetCategory(), newObj.GetCategory()))
	newObj.SetQty(obj.GetQty())
	assert.True(t, reflect.DeepEqual(obj.GetQty(), newObj.GetQty()))
	newObj.SetVat(obj.GetVat())
	assert.True(t, reflect.DeepEqual(obj.GetVat(), newObj.GetVat()))
	newObj.SetStock(obj.GetStock())
	assert.True(t, reflect.DeepEqual(obj.GetStock(), newObj.GetStock()))
}
