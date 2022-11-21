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

var EmailJsonStr string = "{\"id\":1,\"status\":\"sent\",\"sent_date\":\"2022-07-17T13:53:12Z\",\"errors_count\":0,\"error_log\":\"\",\"from_email\":\"test@mail.it\",\"from_name\":\"Test mail\",\"to_email\":\"mail@test.it\",\"to_name\":\"Mario\",\"subject\":\"Test\",\"content\":\"Test send email\",\"copy_to\":\"\",\"recipient_status\":\"unknown\",\"recipient_date\":\"2022-07-17T13:53:12Z\",\"kind\":\"Fatture\",\"attachments\":[]}"

func TestEmail(t *testing.T) {
	obj := NewEmail()
	json.Unmarshal([]byte(EmailJsonStr), obj)

	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EmailJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEmail()
	newObj.SetAttachments(obj.GetAttachments())
	assert.True(t, reflect.DeepEqual(obj.GetAttachments(), newObj.GetAttachments()))
	newObj.SetContent(obj.GetContent())
	assert.True(t, reflect.DeepEqual(obj.GetCopyTo(), newObj.GetCopyTo()))
	newObj.SetCopyTo(obj.GetCopyTo())
	assert.True(t, reflect.DeepEqual(obj.GetCopyTo(), newObj.GetCopyTo()))
	newObj.SetErrorLog(obj.GetErrorLog())
	assert.True(t, reflect.DeepEqual(obj.GetErrorLog(), newObj.GetErrorLog()))
	newObj.SetFromEmail(obj.GetFromEmail())
	assert.True(t, reflect.DeepEqual(obj.GetFromEmail(), newObj.GetFromEmail()))
	newObj.SetFromName(obj.GetFromName())
	assert.True(t, reflect.DeepEqual(obj.GetFromName(), newObj.GetFromName()))
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetKind(obj.GetKind())
	assert.True(t, reflect.DeepEqual(obj.GetKind(), newObj.GetKind()))
	newObj.SetRecipientDate(obj.GetRecipientDate())
	assert.True(t, reflect.DeepEqual(obj.GetRecipientDate(), newObj.GetRecipientDate()))
	newObj.SetRecipientStatus(obj.GetRecipientStatus())
	assert.True(t, reflect.DeepEqual(obj.GetRecipientStatus(), newObj.GetRecipientStatus()))
	newObj.SetSentDate(obj.GetSentDate())
	assert.True(t, reflect.DeepEqual(obj.GetSentDate(), newObj.GetSentDate()))
	newObj.SetStatus(obj.GetStatus())
	assert.True(t, reflect.DeepEqual(obj.GetStatus(), newObj.GetStatus()))
	newObj.SetSubject(obj.GetSubject())
	assert.True(t, reflect.DeepEqual(obj.GetSubject(), newObj.GetSubject()))
	newObj.SetToEmail(obj.GetToEmail())
	assert.True(t, reflect.DeepEqual(obj.GetToEmail(), newObj.GetToEmail()))
	newObj.SetToName(obj.GetToName())
	assert.True(t, reflect.DeepEqual(obj.GetToName(), newObj.GetToName()))
}
