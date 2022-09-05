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

var DetailedCountryJsonStr string = "{\"name\":\"Italia\",\"settings_name\":\"Italia\",\"iso\":\"IT\",\"fiscal_iso\":\"IT\",\"uic\":\"086\"}"

func TestDetailedCountry(t *testing.T) {
	obj := NewDetailedCountry()
	json.Unmarshal([]byte(DetailedCountryJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, DetailedCountryJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewDetailedCountry()
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetSettingsName(obj.GetSettingsName())
	assert.True(t, reflect.DeepEqual(obj.GetSettingsName(), newObj.GetSettingsName()))
	newObj.SetIso(obj.GetIso())
	assert.True(t, reflect.DeepEqual(obj.GetIso(), newObj.GetIso()))
	newObj.SetFiscalIso(obj.GetFiscalIso())
	assert.True(t, reflect.DeepEqual(obj.GetFiscalIso(), newObj.GetFiscalIso()))
	newObj.SetUic(obj.GetUic())
	assert.True(t, reflect.DeepEqual(obj.GetUic(), newObj.GetUic()))
}
