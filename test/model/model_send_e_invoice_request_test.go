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

var SendEInvoiceRequestJsonStr string = "{\"data\":{\"cassa_type\":\"cassa taip\",\"withholding_tax_causal\":\"causal\"}}"

func TestSendEInvoiceRequest(t *testing.T) {
	obj := NewSendEInvoiceRequest()
	json.Unmarshal([]byte(SendEInvoiceRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, SendEInvoiceRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewSendEInvoiceRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
