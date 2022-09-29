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

var ListF24ResponseJsonStr string = "{\"current_page\":10,\"first_page_url\":\"https://www.page.url/\",\"from\":10,\"last_page\":10,\"last_page_url\":\"https://www.page.url/\",\"next_page_url\":\"https://www.page.url/\",\"path\":\"https://www.page.url/\",\"per_page\":10,\"prev_page_url\":\"https://www.page.url/\",\"to\":10,\"total\":10,\"data\":[{\"id\":12345,\"due_date\":\"2021-12-31\",\"status\":\"paid\",\"payment_account\":{\"id\":111,\"name\":\"Indesa - Carta conto\",\"type\":\"standard\"},\"amount\":840.36,\"attachment_token\":\"Adfqregwthwrt6whrtghsrgbsdthyeruerur6u6676e5879\",\"description\":\"PAGAMENTO IVA 2021\"}]}"

func TestListF24Response(t *testing.T) {
	obj := NewListF24Response()
	json.Unmarshal([]byte(ListF24ResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListF24ResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListF24Response()
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
	newObj.SetAggregatedData(obj.GetAggregatedData())
	assert.True(t, reflect.DeepEqual(obj.GetAggregatedData(), newObj.GetAggregatedData()))
}
