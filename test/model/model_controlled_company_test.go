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

var ControlledCompanyJsonStr string = "{\"id\":5,\"name\":\"Azienda 1\",\"type\":\"company\",\"access_token\":\"ABCDEF\",\"connection_id\":94817,\"tax_code\":\"SLVMTT50A01F205L\"}"

func TestControlledCompany(t *testing.T) {
	obj := NewControlledCompany()
	json.Unmarshal([]byte(ControlledCompanyJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ControlledCompanyJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewControlledCompany()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetAccessToken(obj.GetAccessToken())
	assert.True(t, reflect.DeepEqual(obj.GetAccessToken(), newObj.GetAccessToken()))
	newObj.SetConnectionId(obj.GetConnectionId())
	assert.True(t, reflect.DeepEqual(obj.GetConnectionId(), newObj.GetConnectionId()))
	newObj.SetTaxCode(obj.GetTaxCode())
	assert.True(t, reflect.DeepEqual(obj.GetTaxCode(), newObj.GetTaxCode()))
}
