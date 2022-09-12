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

var CreateF24ResponseJsonStr string = "{\"data\":{\"id\":12345,\"due_date\":\"2021-12-31\",\"status\":\"paid\",\"payment_account\":{\"id\":111,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\"},\"amount\":840.36,\"attachment_token\":\"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879\",\"description\":\"PAGAMENTO IVA 2021\"}}"

func TestCreateF24Response(t *testing.T) {
	obj := NewCreateF24Response()
	json.Unmarshal([]byte(CreateF24ResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CreateF24ResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCreateF24Response()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}