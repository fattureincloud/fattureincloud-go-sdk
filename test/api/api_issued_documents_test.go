/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateIssuedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"type":"receipt","year":2021,"numeration":"rec123","stamp_duty":0,"use_gross_prices":false,"e_invoice":false,"delivery_note":false,"accompanying_invoice":false,"amount_net":68.18,"amount_vat":6.82,"amount_gross":75,"amount_due_discount":0,"amount_rivalsa":0,"amount_cassa":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"h_margins":15,"v_margins":16,"show_payment_method":false,"show_payments":true,"show_totals":"all","show_notification_button":false,"is_marked":false,"entity":{"id":54321,"name":"Mary Red S.r.L","vat_number":"IT05432181211","tax_code":"IT05432181211","address_street":"Corso impero, 66","address_postal_code":"20900","address_city":"Milano","address_province":"MI","country":"Italia","certified_email":"mary@pec.red.com","ei_code":"ABCXCR1"},"date":"2021-08-20","number":1,"currency":{"id":"EUR","exchange_rate":"1.00000","symbol":"€"},"language":{"code":"it","name":"Italiano"},"rivalsa":0,"cassa":0,"withholding_tax":0,"withholding_tax_taxable":100,"other_withholding_tax":0,"payment_method":{"id":4,"name":"Credit card"},"use_split_payment":false,"items_list":[{"product_id":5432,"code":"SG3","name":"Soggiorno","net_price":68.18182,"id":277876033,"gross_price":75,"apply_withholding_taxes":true,"qty":1,"vat":{"id":3,"value":10},"stock":false,"not_taxable":false}],"payments_list":[{"amount":75,"due_date":"2020-08-23","id":69078013,"status":"not_paid"}],"attachment_url":"kdijrnf893hnwkfk45f50f.pdf","next_due_date":"2020-08-23","template":{"id":2821,"name":"Light Smoke"},"url":"y12h45rn9yf2mse0p43t7ec90vr.pdf","show_tspay_button":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.CreateIssuedDocument(context.Background(), 2).CreateIssuedDocumentRequest(*NewCreateIssuedDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocument().
		SetId(12345).
		SetEntity(*NewEntity().
			SetId(54321).
			SetName("Mary Red S.r.L").
			SetVatNumber("IT05432181211").
			SetTaxCode("IT05432181211").
			SetAddressStreet("Corso impero, 66").
			SetAddressPostalCode("20900").
			SetAddressCity("Milano").
			SetAddressProvince("MI").
			SetCountry("Italia").
			SetCertifiedEmail("mary@pec.red.com").
			SetEiCode("ABCXCR1")).
		SetType(IssuedDocumentTypes.RECEIPT).
		SetNumber(1).
		SetNumeration("rec123").
		SetDate("2021-08-20").
		SetYear(2021).
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetSymbol("€").
			SetExchangeRate("1.00000")).
		SetLanguage(*NewLanguage().
			SetCode("it").
			SetName("Italiano")).
		SetRivalsa(0).
		SetCassa(0).
		SetAmountCassa(0).
		SetWithholdingTax(0).
		SetWithholdingTaxTaxable(100).
		SetOtherWithholdingTax(0).
		SetStampDuty(0).
		SetPaymentMethod(*NewPaymentMethod().
			SetId(4).
			SetName("Credit card")).
		SetUseSplitPayment(false).
		SetUseGrossPrices(false).
		SetEInvoice(false).
		SetItemsList([]IssuedDocumentItemsListItem{
			*NewIssuedDocumentItemsListItem().
				SetProductId(5432).
				SetCode("SG3").
				SetName("Soggiorno").
				SetNetPrice(68.18182).
				SetId(277876033).
				SetGrossPrice(75).
				SetApplyWithholdingTaxes(true).
				SetQty(1).
				SetVat(*NewVatType().
					SetId(3).
					SetValue(10)).
				SetStock(false).
				SetNotTaxable(false)}).
		SetPaymentsList([]IssuedDocumentPaymentsListItem{
			*NewIssuedDocumentPaymentsListItem().
				SetAmount(75).
				SetDueDate("2020-08-23").
				SetId(69078013).
				SetStatus(IssuedDocumentStatuses.NOT_PAID)}).
		SetTemplate(*NewDocumentTemplate().
			SetId(2821).
			SetName("Light Smoke")).
		SetHMargins(15).
		SetVMargins(16).
		SetShowPayments(true).
		SetShowPaymentMethod(false).
		SetShowTotals(ShowTotalsModes.ALL).
		SetShowNotificationButton(false).
		SetShowTspayButton(false).
		SetDeliveryNote(false).
		SetAccompanyingInvoice(false).
		SetIsMarked(false).
		SetAmountNet(68.18).
		SetAmountVat(6.82).
		SetAmountGross(75).
		SetAmountDueDiscount(0).
		SetAmountRivalsa(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetNextDueDate("2020-08-23").
		SetUrl("y12h45rn9yf2mse0p43t7ec90vr.pdf").
		SetAttachmentUrl("kdijrnf893hnwkfk45f50f.pdf")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteIssuedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	resp, err := apiClient.IssuedDocumentsAPI.DeleteIssuedDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteIssuedDocumentAttachment(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	resp, err := apiClient.IssuedDocumentsAPI.DeleteIssuedDocumentAttachment(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetEmailData(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"recipient_email":"mary.red@example.com","default_sender_email":{"id":0,"email":"no-reply@fattureincloud.it"},"sender_emails_list":[{"id":0,"email":"no-reply@fattureincloud.it"},{"id":888,"email":"mariorossi@fattureincloud.it"}],"cc_email":"m.rossi@exxample.com","subject":"Nostra pro forma nr. 1","body":"Gentile Mario Rossi,<br>per vedere la nostra pro forma di  o per scaricarne una copia in versione PDF prema sul bottone sottoastante.<br><br>{{allegati}}<br><br>Cordiali saluti,<br><b>Mario Rossi<\\/b>","document_exists":true,"delivery_note_exists":false,"attachment_exists":false,"accompanying_invoice_exists":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.GetEmailData(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewEmailData().
		SetRecipientEmail("mary.red@example.com").
		SetDefaultSenderEmail(*NewEmailDataDefaultSenderEmail().
			SetId(0).
			SetEmail("no-reply@fattureincloud.it")).
		SetSenderEmailsList([]SenderEmail{
			*NewSenderEmail().
				SetId(0).
				SetEmail("no-reply@fattureincloud.it"),
			*NewSenderEmail().
				SetId(888).
				SetEmail("mariorossi@fattureincloud.it"),
		}).
		SetCcEmail("m.rossi@exxample.com").
		SetSubject("Nostra pro forma nr. 1").
		SetBody("Gentile Mario Rossi,<br>per vedere la nostra pro forma di  o per scaricarne una copia in versione PDF prema sul bottone sottoastante.<br><br>{{allegati}}<br><br>Cordiali saluti,<br><b>Mario Rossi<\\/b>").
		SetDocumentExists(true).
		SetDeliveryNoteExists(false).
		SetAttachmentExists(false).
		SetAccompanyingInvoiceExists(false)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetExistingIssuedDocumentTotals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"amount_net":10,"amount_vat":10,"amount_gross":10,"amount_withholding_tax":10,"amount_other_withholding_tax":10,"amount_due":10,"payments_sum":10}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.GetExistingIssuedDocumentTotals(context.Background(), 2, 2).GetExistingIssuedDocumentTotalsRequest(*NewGetExistingIssuedDocumentTotalsRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocumentTotals().
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetAmountWithholdingTax(10).
		SetAmountOtherWithholdingTax(10).
		SetAmountDue(10).
		SetPaymentsSum(10)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetIssuedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"type":"receipt","year":2021,"numeration":"rec123","stamp_duty":0,"use_gross_prices":false,"e_invoice":false,"delivery_note":false,"accompanying_invoice":false,"amount_net":68.18,"amount_vat":6.82,"amount_gross":75,"amount_due_discount":0,"amount_rivalsa":0,"amount_cassa":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"h_margins":15,"v_margins":16,"show_payment_method":false,"show_payments":true,"show_totals":"all","show_notification_button":false,"is_marked":false,"entity":{"id":54321,"name":"Mary Red S.r.L","vat_number":"IT05432181211","tax_code":"IT05432181211","address_street":"Corso impero, 66","address_postal_code":"20900","address_city":"Milano","address_province":"MI","country":"Italia","certified_email":"mary@pec.red.com","ei_code":"ABCXCR1"},"date":"2021-08-20","number":1,"currency":{"id":"EUR","exchange_rate":"1.00000","symbol":"€"},"language":{"code":"it","name":"Italiano"},"rivalsa":0,"cassa":0,"withholding_tax":0,"withholding_tax_taxable":100,"other_withholding_tax":0,"payment_method":{"id":4,"name":"Credit card"},"use_split_payment":false,"items_list":[{"product_id":5432,"code":"SG3","name":"Soggiorno","net_price":68.18182,"id":277876033,"gross_price":75,"apply_withholding_taxes":true,"qty":1,"vat":{"id":3,"value":10},"stock":false,"not_taxable":false}],"payments_list":[{"amount":75,"due_date":"2020-08-23","id":69078013,"status":"not_paid"}],"attachment_url":"kdijrnf893hnwkfk45f50f.pdf","next_due_date":"2020-08-23","template":{"id":2821,"name":"Light Smoke"},"url":"y12h45rn9yf2mse0p43t7ec90vr.pdf","show_tspay_button":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.GetIssuedDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocument().
		SetId(12345).
		SetEntity(*NewEntity().
			SetId(54321).
			SetName("Mary Red S.r.L").
			SetVatNumber("IT05432181211").
			SetTaxCode("IT05432181211").
			SetAddressStreet("Corso impero, 66").
			SetAddressPostalCode("20900").
			SetAddressCity("Milano").
			SetAddressProvince("MI").
			SetCountry("Italia").
			SetCertifiedEmail("mary@pec.red.com").
			SetEiCode("ABCXCR1")).
		SetType(IssuedDocumentTypes.RECEIPT).
		SetNumber(1).
		SetNumeration("rec123").
		SetDate("2021-08-20").
		SetYear(2021).
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetSymbol("€").
			SetExchangeRate("1.00000")).
		SetLanguage(*NewLanguage().
			SetCode("it").
			SetName("Italiano")).
		SetRivalsa(0).
		SetCassa(0).
		SetAmountCassa(0).
		SetWithholdingTax(0).
		SetWithholdingTaxTaxable(100).
		SetOtherWithholdingTax(0).
		SetStampDuty(0).
		SetPaymentMethod(*NewPaymentMethod().
			SetId(4).
			SetName("Credit card")).
		SetUseSplitPayment(false).
		SetUseGrossPrices(false).
		SetEInvoice(false).
		SetItemsList([]IssuedDocumentItemsListItem{
			*NewIssuedDocumentItemsListItem().
				SetProductId(5432).
				SetCode("SG3").
				SetName("Soggiorno").
				SetNetPrice(68.18182).
				SetId(277876033).
				SetGrossPrice(75).
				SetApplyWithholdingTaxes(true).
				SetQty(1).
				SetVat(*NewVatType().
					SetId(3).
					SetValue(10)).
				SetStock(false).
				SetNotTaxable(false)}).
		SetPaymentsList([]IssuedDocumentPaymentsListItem{
			*NewIssuedDocumentPaymentsListItem().
				SetAmount(75).
				SetDueDate("2020-08-23").
				SetId(69078013).
				SetStatus(IssuedDocumentStatuses.NOT_PAID)}).
		SetTemplate(*NewDocumentTemplate().
			SetId(2821).
			SetName("Light Smoke")).
		SetHMargins(15).
		SetVMargins(16).
		SetShowPayments(true).
		SetShowPaymentMethod(false).
		SetShowTotals(ShowTotalsModes.ALL).
		SetShowNotificationButton(false).
		SetShowTspayButton(false).
		SetDeliveryNote(false).
		SetAccompanyingInvoice(false).
		SetIsMarked(false).
		SetAmountNet(68.18).
		SetAmountVat(6.82).
		SetAmountGross(75).
		SetAmountDueDiscount(0).
		SetAmountRivalsa(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetNextDueDate("2020-08-23").
		SetUrl("y12h45rn9yf2mse0p43t7ec90vr.pdf").
		SetAttachmentUrl("kdijrnf893hnwkfk45f50f.pdf")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetIssuedDocumentPreCreateInfo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"default_values":{"default_template":{"id":1},"dn_template":{"id":1},"ai_template":{"id":1},"notes":"notes","rivalsa":1,"cassa":1,"withholding_tax":1,"withholding_tax_taxable":1,"other_withholding_tax":1,"use_gross_prices":true,"payment_method":{"id":1}},"extra_data_default_values":{"ts_communication":true,"ts_tipo_spesa":"ts","ts_flag_tipo_spesa":1,"ts_pagamento_tracciato":true},"items_default_values":{"vat":{"id":1}},"countries_list":["Italia","Marocco"],"currencies_list":[{"id":"EUR"},{"id":"DNR"}],"templates_list":[{"id":1},{"id":2}],"dn_templates_list":[{"id":1},{"id":2}],"ai_templates_list":[{"id":1},{"id":2}],"payment_methods_list":[{"id":1},{"id":2}],"payment_accounts_list":[{"id":1},{"id":2}],"vat_types_list":[{"id":1},{"id":2}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.GetIssuedDocumentPreCreateInfo(context.Background(), 2).Type_("invoice").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocumentPreCreateInfo().
		SetNumerations(nil).
		SetDnNumerations(nil).
		SetDefaultValues(*NewIssuedDocumentPreCreateInfoDefaultValues().
			SetDefaultTemplate(*NewDocumentTemplate().SetId(1)).
			SetDnTemplate(*NewDocumentTemplate().SetId(1)).
			SetAiTemplate(*NewDocumentTemplate().SetId(1)).
			SetNotes("notes").
			SetRivalsa(1).
			SetCassa(1).
			SetWithholdingTax(1).
			SetWithholdingTaxTaxable(1).
			SetOtherWithholdingTax(1).
			SetUseGrossPrices(true).
			SetPaymentMethod(*NewPaymentMethod().SetId(1))).
		SetExtraDataDefaultValues(*NewIssuedDocumentPreCreateInfoExtraDataDefaultValues().
			SetTsCommunication(true).
			SetTsTipoSpesa("ts").
			SetTsFlagTipoSpesa(1).
			SetTsPagamentoTracciato(true)).
		SetItemsDefaultValues(*NewIssuedDocumentPreCreateInfoItemsDefaultValues().
			SetVat(*NewVatType().SetId(1))).
		SetCountriesList([]string{"Italia", "Marocco"}).
		SetCurrenciesList([]Currency{*NewCurrency().SetId("EUR"), *NewCurrency().SetId("DNR")}).
		SetTemplatesList([]DocumentTemplate{*NewDocumentTemplate().SetId(1), *NewDocumentTemplate().SetId(2)}).
		SetDnTemplatesList([]DocumentTemplate{*NewDocumentTemplate().SetId(1), *NewDocumentTemplate().SetId(2)}).
		SetAiTemplatesList([]DocumentTemplate{*NewDocumentTemplate().SetId(1), *NewDocumentTemplate().SetId(2)}).
		SetPaymentMethodsList([]PaymentMethod{*NewPaymentMethod().SetId(1), *NewPaymentMethod().SetId(2)}).
		SetPaymentAccountsList([]PaymentAccount{*NewPaymentAccount().SetId(1), *NewPaymentAccount().SetId(2)}).
		SetVatTypesList([]VatType{*NewVatType().SetId(1), *NewVatType().SetId(2)})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestGetNewIssuedDocumentTotals(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"amount_net":10,"amount_vat":10,"amount_gross":10,"amount_withholding_tax":10,"amount_other_withholding_tax":10,"amount_due":10,"payments_sum":10}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.GetNewIssuedDocumentTotals(context.Background(), 2).GetNewIssuedDocumentTotalsRequest(*NewGetNewIssuedDocumentTotalsRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocumentTotals().
		SetAmountNet(10).
		SetAmountVat(10).
		SetAmountGross(10).
		SetAmountWithholdingTax(10).
		SetAmountOtherWithholdingTax(10).
		SetAmountDue(10).
		SetPaymentsSum(10)

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListIssuedDocuments(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12345,"type":"receipt","year":2021,"numeration":"rec123","stamp_duty":0,"use_gross_prices":false,"e_invoice":false,"delivery_note":false,"accompanying_invoice":false,"amount_net":68.18,"amount_vat":6.82,"amount_gross":75,"amount_due_discount":0,"amount_rivalsa":0,"amount_cassa":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"h_margins":15,"v_margins":16,"show_payment_method":false,"show_payments":true,"show_totals":"all","show_notification_button":false,"is_marked":false,"entity":{"id":54321,"name":"Mary Red S.r.L","vat_number":"IT05432181211","tax_code":"IT05432181211","address_street":"Corso impero, 66","address_postal_code":"20900","address_city":"Milano","address_province":"MI","country":"Italia","certified_email":"mary@pec.red.com","ei_code":"ABCXCR1"},"date":"2021-08-20","number":1,"currency":{"id":"EUR","exchange_rate":"1.00000","symbol":"€"},"language":{"code":"it","name":"Italiano"},"rivalsa":0,"cassa":0,"withholding_tax":0,"withholding_tax_taxable":100,"other_withholding_tax":0,"payment_method":{"id":4,"name":"Credit card"},"use_split_payment":false,"items_list":[{"product_id":5432,"code":"SG3","name":"Soggiorno","net_price":68.18182,"id":277876033,"gross_price":75,"apply_withholding_taxes":true,"qty":1,"vat":{"id":3,"value":10},"stock":false,"not_taxable":false}],"payments_list":[{"amount":75,"due_date":"2020-08-23","id":69078013,"status":"not_paid"}],"attachment_url":"kdijrnf893hnwkfk45f50f.pdf","next_due_date":"2020-08-23","template":{"id":2821,"name":"Light Smoke"},"url":"y12h45rn9yf2mse0p43t7ec90vr.pdf","show_tspay_button":false}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.ListIssuedDocuments(context.Background(), 2).Type_("invoice").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocument().
		SetId(12345).
		SetEntity(*NewEntity().
			SetId(54321).
			SetName("Mary Red S.r.L").
			SetVatNumber("IT05432181211").
			SetTaxCode("IT05432181211").
			SetAddressStreet("Corso impero, 66").
			SetAddressPostalCode("20900").
			SetAddressCity("Milano").
			SetAddressProvince("MI").
			SetCountry("Italia").
			SetCertifiedEmail("mary@pec.red.com").
			SetEiCode("ABCXCR1")).
		SetType(IssuedDocumentTypes.RECEIPT).
		SetNumber(1).
		SetNumeration("rec123").
		SetDate("2021-08-20").
		SetYear(2021).
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetSymbol("€").
			SetExchangeRate("1.00000")).
		SetLanguage(*NewLanguage().
			SetCode("it").
			SetName("Italiano")).
		SetRivalsa(0).
		SetCassa(0).
		SetAmountCassa(0).
		SetWithholdingTax(0).
		SetWithholdingTaxTaxable(100).
		SetOtherWithholdingTax(0).
		SetStampDuty(0).
		SetPaymentMethod(*NewPaymentMethod().
			SetId(4).
			SetName("Credit card")).
		SetUseSplitPayment(false).
		SetUseGrossPrices(false).
		SetEInvoice(false).
		SetItemsList([]IssuedDocumentItemsListItem{
			*NewIssuedDocumentItemsListItem().
				SetProductId(5432).
				SetCode("SG3").
				SetName("Soggiorno").
				SetNetPrice(68.18182).
				SetId(277876033).
				SetGrossPrice(75).
				SetApplyWithholdingTaxes(true).
				SetQty(1).
				SetVat(*NewVatType().
					SetId(3).
					SetValue(10)).
				SetStock(false).
				SetNotTaxable(false)}).
		SetPaymentsList([]IssuedDocumentPaymentsListItem{
			*NewIssuedDocumentPaymentsListItem().
				SetAmount(75).
				SetDueDate("2020-08-23").
				SetId(69078013).
				SetStatus(IssuedDocumentStatuses.NOT_PAID)}).
		SetTemplate(*NewDocumentTemplate().
			SetId(2821).
			SetName("Light Smoke")).
		SetHMargins(15).
		SetVMargins(16).
		SetShowPayments(true).
		SetShowPaymentMethod(false).
		SetShowTotals(ShowTotalsModes.ALL).
		SetShowNotificationButton(false).
		SetShowTspayButton(false).
		SetDeliveryNote(false).
		SetAccompanyingInvoice(false).
		SetIsMarked(false).
		SetAmountNet(68.18).
		SetAmountVat(6.82).
		SetAmountGross(75).
		SetAmountDueDiscount(0).
		SetAmountRivalsa(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetNextDueDate("2020-08-23").
		SetUrl("y12h45rn9yf2mse0p43t7ec90vr.pdf").
		SetAttachmentUrl("kdijrnf893hnwkfk45f50f.pdf")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyIssuedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"type":"receipt","year":2021,"numeration":"rec123","stamp_duty":0,"use_gross_prices":false,"e_invoice":false,"delivery_note":false,"accompanying_invoice":false,"amount_net":68.18,"amount_vat":6.82,"amount_gross":75,"amount_due_discount":0,"amount_rivalsa":0,"amount_cassa":0,"amount_withholding_tax":0,"amount_other_withholding_tax":0,"h_margins":15,"v_margins":16,"show_payment_method":false,"show_payments":true,"show_totals":"all","show_notification_button":false,"is_marked":false,"entity":{"id":54321,"name":"Mary Red S.r.L","vat_number":"IT05432181211","tax_code":"IT05432181211","address_street":"Corso impero, 66","address_postal_code":"20900","address_city":"Milano","address_province":"MI","country":"Italia","certified_email":"mary@pec.red.com","ei_code":"ABCXCR1"},"date":"2021-08-20","number":1,"currency":{"id":"EUR","exchange_rate":"1.00000","symbol":"€"},"language":{"code":"it","name":"Italiano"},"rivalsa":0,"cassa":0,"withholding_tax":0,"withholding_tax_taxable":100,"other_withholding_tax":0,"payment_method":{"id":4,"name":"Credit card"},"use_split_payment":false,"items_list":[{"product_id":5432,"code":"SG3","name":"Soggiorno","net_price":68.18182,"id":277876033,"gross_price":75,"apply_withholding_taxes":true,"qty":1,"vat":{"id":3,"value":10},"stock":false,"not_taxable":false}],"payments_list":[{"amount":75,"due_date":"2020-08-23","id":69078013,"status":"not_paid"}],"attachment_url":"kdijrnf893hnwkfk45f50f.pdf","next_due_date":"2020-08-23","template":{"id":2821,"name":"Light Smoke"},"url":"y12h45rn9yf2mse0p43t7ec90vr.pdf","show_tspay_button":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.ModifyIssuedDocument(context.Background(), 2, 2).ModifyIssuedDocumentRequest(*NewModifyIssuedDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewIssuedDocument().
		SetId(12345).
		SetEntity(*NewEntity().
			SetId(54321).
			SetName("Mary Red S.r.L").
			SetVatNumber("IT05432181211").
			SetTaxCode("IT05432181211").
			SetAddressStreet("Corso impero, 66").
			SetAddressPostalCode("20900").
			SetAddressCity("Milano").
			SetAddressProvince("MI").
			SetCountry("Italia").
			SetCertifiedEmail("mary@pec.red.com").
			SetEiCode("ABCXCR1")).
		SetType(IssuedDocumentTypes.RECEIPT).
		SetNumber(1).
		SetNumeration("rec123").
		SetDate("2021-08-20").
		SetYear(2021).
		SetCurrency(*NewCurrency().
			SetId("EUR").
			SetSymbol("€").
			SetExchangeRate("1.00000")).
		SetLanguage(*NewLanguage().
			SetCode("it").
			SetName("Italiano")).
		SetRivalsa(0).
		SetCassa(0).
		SetAmountCassa(0).
		SetWithholdingTax(0).
		SetWithholdingTaxTaxable(100).
		SetOtherWithholdingTax(0).
		SetStampDuty(0).
		SetPaymentMethod(*NewPaymentMethod().
			SetId(4).
			SetName("Credit card")).
		SetUseSplitPayment(false).
		SetUseGrossPrices(false).
		SetEInvoice(false).
		SetItemsList([]IssuedDocumentItemsListItem{
			*NewIssuedDocumentItemsListItem().
				SetProductId(5432).
				SetCode("SG3").
				SetName("Soggiorno").
				SetNetPrice(68.18182).
				SetId(277876033).
				SetGrossPrice(75).
				SetApplyWithholdingTaxes(true).
				SetQty(1).
				SetVat(*NewVatType().
					SetId(3).
					SetValue(10)).
				SetStock(false).
				SetNotTaxable(false)}).
		SetPaymentsList([]IssuedDocumentPaymentsListItem{
			*NewIssuedDocumentPaymentsListItem().
				SetAmount(75).
				SetDueDate("2020-08-23").
				SetId(69078013).
				SetStatus(IssuedDocumentStatuses.NOT_PAID)}).
		SetTemplate(*NewDocumentTemplate().
			SetId(2821).
			SetName("Light Smoke")).
		SetHMargins(15).
		SetVMargins(16).
		SetShowPayments(true).
		SetShowPaymentMethod(false).
		SetShowTotals(ShowTotalsModes.ALL).
		SetShowNotificationButton(false).
		SetShowTspayButton(false).
		SetDeliveryNote(false).
		SetAccompanyingInvoice(false).
		SetIsMarked(false).
		SetAmountNet(68.18).
		SetAmountVat(6.82).
		SetAmountGross(75).
		SetAmountDueDiscount(0).
		SetAmountRivalsa(0).
		SetAmountWithholdingTax(0).
		SetAmountOtherWithholdingTax(0).
		SetNextDueDate("2020-08-23").
		SetUrl("y12h45rn9yf2mse0p43t7ec90vr.pdf").
		SetAttachmentUrl("kdijrnf893hnwkfk45f50f.pdf")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestScheduleEmail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	resp, err := apiClient.IssuedDocumentsAPI.ScheduleEmail(context.Background(), 2, 2).ScheduleEmailRequest(*NewScheduleEmailRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUploadIssuedDocumentAttachment(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(``))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	_, resp, err := apiClient.IssuedDocumentsAPI.UploadIssuedDocumentAttachment(context.Background(), 2).Filename("filename").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransformIssuedDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"type":"invoice"},"options":{"create_from":["82112399"],"transform":true,"keep_copy":true}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.TransformIssuedDocument(context.Background(), 2).OriginalDocumentId(12345).NewType("proforma").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewTransformIssuedDocumentResponse().
		SetData(
			*NewIssuedDocument().
				SetId(12345).
				SetType(IssuedDocumentTypes.INVOICE),
		).
		SetOptions(
			*NewIssuedDocumentOptions().
				SetCreateFrom([]string{"82112399"}).
				SetKeepCopy(true).
				SetTransform(true),
		)

	assert.True(t, reflect.DeepEqual(expected, actual))
}

func TestJoinIssuedDocuments(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"type":"invoice"},"options":{"create_from":["82112399","82112400"]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.IssuedDocumentsAPI.JoinIssuedDocuments(context.Background(), 2).Ids("2332,876").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewJoinIssuedDocumentsResponse().
		SetData(
			*NewIssuedDocument().
				SetId(12345).
				SetType(IssuedDocumentTypes.INVOICE),
		).
		SetOptions(
			*NewIssuedDocumentOptions().
				SetCreateFrom([]string{"82112399", "82112400"}),
		)

	exj, _ := json.Marshal(expected)
	fmt.Print(string(exj))

	assert.True(t, reflect.DeepEqual(expected, actual))
}
