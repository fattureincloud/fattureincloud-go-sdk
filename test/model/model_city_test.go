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

var CityJsonStr string = "{\"postal_code\":\"24016\",\"city\":\"San Pellegrino Terme\",\"province\":\"BG\"}"

func TestCity(t *testing.T) {
	obj := NewCity()
	json.Unmarshal([]byte(CityJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CityJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCity()
	newObj.SetPostalCode(obj.GetPostalCode())
	assert.True(t, reflect.DeepEqual(obj.GetPostalCode(), newObj.GetPostalCode()))
	newObj.SetCity(obj.GetCity())
	assert.True(t, reflect.DeepEqual(obj.GetCity(), newObj.GetCity()))
	newObj.SetProvince(obj.GetProvince())
	assert.True(t, reflect.DeepEqual(obj.GetProvince(), newObj.GetProvince()))
}
