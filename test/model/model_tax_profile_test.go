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

var TaxProfileJsonStr string = "{\"company_type\":\"individual\",\"company_subtype\":\"artigiani\",\"regime\":\"forfettario_5\",\"rivalsa_name\":\"\",\"default_rivalsa\":0,\"cassa_name\":\"\",\"default_cassa\":0,\"default_cassa_taxable\":100,\"cassa2_name\":\"\",\"default_cassa2\":0,\"default_cassa2_taxable\":0,\"default_withholding_tax\":0,\"default_withholding_tax_taxable\":100,\"default_other_withholding_tax\":0,\"enasarco\":false,\"contributions_percentage\":0,\"med\":false,\"default_vat\":{\"id\":66,\"value\":0,\"description\":\"Contribuenti forfettari\",\"notes\":\"Operazione non soggetta a IVA ai sensi dell'art. 1, commi 54-89, Legge n. 190/2014 e succ. modifiche/integrazioni\",\"e_invoice\":true,\"ei_type\":\"2.2\",\"ei_description\":\"Non soggetta art. 1/54-89 L. 190/2014 e succ. modifiche/integrazioni\",\"editable\":false,\"is_disabled\":false,\"default\":true}}"

func TestTaxProfile(t *testing.T) {
	obj := NewTaxProfile()
	json.Unmarshal([]byte(TaxProfileJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, TaxProfileJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewTaxProfile()
	newObj.SetCompanyType(obj.GetCompanyType())
	assert.True(t, reflect.DeepEqual(obj.GetCompanyType(), newObj.GetCompanyType()))
	newObj.SetCompanySubtype(obj.GetCompanySubtype())
	assert.True(t, reflect.DeepEqual(obj.GetCompanySubtype(), newObj.GetCompanySubtype()))
	newObj.SetRegime(obj.GetRegime())
	assert.True(t, reflect.DeepEqual(obj.GetRegime(), newObj.GetRegime()))
	newObj.SetRivalsaName(obj.GetRivalsaName())
	assert.True(t, reflect.DeepEqual(obj.GetRivalsaName(), newObj.GetRivalsaName()))
	newObj.SetDefaultRivalsa(obj.GetDefaultRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultRivalsa(), newObj.GetDefaultRivalsa()))
	newObj.SetCassaName(obj.GetCassaName())
	assert.True(t, reflect.DeepEqual(obj.GetCassaName(), newObj.GetCassaName()))
	newObj.SetDefaultCassa(obj.GetDefaultCassa())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultCassa(), newObj.GetDefaultCassa()))
	newObj.SetDefaultCassaTaxable(obj.GetDefaultCassaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultCassaTaxable(), newObj.GetDefaultCassaTaxable()))
	newObj.SetCassa2Name(obj.GetCassa2Name())
	assert.True(t, reflect.DeepEqual(obj.GetCassa2Name(), newObj.GetCassa2Name()))
	newObj.SetDefaultCassa2(obj.GetDefaultCassa2())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultCassa2(), newObj.GetDefaultCassa2()))
	newObj.SetDefaultCassa2Taxable(obj.GetDefaultCassa2Taxable())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultCassa2Taxable(), newObj.GetDefaultCassa2Taxable()))
	newObj.SetDefaultWithholdingTax(obj.GetDefaultWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultWithholdingTax(), newObj.GetDefaultWithholdingTax()))
	newObj.SetDefaultWithholdingTaxTaxable(obj.GetDefaultWithholdingTaxTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultWithholdingTaxTaxable(), newObj.GetDefaultWithholdingTaxTaxable()))
	newObj.SetDefaultOtherWithholdingTax(obj.GetDefaultOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultOtherWithholdingTax(), newObj.GetDefaultOtherWithholdingTax()))
	newObj.SetEnasarco(obj.GetEnasarco())
	assert.True(t, reflect.DeepEqual(obj.GetEnasarco(), newObj.GetEnasarco()))
	newObj.SetContributionsPercentage(obj.GetContributionsPercentage())
	assert.True(t, reflect.DeepEqual(obj.GetContributionsPercentage(), newObj.GetContributionsPercentage()))
	newObj.SetMed(obj.GetMed())
	assert.True(t, reflect.DeepEqual(obj.GetMed(), newObj.GetMed()))
	newObj.SetDefaultVat(obj.GetDefaultVat())
	assert.True(t, reflect.DeepEqual(obj.GetDefaultVat(), newObj.GetDefaultVat()))
}
