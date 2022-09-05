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

var CreateProductRequestJsonStr string = "{\"data\":{\"id\":12345,\"name\":\"neim\",\"code\":\"cod\",\"net_price\":10,\"gross_price\":10,\"use_gross_price\":true,\"default_vat\":{\"id\":1},\"net_cost\":10,\"measure\":\"big\",\"description\":\"desc\",\"category\":\"cat6\",\"notes\":\"nots\",\"in_stock\":true,\"stock_initial\":10,\"average_cost\":10,\"average_price\":10,\"created_at\":\"2021-10-10\",\"updated_at\":\"2021-10-10\"}}"

func TestCreateProductRequest(t *testing.T) {
	obj := NewCreateProductRequest()
	json.Unmarshal([]byte(CreateProductRequestJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CreateProductRequestJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCreateProductRequest()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
