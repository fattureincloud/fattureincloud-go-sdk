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

var IssuedDocumentExtraDataJsonStr string = "{\"multifatture_sent\":5,\"ts_communication\":true,\"ts_flag_tipo_spesa\":2,\"ts_pagamento_tracciato\":true,\"ts_tipo_spesa\":\"TK\",\"ts_opposizione\":true,\"ts_status\":5,\"ts_file_id\":\"str\",\"ts_sent_date\":\"2021-12-25\",\"ts_full_amount\":true,\"imported_by\":\"me\"}"

func TestIssuedDocumentExtraData(t *testing.T) {
	obj := NewIssuedDocumentExtraData()
	json.Unmarshal([]byte(IssuedDocumentExtraDataJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentExtraDataJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocumentExtraData()
	newObj.SetMultifattureSent(obj.GetMultifattureSent())
	assert.True(t, reflect.DeepEqual(obj.GetMultifattureSent(), newObj.GetMultifattureSent()))
	newObj.SetTsCommunication(obj.GetTsCommunication())
	assert.True(t, reflect.DeepEqual(obj.GetTsCommunication(), newObj.GetTsCommunication()))
	newObj.SetTsFlagTipoSpesa(obj.GetTsFlagTipoSpesa())
	assert.True(t, reflect.DeepEqual(obj.GetTsFlagTipoSpesa(), newObj.GetTsFlagTipoSpesa()))
	newObj.SetTsPagamentoTracciato(obj.GetTsPagamentoTracciato())
	assert.True(t, reflect.DeepEqual(obj.GetTsPagamentoTracciato(), newObj.GetTsPagamentoTracciato()))
	newObj.SetTsTipoSpesa(obj.GetTsTipoSpesa())
	assert.True(t, reflect.DeepEqual(obj.GetTsTipoSpesa(), newObj.GetTsTipoSpesa()))
	newObj.SetTsOpposizione(obj.GetTsOpposizione())
	assert.True(t, reflect.DeepEqual(obj.GetTsOpposizione(), newObj.GetTsOpposizione()))
	newObj.SetTsStatus(obj.GetTsStatus())
	assert.True(t, reflect.DeepEqual(obj.GetTsStatus(), newObj.GetTsStatus()))
	newObj.SetTsFileId(obj.GetTsFileId())
	assert.True(t, reflect.DeepEqual(obj.GetTsFileId(), newObj.GetTsFileId()))
	newObj.SetTsSentDate(obj.GetTsSentDate())
	assert.True(t, reflect.DeepEqual(obj.GetTsSentDate(), newObj.GetTsSentDate()))
	newObj.SetTsFullAmount(obj.GetTsFullAmount())
	assert.True(t, reflect.DeepEqual(obj.GetTsFullAmount(), newObj.GetTsFullAmount()))
	newObj.SetImportedBy(obj.GetImportedBy())
	assert.True(t, reflect.DeepEqual(obj.GetImportedBy(), newObj.GetImportedBy()))
}
