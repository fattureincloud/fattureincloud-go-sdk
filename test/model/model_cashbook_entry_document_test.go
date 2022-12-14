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

var CashbookEntryDocumentJsonStr string = "{\"id\":12345,\"type\":\"issued_document\",\"path\":\"/doc1.pdf\"}"

func TestCashbookEntryDocument(t *testing.T) {
	obj := NewCashbookEntryDocument()
	json.Unmarshal([]byte(CashbookEntryDocumentJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CashbookEntryDocumentJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCashbookEntryDocument()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetPath(obj.GetPath())
	assert.True(t, reflect.DeepEqual(obj.GetPath(), newObj.GetPath()))
}
