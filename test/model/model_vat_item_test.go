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

var VatItemJsonStr string = "{\"amount_net\":10,\"amount_vat\":12}"

func TestVatItem(t *testing.T) {
	obj := NewVatItem()
	json.Unmarshal([]byte(VatItemJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, VatItemJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewVatItem()
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
}
