/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":16451,"code":"AE86","name":"Avv. Maria Rossi","type":"person","first_name":"Maria","last_name":"Rossi","contact_person":"","vat_number":"IT12345640962","tax_code":"BLTGNI5ABCDA794E","address_street":"Via Roma, 1","address_postal_code":"20900","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"maria.rossi@example.com","certified_email":"maria.rossi@pec.example.com","phone":"1234567890","fax":"","notes":"","created_at":"2021-04-29 08:53:07","updated_at":"2021-04-29 08:53:07","default_payment_terms":1,"default_payment_terms_type":"standard","bank_name":"Indesa","bank_iban":"IT40P123456781000000123456","bank_swift_code":"AK86PCT","shipping_address":"Corso Magellano 4","e_invoice":true,"ei_code":"111111","default_vat":{"id":54321,"value":45,"description":"","is_disabled":false},"default_payment_method":{"id":386092,"name":"Credit card"}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ClientsApi.CreateClient(context.Background(), 2).CreateClientRequest(*NewCreateClientRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewClient().
		SetId(16451).
		SetType(ClientTypes.PERSON).
		SetCountry("Italia").
		SetDefaultVat(*NewVatType().
			SetId(54321).
			SetValue(45).
			SetDescription("").
			SetIsDisabled(false)).
		SetDefaultPaymentTermsType(PaymentTermsTypes.STANDARD).
		SetDefaultPaymentMethod(*NewPaymentMethod().
			SetId(386092).
			SetName("Credit card")).
		SetEInvoice(true).
		SetCode("AE86").
		SetName("Avv. Maria Rossi").
		SetFirstName("Maria").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("IT12345640962").
		SetTaxCode("BLTGNI5ABCDA794E").
		SetAddressStreet("Via Roma, 1").
		SetAddressPostalCode("20900").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetEmail("maria.rossi@example.com").
		SetCertifiedEmail("maria.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("").
		SetNotes("").
		SetDefaultPaymentTerms(1).
		SetBankName("Indesa").
		SetBankIban("IT40P123456781000000123456").
		SetBankSwiftCode("AK86PCT").
		SetShippingAddress("Corso Magellano 4").
		SetEiCode("111111").
		SetCreatedAt("2021-04-29 08:53:07").
		SetUpdatedAt("2021-04-29 08:53:07")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteClient(t *testing.T) {
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

	resp, err := apiClient.ClientsApi.DeleteClient(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetClient(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":16451,"code":"AE86","name":"Avv. Maria Rossi","type":"person","first_name":"Maria","last_name":"Rossi","contact_person":"","vat_number":"IT12345640962","tax_code":"BLTGNI5ABCDA794E","address_street":"Via Roma, 1","address_postal_code":"20900","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"maria.rossi@example.com","certified_email":"maria.rossi@pec.example.com","phone":"1234567890","fax":"","notes":"","created_at":"2021-04-29 08:53:07","updated_at":"2021-04-29 08:53:07","default_payment_terms":1,"default_payment_terms_type":"standard","bank_name":"Indesa","bank_iban":"IT40P123456781000000123456","bank_swift_code":"AK86PCT","shipping_address":"Corso Magellano 4","e_invoice":true,"ei_code":"111111","default_vat":{"id":54321,"value":45,"description":"","is_disabled":false},"default_payment_method":{"id":386092,"name":"Credit card"}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ClientsApi.GetClient(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewClient().
		SetId(16451).
		SetType(ClientTypes.PERSON).
		SetCountry("Italia").
		SetDefaultVat(*NewVatType().
			SetId(54321).
			SetValue(45).
			SetDescription("").
			SetIsDisabled(false)).
		SetDefaultPaymentTermsType(PaymentTermsTypes.STANDARD).
		SetDefaultPaymentMethod(*NewPaymentMethod().
			SetId(386092).
			SetName("Credit card")).
		SetEInvoice(true).
		SetCode("AE86").
		SetName("Avv. Maria Rossi").
		SetFirstName("Maria").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("IT12345640962").
		SetTaxCode("BLTGNI5ABCDA794E").
		SetAddressStreet("Via Roma, 1").
		SetAddressPostalCode("20900").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetEmail("maria.rossi@example.com").
		SetCertifiedEmail("maria.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("").
		SetNotes("").
		SetDefaultPaymentTerms(1).
		SetBankName("Indesa").
		SetBankIban("IT40P123456781000000123456").
		SetBankSwiftCode("AK86PCT").
		SetShippingAddress("Corso Magellano 4").
		SetEiCode("111111").
		SetCreatedAt("2021-04-29 08:53:07").
		SetUpdatedAt("2021-04-29 08:53:07")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListClients(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":16451,"code":"AE86","name":"Avv. Maria Rossi","type":"person","first_name":"Maria","last_name":"Rossi","contact_person":"","vat_number":"IT12345640962","tax_code":"BLTGNI5ABCDA794E","address_street":"Via Roma, 1","address_postal_code":"20900","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"maria.rossi@example.com","certified_email":"maria.rossi@pec.example.com","phone":"1234567890","fax":"","notes":"","created_at":"2021-04-29 08:53:07","updated_at":"2021-04-29 08:53:07","default_payment_terms":1,"default_payment_terms_type":"standard","bank_name":"Indesa","bank_iban":"IT40P123456781000000123456","bank_swift_code":"AK86PCT","shipping_address":"Corso Magellano 4","e_invoice":true,"ei_code":"111111","default_vat":{"id":54321,"value":45,"description":"","is_disabled":false},"default_payment_method":{"id":386092,"name":"Credit card"}}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ClientsApi.ListClients(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewClient().
		SetId(16451).
		SetType(ClientTypes.PERSON).
		SetCountry("Italia").
		SetDefaultVat(*NewVatType().
			SetId(54321).
			SetValue(45).
			SetDescription("").
			SetIsDisabled(false)).
		SetDefaultPaymentTermsType(PaymentTermsTypes.STANDARD).
		SetDefaultPaymentMethod(*NewPaymentMethod().
			SetId(386092).
			SetName("Credit card")).
		SetEInvoice(true).
		SetCode("AE86").
		SetName("Avv. Maria Rossi").
		SetFirstName("Maria").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("IT12345640962").
		SetTaxCode("BLTGNI5ABCDA794E").
		SetAddressStreet("Via Roma, 1").
		SetAddressPostalCode("20900").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetEmail("maria.rossi@example.com").
		SetCertifiedEmail("maria.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("").
		SetNotes("").
		SetDefaultPaymentTerms(1).
		SetBankName("Indesa").
		SetBankIban("IT40P123456781000000123456").
		SetBankSwiftCode("AK86PCT").
		SetShippingAddress("Corso Magellano 4").
		SetEiCode("111111").
		SetCreatedAt("2021-04-29 08:53:07").
		SetUpdatedAt("2021-04-29 08:53:07")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyClient(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":16451,"code":"AE86","name":"Avv. Maria Rossi","type":"person","first_name":"Maria","last_name":"Rossi","contact_person":"","vat_number":"IT12345640962","tax_code":"BLTGNI5ABCDA794E","address_street":"Via Roma, 1","address_postal_code":"20900","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"maria.rossi@example.com","certified_email":"maria.rossi@pec.example.com","phone":"1234567890","fax":"","notes":"","created_at":"2021-04-29 08:53:07","updated_at":"2021-04-29 08:53:07","default_payment_terms":1,"default_payment_terms_type":"standard","bank_name":"Indesa","bank_iban":"IT40P123456781000000123456","bank_swift_code":"AK86PCT","shipping_address":"Corso Magellano 4","e_invoice":true,"ei_code":"111111","default_vat":{"id":54321,"value":45,"description":"","is_disabled":false},"default_payment_method":{"id":386092,"name":"Credit card"}}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ClientsApi.ModifyClient(context.Background(), 2, 2).ModifyClientRequest(*NewModifyClientRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewClient().
		SetId(16451).
		SetType(ClientTypes.PERSON).
		SetCountry("Italia").
		SetDefaultVat(*NewVatType().
			SetId(54321).
			SetValue(45).
			SetDescription("").
			SetIsDisabled(false)).
		SetDefaultPaymentTermsType(PaymentTermsTypes.STANDARD).
		SetDefaultPaymentMethod(*NewPaymentMethod().
			SetId(386092).
			SetName("Credit card")).
		SetEInvoice(true).
		SetCode("AE86").
		SetName("Avv. Maria Rossi").
		SetFirstName("Maria").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("IT12345640962").
		SetTaxCode("BLTGNI5ABCDA794E").
		SetAddressStreet("Via Roma, 1").
		SetAddressPostalCode("20900").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetEmail("maria.rossi@example.com").
		SetCertifiedEmail("maria.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("").
		SetNotes("").
		SetDefaultPaymentTerms(1).
		SetBankName("Indesa").
		SetBankIban("IT40P123456781000000123456").
		SetBankSwiftCode("AK86PCT").
		SetShippingAddress("Corso Magellano 4").
		SetEiCode("111111").
		SetCreatedAt("2021-04-29 08:53:07").
		SetUpdatedAt("2021-04-29 08:53:07")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
