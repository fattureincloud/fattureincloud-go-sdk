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

var ReceivedDocumentInfoJsonStr string = "{\"default_values\":{\"detailed\":false},\"items_default_values\":{\"vat\":10},\"countries_list\":[\"Italia\"],\"currencies_list\":[{\"id\":\"EUR\"}],\"categories_list\":[\"cat6\"],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"}],\"vat_types_list\":[{\"id\":1}]}"

func TestReceivedDocumentInfo(t *testing.T) {
	obj := NewReceivedDocumentInfo()
	json.Unmarshal([]byte(ReceivedDocumentInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ReceivedDocumentInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewReceivedDocumentInfo()
	newObj.SetDefaultValues(obj.GetDefaultValues())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultValues(), newObj.GetDefaultValues()))
	newObj.SetItemsDefaultValues(obj.GetItemsDefaultValues())
	assert.True(t, reflect.DeepEqual(obj.GetItemsDefaultValues(), newObj.GetItemsDefaultValues()))
	newObj.SetCountriesList(obj.GetCountriesList())
	assert.True(t, reflect.DeepEqual(obj.GetCountriesList(), newObj.GetCountriesList()))
	newObj.SetCurrenciesList(obj.GetCurrenciesList())
	assert.True(t, reflect.DeepEqual(obj.GetCurrenciesList(), newObj.GetCurrenciesList()))
	newObj.SetCategoriesList(obj.GetCategoriesList())
	assert.True(t, reflect.DeepEqual(obj.GetCategoriesList(), newObj.GetCategoriesList()))
	newObj.SetPaymentAccountsList(obj.GetPaymentAccountsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountsList(), newObj.GetPaymentAccountsList()))
	newObj.SetVatTypesList(obj.GetVatTypesList())
	assert.True(t, reflect.DeepEqual(obj.GetVatTypesList(), newObj.GetVatTypesList()))
}
