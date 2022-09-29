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

var GetIssuedDocumentPreCreateInfoResponseJsonStr string = "{\"data\":{\"default_values\":{\"default_template\":{\"id\":1},\"dn_template\":{\"id\":1},\"ai_template\":{\"id\":1},\"notes\":\"notes\",\"rivalsa\":1,\"cassa\":1,\"withholding_tax\":1,\"withholding_tax_taxable\":1,\"other_withholding_tax\":1,\"use_gross_prices\":true,\"payment_method\":{\"id\":1,\"type\":\"standard\"}},\"extra_data_default_values\":{\"ts_communication\":true,\"ts_tipo_spesa\":\"ts\",\"ts_flag_tipo_spesa\":1,\"ts_pagamento_tracciato\":true},\"items_default_values\":{\"vat\":{\"id\":1}},\"countries_list\":[\"Italia\",\"Marocco\"],\"currencies_list\":[{\"id\":\"EUR\"},{\"id\":\"DNR\"}],\"templates_list\":[{\"id\":1},{\"id\":2}],\"dn_templates_list\":[{\"id\":1},{\"id\":2}],\"ai_templates_list\":[{\"id\":1},{\"id\":2}],\"payment_methods_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"payment_accounts_list\":[{\"id\":1,\"type\":\"standard\"},{\"id\":2,\"type\":\"standard\"}],\"vat_types_list\":[{\"id\":1},{\"id\":2}]}}"

func TestGetIssuedDocumentPreCreateInfoResponse(t *testing.T) {
	obj := NewGetIssuedDocumentPreCreateInfoResponse()
	json.Unmarshal([]byte(GetIssuedDocumentPreCreateInfoResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetIssuedDocumentPreCreateInfoResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetIssuedDocumentPreCreateInfoResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
