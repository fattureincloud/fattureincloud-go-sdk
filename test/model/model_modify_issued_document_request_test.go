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

var ModifyIssuedDocumentRequestJsonStr string = "{\"data\":{\"id\":12345,\"type\":\"invoice\",\"notes\":\"bando\",\"show_totals\":\"all\"},\"options\":{\"fix_payments\":true}}"

func TestModifyIssuedDocumentRequest(t *testing.T) {
	obj := NewModifyIssuedDocumentRequest()
	json.Unmarshal([]byte(ModifyIssuedDocumentRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ModifyIssuedDocumentRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewModifyIssuedDocumentRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
	newObj.SetOptions(obj.GetOptions())
	assert.True(t, reflect.DeepEqual(obj.GetOptions(), newObj.GetOptions()))
}
