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

var GetReceiptPreCreateInfoResponseJsonStr string = "{\"data\":{\"numerations_list\":[\"num\"],\"rc_centers_list\":[\"bg\"],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"}],\"categories_list\":[\"cat6\"],\"vat_types_list\":[{\"id\":1}]}}"

func TestGetReceiptPreCreateInfoResponse(t *testing.T) {
	obj := NewGetReceiptPreCreateInfoResponse()
	json.Unmarshal([]byte(GetReceiptPreCreateInfoResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetReceiptPreCreateInfoResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetReceiptPreCreateInfoResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
