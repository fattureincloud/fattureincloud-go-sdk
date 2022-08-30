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

var CreateReceiptRequestJsonStr string = "{\"data\":{\"id\":12345,\"date\":\"2021-12-25\",\"number\":10,\"numeration\":\"num\",\"amount_net\":10,\"amount_vat\":10,\"amount_gross\":10,\"use_gross_prices\":true,\"type\":\"till_receipt\",\"description\":\"descr\",\"rc_center\":\"bg\",\"created_at\":\"2021-10-10\",\"updated_at\":\"2021-10-10\",\"payment_account\":{\"id\":1,\"type\":\"standard\"},\"items_list\":[{\"id\":1}]},\"autocomplete_number\":true}"

func TestCreateReceiptRequest(t *testing.T) {
	obj := NewCreateReceiptRequest()
	json.Unmarshal([]byte(CreateReceiptRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CreateReceiptRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCreateReceiptRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
	newObj.SetAutocompleteNumber(obj.GetAutocompleteNumber())
	assert.True(t, reflect.DeepEqual(obj.GetAutocompleteNumber(), newObj.GetAutocompleteNumber()))
}
