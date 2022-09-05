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

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateSupplier(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"code":"AE86","name":"Mario Rossi S.R.L.","type":"company","first_name":"Mario","last_name":"Rossi","contact_person":"","vat_number":"111222333","tax_code":"111122233","address_street":"Corso Magellano, 46","address_postal_code":"20146","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"mario.rossi@example.com","certified_email":"mario.rossi@pec.example.com","phone":"1234567890","fax":"123456789","notes":"","created_at":"2021-15-08","updated_at":"2021-15-08"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SuppliersApi.CreateSupplier(context.Background(), 2).CreateSupplierRequest(*NewCreateSupplierRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewSupplier().
		SetId(12345).
		SetCode("AE86").
		SetName("Mario Rossi S.R.L.").
		SetType(SupplierTypes.COMPANY).
		SetFirstName("Mario").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("111222333").
		SetTaxCode("111122233").
		SetAddressStreet("Corso Magellano, 46").
		SetAddressPostalCode("20146").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetCountry("Italia").
		SetEmail("mario.rossi@example.com").
		SetCertifiedEmail("mario.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("123456789").
		SetNotes("").
		SetCreatedAt("2021-15-08").
		SetUpdatedAt("2021-15-08")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteSupplier(t *testing.T) {
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

	resp, err := apiClient.SuppliersApi.DeleteSupplier(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetSupplier(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"code":"AE86","name":"Mario Rossi S.R.L.","type":"company","first_name":"Mario","last_name":"Rossi","contact_person":"","vat_number":"111222333","tax_code":"111122233","address_street":"Corso Magellano, 46","address_postal_code":"20146","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"mario.rossi@example.com","certified_email":"mario.rossi@pec.example.com","phone":"1234567890","fax":"123456789","notes":"","created_at":"2021-15-08","updated_at":"2021-15-08"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SuppliersApi.GetSupplier(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewSupplier().
		SetId(12345).
		SetCode("AE86").
		SetName("Mario Rossi S.R.L.").
		SetType(SupplierTypes.COMPANY).
		SetFirstName("Mario").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("111222333").
		SetTaxCode("111122233").
		SetAddressStreet("Corso Magellano, 46").
		SetAddressPostalCode("20146").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetCountry("Italia").
		SetEmail("mario.rossi@example.com").
		SetCertifiedEmail("mario.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("123456789").
		SetNotes("").
		SetCreatedAt("2021-15-08").
		SetUpdatedAt("2021-15-08")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListSuppliers(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12345,"code":"AE86","name":"Mario Rossi S.R.L.","type":"company","first_name":"Mario","last_name":"Rossi","contact_person":"","vat_number":"111222333","tax_code":"111122233","address_street":"Corso Magellano, 46","address_postal_code":"20146","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"mario.rossi@example.com","certified_email":"mario.rossi@pec.example.com","phone":"1234567890","fax":"123456789","notes":"","created_at":"2021-15-08","updated_at":"2021-15-08"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SuppliersApi.ListSuppliers(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewSupplier().
		SetId(12345).
		SetCode("AE86").
		SetName("Mario Rossi S.R.L.").
		SetType(SupplierTypes.COMPANY).
		SetFirstName("Mario").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("111222333").
		SetTaxCode("111122233").
		SetAddressStreet("Corso Magellano, 46").
		SetAddressPostalCode("20146").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetCountry("Italia").
		SetEmail("mario.rossi@example.com").
		SetCertifiedEmail("mario.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("123456789").
		SetNotes("").
		SetCreatedAt("2021-15-08").
		SetUpdatedAt("2021-15-08")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifySupplier(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"code":"AE86","name":"Mario Rossi S.R.L.","type":"company","first_name":"Mario","last_name":"Rossi","contact_person":"","vat_number":"111222333","tax_code":"111122233","address_street":"Corso Magellano, 46","address_postal_code":"20146","address_city":"Milano","address_province":"MI","address_extra":"","country":"Italia","email":"mario.rossi@example.com","certified_email":"mario.rossi@pec.example.com","phone":"1234567890","fax":"123456789","notes":"","created_at":"2021-15-08","updated_at":"2021-15-08"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.SuppliersApi.ModifySupplier(context.Background(), 2, 2).ModifySupplierRequest(*NewModifySupplierRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewSupplier().
		SetId(12345).
		SetCode("AE86").
		SetName("Mario Rossi S.R.L.").
		SetType(SupplierTypes.COMPANY).
		SetFirstName("Mario").
		SetLastName("Rossi").
		SetContactPerson("").
		SetVatNumber("111222333").
		SetTaxCode("111122233").
		SetAddressStreet("Corso Magellano, 46").
		SetAddressPostalCode("20146").
		SetAddressCity("Milano").
		SetAddressProvince("MI").
		SetAddressExtra("").
		SetCountry("Italia").
		SetEmail("mario.rossi@example.com").
		SetCertifiedEmail("mario.rossi@pec.example.com").
		SetPhone("1234567890").
		SetFax("123456789").
		SetNotes("").
		SetCreatedAt("2021-15-08").
		SetUpdatedAt("2021-15-08")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
