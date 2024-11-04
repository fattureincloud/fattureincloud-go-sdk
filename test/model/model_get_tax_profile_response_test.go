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

var GetTaxProfileResponseJsonStr string = "{\"data\":{\"company_type\":\"individual\",\"company_subtype\":\"artigiani\",\"regime\":\"forfettario_5\",\"rivalsa_name\":\"\",\"default_rivalsa\":0,\"cassa_name\":\"\",\"default_cassa\":0,\"default_cassa_taxable\":100,\"cassa2_name\":\"\",\"default_cassa2\":0,\"default_cassa2_taxable\":0,\"default_withholding_tax\":0,\"default_withholding_tax_taxable\":100,\"default_other_withholding_tax\":0,\"enasarco\":false,\"contributions_percentage\":0,\"med\":false,\"default_vat\":{\"id\":66,\"value\":0,\"description\":\"Contribuenti forfettari\",\"notes\":\"Operazione non soggetta a IVA ai sensi dell'art. 1, commi 54-89, Legge n. 190/2014 e succ. modifiche/integrazioni\",\"e_invoice\":true,\"ei_type\":\"2.2\",\"ei_description\":\"Non soggetta art. 1/54-89 L. 190/2014 e succ. modifiche/integrazioni\",\"editable\":false,\"is_disabled\":false,\"default\":true}}}"

func TestGetTaxProfileResponse(t *testing.T) {
	obj := NewGetTaxProfileResponse()
	json.Unmarshal([]byte(GetTaxProfileResponseJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, GetTaxProfileResponseJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewGetTaxProfileResponse()
	newObj.SetData(obj.GetData())
	assert.True(t, reflect.DeepEqual(obj.GetData(), newObj.GetData()))
}
