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

var EmailScheduleJsonStr string = "{\"sender_id\":5,\"sender_email\":\"mariorossi@fattureincloud.it\",\"recipient_email\":\"mary.red@example.com\",\"subject\":\"Nostra pro forma nr. 1\",\"body\":\"Pro forma body\",\"include\":{\"document\":false,\"delivery_note\":false,\"attachment\":false,\"accompanying_invoice\":false},\"attach_pdf\":true,\"send_copy\":false}"

func TestEmailSchedule(t *testing.T) {
	obj := NewEmailSchedule()
	json.Unmarshal([]byte(EmailScheduleJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EmailScheduleJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEmailSchedule()
	newObj.SetSenderId(obj.GetSenderId())
	assert.True(t, reflect.DeepEqual(obj.GetSenderId(), newObj.GetSenderId()))
	newObj.SetSenderEmail(obj.GetSenderEmail())
	assert.True(t, reflect.DeepEqual(obj.GetSenderEmail(), newObj.GetSenderEmail()))
	newObj.SetRecipientEmail(obj.GetRecipientEmail())
	assert.True(t, reflect.DeepEqual(obj.GetRecipientEmail(), newObj.GetRecipientEmail()))
	newObj.SetSubject(obj.GetSubject())
	assert.True(t, reflect.DeepEqual(obj.GetSubject(), newObj.GetSubject()))
	newObj.SetBody(obj.GetBody())
	assert.True(t, reflect.DeepEqual(obj.GetBody(), newObj.GetBody()))
	newObj.SetInclude(obj.GetInclude())
	assert.True(t, reflect.DeepEqual(obj.GetInclude(), newObj.GetInclude()))
	newObj.SetAttachPdf(obj.GetAttachPdf())
	assert.True(t, reflect.DeepEqual(obj.GetAttachPdf(), newObj.GetAttachPdf()))
	newObj.SetSendCopy(obj.GetSendCopy())
	assert.True(t, reflect.DeepEqual(obj.GetSendCopy(), newObj.GetSendCopy()))
}
