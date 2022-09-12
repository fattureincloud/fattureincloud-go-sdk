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

var GetExistingReceivedDocumentTotalsResponseJsonStr string = "{\"data\":{\"amount_net\":10,\"amount_vat\":10,\"amount_gross\":10,\"amount_withholding_tax\":10,\"amount_other_withholding_tax\":10,\"amount_due\":10,\"payments_sum\":10}}"

func TestGetExistingReceivedDocumentTotalsResponse(t *testing.T) {
	obj := NewGetExistingReceivedDocumentTotalsResponse()
	json.Unmarshal([]byte(GetExistingReceivedDocumentTotalsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetExistingReceivedDocumentTotalsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetExistingReceivedDocumentTotalsResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}