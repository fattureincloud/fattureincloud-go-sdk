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

var EInvoiceRejectionReasonJsonStr string = "{\"reason\":\"invalid date\",\"ei_status\":\"rejected\",\"solution\":\"set a valid date\",\"code\":\"c01\",\"date\":\"2022-10-10T23:22:21.000000001+02:00\"}"

func TestEInvoiceRejectionReason(t *testing.T) {
	obj := NewEInvoiceRejectionReason()
	json.Unmarshal([]byte(EInvoiceRejectionReasonJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, EInvoiceRejectionReasonJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewEInvoiceRejectionReason()
	newObj.SetReason(obj.GetReason())
	assert.True(t, reflect.DeepEqual(obj.GetReason(), newObj.GetReason()))
	newObj.SetEiStatus(obj.GetEiStatus())
	assert.True(t, reflect.DeepEqual(obj.GetEiStatus(), newObj.GetEiStatus()))
	newObj.SetSolution(obj.GetSolution())
	assert.True(t, reflect.DeepEqual(obj.GetSolution(), newObj.GetSolution()))
	newObj.SetCode(obj.GetCode())
	assert.True(t, reflect.DeepEqual(obj.GetCode(), newObj.GetCode()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
}
