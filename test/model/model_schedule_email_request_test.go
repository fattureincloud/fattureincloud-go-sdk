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

var ScheduleEmailRequestJsonStr string = "{\"data\":{\"sender_id\":12345,\"sender_email\":\"mariorossi@fattureincloud.it\",\"recipient_email\":\"mary.red@example.com\",\"subject\":\"Nostra pro forma nr. 1\",\"body\":\"Pro forma body\",\"include\":{\"document\":false,\"delivery_note\":false,\"attachment\":false,\"accompanying_invoice\":false},\"attach_pdf\":true,\"send_copy\":false}}"

func TestScheduleEmailRequest(t *testing.T) {
	obj := NewScheduleEmailRequest()
	json.Unmarshal([]byte(ScheduleEmailRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ScheduleEmailRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewScheduleEmailRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
