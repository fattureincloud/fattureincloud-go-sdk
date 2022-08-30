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

var PaymentMethodDetailsJsonStr string = "{\"title\":\"details\",\"description\":\"new details\"}"

func TestPaymentMethodDetails(t *testing.T) {
	obj := NewPaymentMethodDetails()
	json.Unmarshal([]byte(PaymentMethodDetailsJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, PaymentMethodDetailsJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewPaymentMethodDetails()
	newObj.SetTitle(obj.GetTitle())
	assert.True(t, reflect.DeepEqual(obj.GetTitle(), newObj.GetTitle()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
}
