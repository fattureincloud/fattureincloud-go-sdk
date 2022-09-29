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

var GetNewIssuedDocumentTotalsResponseJsonStr string = "{\"data\":{\"amount_net\":10,\"amount_rivalsa\":10,\"amount_net_with_rivalsa\":10,\"amount_cassa\":10,\"taxable_amount\":10,\"not_taxable_amount\":10,\"amount_vat\":10,\"amount_gross\":10,\"taxable_amount_withholding_tax\":10,\"amount_withholding_tax\":10,\"taxable_amount_other_withholding_tax\":10,\"amount_other_withholding_tax\":10,\"stamp_duty\":10,\"amount_due\":10,\"is_enasarco_maximal_exceeded\":true,\"payments_sum\":10}}"

func TestGetNewIssuedDocumentTotalsResponse(t *testing.T) {
	obj := NewGetNewIssuedDocumentTotalsResponse()
	json.Unmarshal([]byte(GetNewIssuedDocumentTotalsResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetNewIssuedDocumentTotalsResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetNewIssuedDocumentTotalsResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
