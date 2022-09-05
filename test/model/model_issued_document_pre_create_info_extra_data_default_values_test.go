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

var IssuedDocumentPreCreateInfoExtraDataDefaultValuesJsonStr string = "{\"ts_communication\":true,\"ts_tipo_spesa\":\"ts\",\"ts_flag_tipo_spesa\":1,\"ts_pagamento_tracciato\":true}"

func TestIssuedDocumentPreCreateInfoExtraDataDefaultValues(t *testing.T) {
	obj := NewIssuedDocumentPreCreateInfoExtraDataDefaultValues()
	json.Unmarshal([]byte(IssuedDocumentPreCreateInfoExtraDataDefaultValuesJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentPreCreateInfoExtraDataDefaultValuesJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentPreCreateInfoExtraDataDefaultValues()
	newObj.SetTsCommunication(obj.GetTsCommunication())
	assert.True(t, reflect.DeepEqual(obj.GetTsCommunication(), newObj.GetTsCommunication()))
	newObj.SetTsTipoSpesa(obj.GetTsTipoSpesa())
	assert.True(t, reflect.DeepEqual(obj.GetTsTipoSpesa(), newObj.GetTsTipoSpesa()))
	newObj.SetTsFlagTipoSpesa(obj.GetTsFlagTipoSpesa())
	assert.True(t, reflect.DeepEqual(obj.GetTsFlagTipoSpesa(), newObj.GetTsFlagTipoSpesa()))
	newObj.SetTsPagamentoTracciato(obj.GetTsPagamentoTracciato())
	assert.True(t, reflect.DeepEqual(obj.GetTsPagamentoTracciato(), newObj.GetTsPagamentoTracciato()))
}
