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

var ListF24ResponseAggregationJsonStr string = "{\"aggregated_data\":{\"amount\":10}}"

func TestListF24ResponseAggregation(t *testing.T) {
	obj := NewListF24ResponseAggregation()
	json.Unmarshal([]byte(ListF24ResponseAggregationJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, ListF24ResponseAggregationJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewListF24ResponseAggregation()
	newObj.SetAggregatedData(obj.GetAggregatedData())
	assert.True(t, reflect.DeepEqual(obj.GetAggregatedData(), newObj.GetAggregatedData()))
}
