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

func TestGetUserInfo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"name":"Mario Rossi","first_name":"Mario","last_name":"Rossi","email":"mario.rossi@example.com","hash":"5add29e1234532a1bf2ed7b612043029","picture":"picture.jpg"},"info":{"need_marketing_consents_confirmation":false,"need_password_change":false,"need_terms_of_service_confirmation":false},"email_confirmation_state":{"need_confirmation":false}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.UserApi.GetUserInfo(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewUser().
		SetId(12345).
		SetName("Mario Rossi").
		SetFirstName("Mario").
		SetLastName("Rossi").
		SetEmail("mario.rossi@example.com").
		SetHash("5add29e1234532a1bf2ed7b612043029").
		SetPicture("picture.jpg")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListUserCompanies(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"companies":[{"id":2,"name":"Studio Commercialista","tax_code":"SLVMTT50A01F205L","type":"accountant","connection_id":94816,"controlled_companies":[{"id":5,"name":"Azienda 1","tax_code":"SLVMTT50A01F205L","type":"company","connection_id":94817}]}]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.UserApi.ListUserCompanies(context.Background()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCompany().
		SetId(2).
		SetName("Studio Commercialista").
		SetTaxCode("SLVMTT50A01F205L").
		SetType(CompanyTypes.ACCOUNTANT).
		SetConnectionId(94816).
		SetControlledCompanies([]ControlledCompany{
			*NewControlledCompany().
				SetId(5).
				SetName("Azienda 1").
				SetTaxCode("SLVMTT50A01F205L").
				SetType(CompanyTypes.COMPANY).
				SetConnectionId(94817)})

	assert.True(t, reflect.DeepEqual(expected, &actual.Data.Get().Companies[0]))
}
