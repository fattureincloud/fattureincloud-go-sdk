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

var GetEntityClientPreCreateInfoResponseJsonStr string = "{\"data\":{\"countries_list\":[\"Italia\",\"Marocco\"],\"payment_methods_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"vat_types_list\":[{\"id\":1},{\"id\":2}],\"price_lists\":[{\"id\":\"12345\"}]}}"

func TestGetEntityClientPreCreateInfoResponse(t *testing.T) {
	obj := NewGetEntityClientPreCreateInfoResponse()
	json.Unmarshal([]byte(GetEntityClientPreCreateInfoResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetEntityClientPreCreateInfoResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetEntityClientPreCreateInfoResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
