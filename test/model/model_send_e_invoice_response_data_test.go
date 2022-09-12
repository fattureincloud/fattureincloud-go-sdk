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

var SendEInvoiceResponseDataJsonStr string = "{\"name\":\"neim\",\"date\":\"2021-12-31\"}"

func TestSendEInvoiceResponseData(t *testing.T) {
	obj := NewSendEInvoiceResponseData()
	json.Unmarshal([]byte(SendEInvoiceResponseDataJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, SendEInvoiceResponseDataJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewSendEInvoiceResponseData()
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
}