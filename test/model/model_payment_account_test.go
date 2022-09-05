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

var PaymentAccountJsonStr string = "{\"id\":21,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false}"

func TestPaymentAccount(t *testing.T) {
	obj := NewPaymentAccount()
	json.Unmarshal([]byte(PaymentAccountJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PaymentAccountJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewPaymentAccount()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetIban(obj.GetIban())
	assert.True(t, reflect.DeepEqual(obj.GetIban(), newObj.GetIban()))
	newObj.SetSia(obj.GetSia())
	assert.True(t, reflect.DeepEqual(obj.GetSia(), newObj.GetSia()))
	newObj.SetCuc(obj.GetCuc())
	assert.True(t, reflect.DeepEqual(obj.GetCuc(), newObj.GetCuc()))
	newObj.SetVirtual(obj.GetVirtual())
	assert.True(t, reflect.DeepEqual(obj.GetVirtual(), newObj.GetVirtual()))
}
