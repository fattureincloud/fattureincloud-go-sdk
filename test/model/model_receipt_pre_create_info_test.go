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

var ReceiptPreCreateInfoJsonStr string = "{\"numerations\":{\"2022\":{\"cloud\":1988,\"in\":2,\"fatture\":22},\"2021\":{\"a\":1,\"b\":2,\"c\":3}},\"numerations_list\":[\"num\"],\"rc_centers_list\":[\"bg\"],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"}],\"categories_list\":[\"cat6\"],\"vat_types_list\":[{\"id\":1}]}"

func TestReceiptPreCreateInfo(t *testing.T) {
	obj := NewReceiptPreCreateInfo()
	json.Unmarshal([]byte(ReceiptPreCreateInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceiptPreCreateInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceiptPreCreateInfo()
	newObj.SetNumerations(obj.GetNumerations())
	assert.True(t, reflect.DeepEqual(obj.GetNumerations(), newObj.GetNumerations()))
	newObj.SetNumerationsList(obj.GetNumerationsList())
	assert.True(t, reflect.DeepEqual(obj.GetNumerationsList(), newObj.GetNumerationsList()))
	newObj.SetRcCentersList(obj.GetRcCentersList())
	assert.True(t, reflect.DeepEqual(obj.GetRcCentersList(), newObj.GetRcCentersList()))
	newObj.SetPaymentAccountsList(obj.GetPaymentAccountsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountsList(), newObj.GetPaymentAccountsList()))
	newObj.SetCategoriesList(obj.GetCategoriesList())
	assert.True(t, reflect.DeepEqual(obj.GetCategoriesList(), newObj.GetCategoriesList()))
	newObj.SetVatTypesList(obj.GetVatTypesList())
	assert.True(t, reflect.DeepEqual(obj.GetVatTypesList(), newObj.GetVatTypesList()))
}
