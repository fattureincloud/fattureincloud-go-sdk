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

var IssuedDocumentJsonStr string = "{\"id\":12345,\"type\":\"invoice\",\"notes\":\"bando\",\"show_totals\":\"all\"}"

func TestIssuedDocument(t *testing.T) {
	obj := NewIssuedDocument()
	json.Unmarshal([]byte(IssuedDocumentJsonStr), obj)
	jsonStr2, _ := json.Marshal(obj)

	assert.JSONEq(t, IssuedDocumentJsonStr, string(jsonStr2), "they should be equal")

	newObj := NewIssuedDocument()
	newObj.SetId(obj.GetId())
	assert.True(t, reflect.DeepEqual(obj.GetId(), newObj.GetId()))
	newObj.SetEntity(obj.GetEntity())
	assert.True(t, reflect.DeepEqual(obj.GetEntity(), newObj.GetEntity()))
	newObj.SetType(obj.GetType())
	assert.True(t, reflect.DeepEqual(obj.GetType(), newObj.GetType()))
	newObj.SetNumber(obj.GetNumber())
	assert.True(t, reflect.DeepEqual(obj.GetNumber(), newObj.GetNumber()))
	newObj.SetNumeration(obj.GetNumeration())
	assert.True(t, reflect.DeepEqual(obj.GetNumeration(), newObj.GetNumeration()))
	newObj.SetDate(obj.GetDate())
	assert.True(t, reflect.DeepEqual(obj.GetDate(), newObj.GetDate()))
	newObj.SetYear(obj.GetYear())
	assert.True(t, reflect.DeepEqual(obj.GetYear(), newObj.GetYear()))
	newObj.SetCurrency(obj.GetCurrency())
	assert.True(t, reflect.DeepEqual(obj.GetCurrency(), newObj.GetCurrency()))
	newObj.SetLanguage(obj.GetLanguage())
	assert.True(t, reflect.DeepEqual(obj.GetLanguage(), newObj.GetLanguage()))
	newObj.SetSubject(obj.GetSubject())
	assert.True(t, reflect.DeepEqual(obj.GetSubject(), newObj.GetSubject()))
	newObj.SetVisibleSubject(obj.GetVisibleSubject())
	assert.True(t, reflect.DeepEqual(obj.GetVisibleSubject(), newObj.GetVisibleSubject()))
	newObj.SetRcCenter(obj.GetRcCenter())
	assert.True(t, reflect.DeepEqual(obj.GetRcCenter(), newObj.GetRcCenter()))
	newObj.SetNotes(obj.GetNotes())
	assert.True(t, reflect.DeepEqual(obj.GetNotes(), newObj.GetNotes()))
	newObj.SetRivalsa(obj.GetRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetRivalsa(), newObj.GetRivalsa()))
	newObj.SetCassa(obj.GetCassa())
	assert.True(t, reflect.DeepEqual(obj.GetCassa(), newObj.GetCassa()))
	newObj.SetAmountCassa(obj.GetAmountCassa())
	assert.True(t, reflect.DeepEqual(obj.GetAmountCassa(), newObj.GetAmountCassa()))
	newObj.SetCassaTaxable(obj.GetCassaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetCassaTaxable(), newObj.GetCassaTaxable()))
	newObj.SetAmountCassaTaxable(obj.GetAmountCassaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountCassaTaxable(), newObj.GetAmountCassaTaxable()))
	newObj.SetCassa2(obj.GetCassa2())
	assert.True(t, reflect.DeepEqual(obj.GetCassa2(), newObj.GetCassa2()))
	newObj.SetAmountCassa2(obj.GetAmountCassa2())
	assert.True(t, reflect.DeepEqual(obj.GetAmountCassa2(), newObj.GetAmountCassa2()))
	newObj.SetCassa2Taxable(obj.GetCassa2Taxable())
	assert.True(t, reflect.DeepEqual(obj.GetCassa2Taxable(), newObj.GetCassa2Taxable()))
	newObj.SetAmountCassa2Taxable(obj.GetAmountCassa2Taxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountCassa2Taxable(), newObj.GetAmountCassa2Taxable()))
	newObj.SetGlobalCassaTaxable(obj.GetGlobalCassaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetGlobalCassaTaxable(), newObj.GetGlobalCassaTaxable()))
	newObj.SetAmountGlobalCassaTaxable(obj.GetAmountGlobalCassaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGlobalCassaTaxable(), newObj.GetAmountGlobalCassaTaxable()))
	newObj.SetWithholdingTax(obj.GetWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetWithholdingTax(), newObj.GetWithholdingTax()))
	newObj.SetWithholdingTaxTaxable(obj.GetWithholdingTaxTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetWithholdingTaxTaxable(), newObj.GetWithholdingTaxTaxable()))
	newObj.SetOtherWithholdingTax(obj.GetOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetOtherWithholdingTax(), newObj.GetOtherWithholdingTax()))
	newObj.SetStampDuty(obj.GetStampDuty())
	assert.True(t, reflect.DeepEqual(obj.GetStampDuty(), newObj.GetStampDuty()))
	newObj.SetPaymentMethod(obj.GetPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentMethod(), newObj.GetPaymentMethod()))
	newObj.SetUseSplitPayment(obj.GetUseSplitPayment())
	assert.True(t, reflect.DeepEqual(obj.GetUseSplitPayment(), newObj.GetUseSplitPayment()))
	newObj.SetUseGrossPrices(obj.GetUseGrossPrices())
	assert.True(t, reflect.DeepEqual(obj.GetUseGrossPrices(), newObj.GetUseGrossPrices()))
	newObj.SetEInvoice(obj.GetEInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetEInvoice(), newObj.GetEInvoice()))
	newObj.SetEiData(obj.GetEiData())
	assert.True(t, reflect.DeepEqual(obj.GetEiData(), newObj.GetEiData()))
	newObj.SetEiCassaType(obj.GetEiCassaType())
	assert.True(t, reflect.DeepEqual(obj.GetEiCassaType(), newObj.GetEiCassaType()))
	newObj.SetEiCassa2Type(obj.GetEiCassa2Type())
	assert.True(t, reflect.DeepEqual(obj.GetEiCassa2Type(), newObj.GetEiCassa2Type()))
	newObj.SetEiWithholdingTaxCausal(obj.GetEiWithholdingTaxCausal())
	assert.True(t, reflect.DeepEqual(obj.GetEiWithholdingTaxCausal(), newObj.GetEiWithholdingTaxCausal()))
	newObj.SetEiOtherWithholdingTaxType(obj.GetEiOtherWithholdingTaxType())
	assert.True(t, reflect.DeepEqual(obj.GetEiOtherWithholdingTaxType(), newObj.GetEiOtherWithholdingTaxType()))
	newObj.SetEiOtherWithholdingTaxCausal(obj.GetEiOtherWithholdingTaxCausal())
	assert.True(t, reflect.DeepEqual(obj.GetEiOtherWithholdingTaxCausal(), newObj.GetEiOtherWithholdingTaxCausal()))
	newObj.SetItemsList(obj.GetItemsList())
	assert.True(t, reflect.DeepEqual(obj.GetItemsList(), newObj.GetItemsList()))
	newObj.SetPaymentsList(obj.GetPaymentsList())
	assert.True(t, reflect.DeepEqual(obj.GetPaymentsList(), newObj.GetPaymentsList()))
	newObj.SetTemplate(obj.GetTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetTemplate(), newObj.GetTemplate()))
	newObj.SetDeliveryNoteTemplate(obj.GetDeliveryNoteTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetDeliveryNoteTemplate(), newObj.GetDeliveryNoteTemplate()))
	newObj.SetAccInvTemplate(obj.GetAccInvTemplate())
	assert.True(t, reflect.DeepEqual(obj.GetAccInvTemplate(), newObj.GetAccInvTemplate()))
	newObj.SetHMargins(obj.GetHMargins())
	assert.True(t, reflect.DeepEqual(obj.GetHMargins(), newObj.GetHMargins()))
	newObj.SetVMargins(obj.GetVMargins())
	assert.True(t, reflect.DeepEqual(obj.GetVMargins(), newObj.GetVMargins()))
	newObj.SetShowPayments(obj.GetShowPayments())
	assert.True(t, reflect.DeepEqual(obj.GetShowPayments(), newObj.GetShowPayments()))
	newObj.SetShowPaymentMethod(obj.GetShowPaymentMethod())
	assert.True(t, reflect.DeepEqual(obj.GetShowPaymentMethod(), newObj.GetShowPaymentMethod()))
	newObj.SetShowTotals(obj.GetShowTotals())
	assert.True(t, reflect.DeepEqual(obj.GetShowTotals(), newObj.GetShowTotals()))
	newObj.SetShowNotificationButton(obj.GetShowNotificationButton())
	assert.True(t, reflect.DeepEqual(obj.GetShowNotificationButton(), newObj.GetShowNotificationButton()))
	newObj.SetShowTspayButton(obj.GetShowTspayButton())
	assert.True(t, reflect.DeepEqual(obj.GetShowTspayButton(), newObj.GetShowTspayButton()))
	newObj.SetDeliveryNote(obj.GetDeliveryNote())
	assert.True(t, reflect.DeepEqual(obj.GetDeliveryNote(), newObj.GetDeliveryNote()))
	newObj.SetAccompanyingInvoice(obj.GetAccompanyingInvoice())
	assert.True(t, reflect.DeepEqual(obj.GetAccompanyingInvoice(), newObj.GetAccompanyingInvoice()))
	newObj.SetDnNumber(obj.GetDnNumber())
	assert.True(t, reflect.DeepEqual(obj.GetDnNumber(), newObj.GetDnNumber()))
	newObj.SetDnDate(obj.GetDnDate())
	assert.True(t, reflect.DeepEqual(obj.GetDnDate(), newObj.GetDnDate()))
	newObj.SetDnAiPackagesNumber(obj.GetDnAiPackagesNumber())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiPackagesNumber(), newObj.GetDnAiPackagesNumber()))
	newObj.SetDnAiWeight(obj.GetDnAiWeight())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiWeight(), newObj.GetDnAiWeight()))
	newObj.SetDnAiCausal(obj.GetDnAiCausal())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiCausal(), newObj.GetDnAiCausal()))
	newObj.SetDnAiDestination(obj.GetDnAiDestination())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiDestination(), newObj.GetDnAiDestination()))
	newObj.SetDnAiTransporter(obj.GetDnAiTransporter())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiTransporter(), newObj.GetDnAiTransporter()))
	newObj.SetDnAiNotes(obj.GetDnAiNotes())
	assert.True(t, reflect.DeepEqual(obj.GetDnAiNotes(), newObj.GetDnAiNotes()))
	newObj.SetIsMarked(obj.GetIsMarked())
	assert.True(t, reflect.DeepEqual(obj.GetIsMarked(), newObj.GetIsMarked()))
	newObj.SetAmountNet(obj.GetAmountNet())
	assert.True(t, reflect.DeepEqual(obj.GetAmountNet(), newObj.GetAmountNet()))
	newObj.SetAmountVat(obj.GetAmountVat())
	assert.True(t, reflect.DeepEqual(obj.GetAmountVat(), newObj.GetAmountVat()))
	newObj.SetAmountGross(obj.GetAmountGross())
	assert.True(t, reflect.DeepEqual(obj.GetAmountGross(), newObj.GetAmountGross()))
	newObj.SetAmountDueDiscount(obj.GetAmountDueDiscount())
	assert.True(t, reflect.DeepEqual(obj.GetAmountDueDiscount(), newObj.GetAmountDueDiscount()))
	newObj.SetAmountRivalsa(obj.GetAmountRivalsa())
	assert.True(t, reflect.DeepEqual(obj.GetAmountRivalsa(), newObj.GetAmountRivalsa()))
	newObj.SetAmountRivalsaTaxable(obj.GetAmountRivalsaTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountRivalsaTaxable(), newObj.GetAmountRivalsaTaxable()))
	newObj.SetAmountWithholdingTax(obj.GetAmountWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountWithholdingTax(), newObj.GetAmountWithholdingTax()))
	newObj.SetAmountWithholdingTaxTaxable(obj.GetAmountWithholdingTaxTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountWithholdingTaxTaxable(), newObj.GetAmountWithholdingTaxTaxable()))
	newObj.SetAmountOtherWithholdingTax(obj.GetAmountOtherWithholdingTax())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOtherWithholdingTax(), newObj.GetAmountOtherWithholdingTax()))
	newObj.SetAmountOtherWithholdingTaxTaxable(obj.GetAmountOtherWithholdingTaxTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountOtherWithholdingTaxTaxable(), newObj.GetAmountOtherWithholdingTaxTaxable()))
	newObj.SetAmountEnasarcoTaxable(obj.GetAmountEnasarcoTaxable())
	assert.True(t, reflect.DeepEqual(obj.GetAmountEnasarcoTaxable(), newObj.GetAmountEnasarcoTaxable()))
	newObj.SetExtraData(obj.GetExtraData())
	assert.True(t, reflect.DeepEqual(obj.GetExtraData(), newObj.GetExtraData()))
	newObj.SetSeenDate(obj.GetSeenDate())
	assert.True(t, reflect.DeepEqual(obj.GetSeenDate(), newObj.GetSeenDate()))
	newObj.SetNextDueDate(obj.GetNextDueDate())
	assert.True(t, reflect.DeepEqual(obj.GetNextDueDate(), newObj.GetNextDueDate()))
	newObj.SetUrl(obj.GetUrl())
	assert.True(t, reflect.DeepEqual(obj.GetUrl(), newObj.GetUrl()))
	newObj.SetAttachmentUrl(obj.GetAttachmentUrl())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentUrl(), newObj.GetAttachmentUrl()))
	newObj.SetAttachmentToken(obj.GetAttachmentToken())
	assert.True(t, reflect.DeepEqual(obj.GetAttachmentToken(), newObj.GetAttachmentToken()))
	newObj.SetEiRaw(obj.GetEiRaw())
	assert.True(t, reflect.DeepEqual(obj.GetEiRaw(), newObj.GetEiRaw()))
	newObj.SetEiStatus(obj.GetEiStatus())
	assert.True(t, reflect.DeepEqual(obj.GetEiStatus(), newObj.GetEiStatus()))
}
