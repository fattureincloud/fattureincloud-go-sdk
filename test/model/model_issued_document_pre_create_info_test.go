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

var IssuedDocumentPreCreateInfoJsonStr string = "{\"numerations\":{\"2022\":{\"cloud\":1988,\"in\":2,\"fatture\":22},\"2021\":{\"a\":1,\"b\":2,\"c\":3}},\"dn_numerations\":{\"2022\":{\"a\":0,\"b\":11,\"c\":222},\"2021\":{\"d\":3333,\"e\":44444,\"f\":555555}},\"default_values\":{\"default_template\":{\"id\":1},\"dn_template\":{\"id\":1},\"ai_template\":{\"id\":1},\"notes\":\"notes\",\"rivalsa\":1,\"cassa\":1,\"withholding_tax\":1,\"withholding_tax_taxable\":1,\"other_withholding_tax\":1,\"use_gross_prices\":true,\"payment_method\":{\"id\":1,\"type\":\"standard\"}},\"extra_data_default_values\":{\"ts_communication\":true,\"ts_tipo_spesa\":\"ts\",\"ts_flag_tipo_spesa\":1,\"ts_pagamento_tracciato\":true},\"items_default_values\":{\"vat\":{\"id\":1}},\"countries_list\":[\"Italia\",\"Marocco\"],\"currencies_list\":[{\"id\":\"EUR\"},{\"id\":\"DNR\"}],\"templates_list\":[{\"id\":1},{\"id\":2}],\"dn_templates_list\":[{\"id\":1},{\"id\":2}],\"ai_templates_list\":[{\"id\":1},{\"id\":2}],\"payment_methods_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"vat_types_list\":[{\"id\":1},{\"id\":2}]}"

func TestIssuedDocumentPreCreateInfo(t *testing.T) {
	obj := NewIssuedDocumentPreCreateInfo()
	json.Unmarshal([]byte(IssuedDocumentPreCreateInfoJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentPreCreateInfoJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentPreCreateInfo()
	newObj.SetNumerations(obj.GetNumerations())
	assert.True(t, reflect.DeepEqual(obj.GetNumerations(), newObj.GetNumerations()))
	newObj.SetDnNumerations(obj.GetDnNumerations())
	assert.True(t, reflect.DeepEqual(obj.GetDnNumerations(), newObj.GetDnNumerations()))
	newObj.SetDefaultValues(obj.GetDefaultValues())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultValues(), newObj.GetDefaultValues()))
	newObj.SetExtraDataDefaultValues(obj.GetExtraDataDefaultValues())
	assert.True(t, reflect.DeepEqual(obj.GetExtraDataDefaultValues(), newObj.GetExtraDataDefaultValues()))
	newObj.SetItemsDefaultValues(obj.GetItemsDefaultValues())
	assert.True(t, reflect.DeepEqual(obj.GetItemsDefaultValues(), newObj.GetItemsDefaultValues()))
	newObj.SetCountriesList(obj.GetCountriesList())
	assert.True(t, reflect.DeepEqual(obj.GetCountriesList(), newObj.GetCountriesList()))
	newObj.SetCurrenciesList(obj.GetCurrenciesList())
	assert.True(t, reflect.DeepEqual(obj.GetCurrenciesList(), newObj.GetCurrenciesList()))
	newObj.SetTemplatesList(obj.GetTemplatesList())
	assert.True(t, reflect.DeepEqual(obj.GetTemplatesList(), newObj.GetTemplatesList()))
	newObj.SetDnTemplatesList(obj.GetDnTemplatesList())
	assert.True(t, reflect.DeepEqual(obj.GetDnTemplatesList(), newObj.GetDnTemplatesList()))
	newObj.SetAiTemplatesList(obj.GetAiTemplatesList())
	assert.True(t, reflect.DeepEqual(obj.GetAiTemplatesList(), newObj.GetAiTemplatesList()))
	newObj.SetPaymentMethodsList(obj.GetPaymentMethodsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentMethodsList(), newObj.GetPaymentMethodsList()))
	newObj.SetPaymentAccountsList(obj.GetPaymentAccountsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountsList(), newObj.GetPaymentAccountsList()))
	newObj.SetVatTypesList(obj.GetVatTypesList())
	assert.True(t, reflect.DeepEqual(obj.GetVatTypesList(), newObj.GetVatTypesList()))
}
