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

var ListProductsResponseJsonStr string = "{\"current_page\":10,\"first_page_url\":\"https://www.page.url/\",\"from\":10,\"last_page\":10,\"last_page_url\":\"https://www.page.url/\",\"next_page_url\":\"https://www.page.url/\",\"path\":\"https://www.page.url/\",\"per_page\":10,\"prev_page_url\":\"https://www.page.url/\",\"to\":10,\"total\":10,\"data\":[{\"id\":12345,\"name\":\"neim\",\"code\":\"cod\",\"net_price\":10,\"gross_price\":10,\"use_gross_price\":true,\"default_vat\":{\"id\":1},\"net_cost\":10,\"measure\":\"big\",\"description\":\"desc\",\"category\":\"cat6\",\"notes\":\"nots\",\"in_stock\":true,\"stock_initial\":10,\"average_cost\":10,\"average_price\":10,\"created_at\":\"2021-10-10\",\"updated_at\":\"2021-10-10\"}]}"

func TestListProductsResponse(t *testing.T) {
	obj := NewListProductsResponse()
	json.Unmarshal([]byte(ListProductsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListProductsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListProductsResponse()
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
