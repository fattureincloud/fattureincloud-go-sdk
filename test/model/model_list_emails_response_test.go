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

var ListEmailsResponseJsonStr string = "{\"current_page\":1,\"data\":[{\"id\":1,\"status\":\"sent\",\"sent_date\":\"2022-07-17T13:53:12Z\",\"errors_count\":0,\"error_log\":\"\",\"from_email\":\"test@mail.it\",\"from_name\":\"Test mail\",\"to_email\":\"mail@test.it\",\"to_name\":\"Mario\",\"subject\":\"Test\",\"content\":\"Test send email\",\"copy_to\":\"\",\"recipient_status\":\"unknown\",\"recipient_date\":\"2022-07-18T13:53:12Z\",\"kind\":\"Fatture\",\"attachments\":[]},{\"id\":2,\"status\":\"sent\",\"sent_date\":\"2022-07-18T13:53:12Z\",\"errors_count\":0,\"error_log\":\"\",\"from_email\":\"test@mail.it\",\"from_name\":\"Test mail\",\"to_email\":\"mail@test.it\",\"to_name\":\"Maria\",\"subject\":\"Test\",\"content\":\"Test send email\",\"copy_to\":\"\",\"recipient_status\":\"unknown\",\"recipient_date\":\"2022-07-18T13:53:12Z\",\"kind\":\"Fatture\",\"attachments\":[]}],\"first_page_url\":\"emails?page=1\",\"next_page_url\":\"emails?page=1\",\"from\":1,\"last_page\":1,\"last_page_url\":\"emails?page=1\",\"path\":\"emails\",\"per_page\":50,\"prev_page_url\":\"emails?page=1\",\"to\":2,\"total\":2}"

func TestListEmailsResponse(t *testing.T) {
	obj := NewListEmailsResponse()
	json.Unmarshal([]byte(ListEmailsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListEmailsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListEmailsResponse()
	newObj.SetCurrentPage(obj.GetCurrentPage())
	assert.True(t, reflect.DeepEqual(obj.GetCurrentPage(), newObj.GetCurrentPage()))
	newObj.SetFirstPageUrl(obj.GetFirstPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetFirstPageUrl(), newObj.GetFirstPageUrl()))
	newObj.SetFrom(obj.GetFrom())
	assert.True(t, reflect.DeepEqual(obj.GetFrom(), newObj.GetFrom()))
	newObj.SetLastPage(obj.GetLastPage())
	assert.True(t, reflect.DeepEqual(obj.GetLastPage(), newObj.GetLastPage()))
	newObj.SetLastPageUrl(obj.GetLastPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetLastPageUrl(), newObj.GetLastPageUrl()))
	newObj.SetNextPageUrl(obj.GetNextPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetNextPageUrl(), newObj.GetNextPageUrl()))
	newObj.SetPath(obj.GetPath())
	assert.True(t, reflect.DeepEqual(obj.GetPath(), newObj.GetPath()))
	newObj.SetPerPage(obj.GetPerPage())
	assert.True(t, reflect.DeepEqual(obj.GetPerPage(), newObj.GetPerPage()))
	newObj.SetPrevPageUrl(obj.GetPrevPageUrl())
	assert.True(t, reflect.DeepEqual(obj.GetPrevPageUrl(), newObj.GetPrevPageUrl()))
	newObj.SetTo(obj.GetTo())
	assert.True(t, reflect.DeepEqual(obj.GetTo(), newObj.GetTo()))
	newObj.SetTotal(obj.GetTotal())
	assert.True(t, reflect.DeepEqual(obj.GetTotal(), newObj.GetTotal()))
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
