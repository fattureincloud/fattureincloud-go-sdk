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

var F24JsonStr string = "{\"id\":12345,\"due_date\":\"2021-12-31\",\"status\":\"paid\",\"payment_account\":{\"id\":111,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\"},\"amount\":840.36,\"attachment_token\":\"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879\",\"description\":\"PAGAMENTO IVA 2021\"}"

func TestF24(t *testing.T) {
	obj := NewF24()
	json.Unmarshal([]byte(F24JsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, F24JsonStr, string(jsonStr2), "they should be equal")

	newObj := NewF24()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetDueDate(obj.GetDueDate())
	assert.True(t, reflect.DeepEqual(obj.GetDueDate(), newObj.GetDueDate()))
	newObj.SetStatus(obj.GetStatus())
	assert.True(t, reflect.DeepEqual(obj.GetStatus(), newObj.GetStatus()))
	newObj.SetPaymentAccount(obj.GetPaymentAccount())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentAccount(), newObj.GetPaymentAccount()))
	newObj.SetAmount(obj.GetAmount())
	assert.True(t, reflect.DeepEqual(obj.GetAmount(), newObj.GetAmount()))
	newObj.SetAttachmentUrl(obj.GetAttachmentUrl())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentUrl(), newObj.GetAttachmentUrl()))
	newObj.SetAttachmentToken(obj.GetAttachmentToken())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentToken(), newObj.GetAttachmentToken()))
	newObj.SetDescription(obj.GetDescription())
	assert.True(t, reflect.DeepEqual(obj.GetDescription(), newObj.GetDescription()))
}
