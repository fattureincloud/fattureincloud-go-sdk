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

var EntityClientPreCreateInfoJsonStr string = "{\"countries_list\":[\"Italia\",\"Marocco\"],\"payment_methods_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"vat_types_list\":[{\"id\":1},{\"id\":2}],\"price_lists\":[{\"id\":\"12345\"}]}"

func TestEntityClientPreCreateInfo(t *testing.T) {
	obj := NewEntityClientPreCreateInfo()
	json.Unmarshal([]byte(EntityClientPreCreateInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EntityClientPreCreateInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEntityClientPreCreateInfo()
	newObj.SetCountriesList(obj.GetCountriesList())
	assert.True(t, reflect.DeepEqual(obj.GetCountriesList(), newObj.GetCountriesList()))
	newObj.SetPaymentMethodsList(obj.GetPaymentMethodsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentMethodsList(), newObj.GetPaymentMethodsList()))
	newObj.SetPaymentAccountsList(obj.GetPaymentAccountsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountsList(), newObj.GetPaymentAccountsList()))
	newObj.SetVatTypesList(obj.GetVatTypesList())
	assert.True(t, reflect.DeepEqual(obj.GetVatTypesList(), newObj.GetVatTypesList()))
	newObj.SetPriceLists(obj.GetPriceLists())
	assert.True(t, reflect.DeepEqual(obj.GetPriceLists(), newObj.GetPriceLists()))
}
