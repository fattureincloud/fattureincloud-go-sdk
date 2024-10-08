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

var CompanyJsonStr string = "{\"id\":5,\"name\":\"Azienda 1\",\"type\":\"company\",\"access_token\":\"ABCDEF\",\"controlled_companies\":[],\"connection_id\":94817,\"tax_code\":\"SLVMTT50A01F205L\",\"vat_number\":\"IT1235657345A\",\"fic_plan\":\"standard\",\"fic_license_expire\":\"2024-10-10\"}"

func TestCompany(t *testing.T) {
	obj := NewCompany()
	json.Unmarshal([]byte(CompanyJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, CompanyJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewCompany()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetName(obj.GetName())
	assert.True(t, reflect.DeepEqual(obj.GetName(), newObj.GetName()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetAccessToken(obj.GetAccessToken())
	assert.True(t, reflect.DeepEqual(obj.GetAccessToken(), newObj.GetAccessToken()))
	newObj.SetControlledCompanies(obj.GetControlledCompanies())
	assert.True(t, reflect.DeepEqual(obj.GetControlledCompanies(), newObj.GetControlledCompanies()))
	newObj.SetConnectionId(obj.GetConnectionId())
	assert.True(t, reflect.DeepEqual(obj.GetConnectionId(), newObj.GetConnectionId()))
	newObj.SetTaxCode(obj.GetTaxCode())
	assert.True(t, reflect.DeepEqual(obj.GetTaxCode(), newObj.GetTaxCode()))
	newObj.SetVatNumber(obj.GetVatNumber())
	assert.True(t, reflect.DeepEqual(obj.GetVatNumber(), newObj.GetVatNumber()))
	newObj.SetFicPlan(obj.GetFicPlan())
	assert.True(t, reflect.DeepEqual(obj.GetFicPlan(), newObj.GetFicPlan()))
	newObj.SetFicLicenseExpire(obj.GetFicLicenseExpire())
	assert.True(t, reflect.DeepEqual(obj.GetFicLicenseExpire(), newObj.GetFicLicenseExpire()))
}
