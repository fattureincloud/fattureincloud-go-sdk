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

var ListF24ResponseAggregatedDataJsonStr string = "{\"amount\":10}"

func TestListF24ResponseAggregatedData(t *testing.T) {
	obj := NewListF24ResponseAggregatedData()
	json.Unmarshal([]byte(ListF24ResponseAggregatedDataJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListF24ResponseAggregatedDataJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListF24ResponseAggregatedData()
	newObj.SetAmount(obj.GetAmount())
	assert.True(t, reflect.DeepEqual(obj.GetAmount(), newObj.GetAmount()))
}
