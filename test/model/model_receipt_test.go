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

var ReceiptJsonStr string = "{\"id\":10,\"date\":\"2021-12-25\",\"number\":10,\"numeration\":\"num\",\"amount_net\":10,\"amount_vat\":10,\"amount_gross\":10,\"use_gross_prices\":true,\"type\":\"till_receipt\",\"description\":\"descr\",\"rc_center\":\"bg\",\"created_at\":\"2021-10-10\",\"updated_at\":\"2021-10-10\",\"payment_account\":{\"id\":1,\"type\":\"standard\"},\"items_list\":[{\"id\":1}]}"

func TestReceipt(t *testing.T) {
	obj := NewReceipt()
	json.Unmarshal([]byte(ReceiptJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceiptJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceipt()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
	newObj.SetNumber(obj.GetNumber())
	assert.True(t, reflect.DeepEqual(obj.GetNumber(), newObj.GetNumber()))
	newObj.SetNumeration(obj.GetNumeration())
	assert.True(t, reflect.DeepEqual(obj.GetNumeration(), newObj.GetNumeration()))
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetUseGrossPrices(obj.GetUseGrossPrices())
	assert.True(t, reflect.DeepEqual(obj.GetUseGrossPrices(), newObj.GetUseGrossPrices()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetRcCenter(obj.GetRcCenter())
	assert.True(t, reflect.DeepEqual(obj.GetRcCenter(), newObj.GetRcCenter()))
	newObj.SetCreatedAt(obj.GetCreatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetCreatedAt(), newObj.GetCreatedAt()))
	newObj.SetUpdatedAt(obj.GetUpdatedAt())
	assert.True(t, reflect.DeepEqual(obj.GetUpdatedAt(), newObj.GetUpdatedAt()))
	newObj.SetPaymentAccount(obj.GetPaymentAccount())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccount(), newObj.GetPaymentAccount()))
	newObj.SetItemsList(obj.GetItemsList())
	assert.True(t, reflect.DeepEqual(obj.GetItemsList(), newObj.GetItemsList()))
}
