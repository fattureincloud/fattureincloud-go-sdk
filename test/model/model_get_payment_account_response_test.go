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

var GetPaymentAccountResponseJsonStr string = "{\"data\":{\"id\":12345,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false}}"

func TestGetPaymentAccountResponse(t *testing.T) {
	obj := NewGetPaymentAccountResponse()
	json.Unmarshal([]byte(GetPaymentAccountResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetPaymentAccountResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetPaymentAccountResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
