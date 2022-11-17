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

var ListEmailsResponsePageJsonStr string = "{\"data\":[{\"id\":1,\"status\":\"sent\",\"sent_date\":\"2022-07-17T13:53:12Z\",\"errors_count\":0,\"error_log\":\"\",\"from_email\":\"test@mail.it\",\"from_name\":\"Test mail\",\"to_email\":\"mail@test.it\",\"to_name\":\"Mario\",\"subject\":\"Test\",\"content\":\"Test send email\",\"copy_to\":\"\",\"recipient_status\":\"unknown\",\"recipient_date\":\"2022-07-18T13:53:12Z\",\"kind\":\"Fatture\",\"attachments\":[]},{\"id\":2,\"status\":\"sent\",\"sent_date\":\"2022-07-18T13:53:12Z\",\"errors_count\":0,\"error_log\":\"\",\"from_email\":\"test@mail.it\",\"from_name\":\"Test mail\",\"to_email\":\"mail@test.it\",\"to_name\":\"Maria\",\"subject\":\"Test\",\"content\":\"Test send email\",\"copy_to\":\"\",\"recipient_status\":\"unknown\",\"recipient_date\":\"2022-07-18T13:53:12Z\",\"kind\":\"Fatture\",\"attachments\":[]}]}"

func TestListEmailsResponsePage(t *testing.T) {
	obj := NewListEmailsResponsePage()
	json.Unmarshal([]byte(ListEmailsResponsePageJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListEmailsResponsePageJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListEmailsResponsePage()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
