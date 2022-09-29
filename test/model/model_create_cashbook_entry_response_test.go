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

var CreateCashbookEntryResponseJsonStr string = "{\"data\":{\"id\":\"1\",\"date\":\"2021-12-19\",\"description\":\"Fattura n. 201/2021\",\"kind\":\"issued_document\",\"type\":\"in\",\"entity_name\":\"Rossi S.r.l.\",\"document\":{\"id\":12345,\"type\":\"issued_document\",\"path\":\"/doc1.pdf\"},\"amount_in\":10,\"payment_account_in\":{\"id\":21,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false},\"amount_out\":10,\"payment_account_out\":{\"id\":21,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false}}}"

func TestCreateCashbookEntryResponse(t *testing.T) {
	obj := NewCreateCashbookEntryResponse()
	json.Unmarshal([]byte(CreateCashbookEntryResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CreateCashbookEntryResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCreateCashbookEntryResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
