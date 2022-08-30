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

var GetPaymentMethodResponseJsonStr string = "{\"data\":{\"id\":12345,\"name\":\"Bonifico bancario\",\"type\":\"standard\",\"is_default\":true,\"default_payment_account\":{\"id\":21,\"name\":\"n1\",\"type\":\"standard\"},\"details\":[{\"title\":\"t1\"}],\"bank_iban\":\"IT62W0300203280486429468578\",\"bank_name\":\"Indesa\",\"bank_beneficiary\":\"mamma\",\"ei_payment_method\":\"2\"}}"

func TestGetPaymentMethodResponse(t *testing.T) {
	obj := NewGetPaymentMethodResponse()
	json.Unmarshal([]byte(GetPaymentMethodResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetPaymentMethodResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetPaymentMethodResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
