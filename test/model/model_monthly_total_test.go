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

var MonthlyTotalJsonStr string = "{\"net\":10,\"gross\":10,\"count\":10}"

func TestMonthlyTotal(t *testing.T) {
	obj := NewMonthlyTotal()
	json.Unmarshal([]byte(MonthlyTotalJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, MonthlyTotalJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewMonthlyTotal()
	newObj.SetNet(obj.GetNet())
	assert.True(t, reflect.DeepEqual(obj.GetNet(), newObj.GetNet()))
	newObj.SetGross(obj.GetGross())
	assert.True(t, reflect.DeepEqual(obj.GetGross(), newObj.GetGross()))
	newObj.SetCount(obj.GetCount())
	assert.True(t, reflect.DeepEqual(obj.GetCount(), newObj.GetCount()))
}
