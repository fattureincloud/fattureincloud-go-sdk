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

var ListPriceListsResponseJsonStr string = "{\"data\":[{\"id\":\"10\",\"name\":\"Listino 1\",\"prices_type\":\"net\",\"is_default\":true,\"valid_from\":\"2023-01-01\",\"valid_to\":\"2023-12-31\",\"type\":\"sell\"}]}"

func TestListPriceListsResponse(t *testing.T) {
	obj := NewListPriceListsResponse()
	json.Unmarshal([]byte(ListPriceListsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListPriceListsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListPriceListsResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
