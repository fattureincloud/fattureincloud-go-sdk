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

var CashbookEntryJsonStr string = "{\"id\":\"1\",\"date\":\"2021-12-19\",\"description\":\"Fattura n. 201/2021\",\"kind\":\"issued_document\",\"type\":\"in\",\"entity_name\":\"Rossi S.r.l.\",\"document\":{\"id\":12345,\"type\":\"issued_document\",\"path\":\"/doc1.pdf\"},\"amount_in\":10,\"payment_account_in\":{\"id\":21,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false},\"amount_out\":10,\"payment_account_out\":{\"id\":21,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\",\"iban\":\"IT84Y0300203280294126225888\",\"sia\":\"sai\",\"cuc\":\"cuc\",\"virtual\":false}}"

func TestCashbookEntry(t *testing.T) {
	obj := NewCashbookEntry()
	json.Unmarshal([]byte(CashbookEntryJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CashbookEntryJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCashbookEntry()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
	newObj.SetKind(obj.GetKind())
	assert.True(t, reflect.DeepEqual(obj.GetKind(), newObj.GetKind()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetEntityName(obj.GetEntityName())
	assert.True(t, reflect.DeepEqual(obj.GetEntityName(), newObj.GetEntityName()))
	newObj.SetDocument(obj.GetDocument())
	assert.True(t, reflect.DeepEqual(obj.GetDocument(), newObj.GetDocument()))
	newObj.SetAmountIn(obj.GetAmountIn())
	assert.True(t, reflect.DeepEqual(obj.GetAmountIn(), newObj.GetAmountIn()))
	newObj.SetPaymentAccountIn(obj.GetPaymentAccountIn())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountIn(), newObj.GetPaymentAccountIn()))
	newObj.SetAmountOut(obj.GetAmountOut())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOut(), newObj.GetAmountOut()))
	newObj.SetPaymentAccountOut(obj.GetPaymentAccountOut())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccountOut(), newObj.GetPaymentAccountOut()))
}
