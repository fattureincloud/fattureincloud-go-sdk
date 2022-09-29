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

var CreateVatTypeResponseJsonStr string = "{\"data\":{\"id\":12345,\"value\":22,\"description\":\"Non imponibile art. 123\",\"notes\":\"IVA non imponibile\",\"e_invoice\":true,\"ei_type\":\"2\",\"ei_description\":\"desc\",\"is_disabled\":false}}"

func TestCreateVatTypeResponse(t *testing.T) {
	obj := NewCreateVatTypeResponse()
	json.Unmarshal([]byte(CreateVatTypeResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CreateVatTypeResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCreateVatTypeResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
