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
	"github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateWebhooksSubscription(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"SUB123","sink":"https://endpoint.test","verified":true,"types":["it.fattureincloud.webhooks.cashbook.create"]},"warnings":["error"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.WebhooksApi.CreateWebhooksSubscription(context.Background(), 2).CreateWebhooksSubscriptionRequest(*NewCreateWebhooksSubscriptionRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewWebhooksSubscription().
		SetId("SUB123").
		SetSink("https://endpoint.test").
		SetVerified(true).
		SetTypes([]model.EventType{EventTypes.CASHBOOK_CREATE})

	expectedWarnings := []string{"error"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
	assert.True(t, reflect.DeepEqual(expectedWarnings, actual.Warnings))
}

func TestDeleteWebhooksSubscription(t *testing.T) {
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

	actual, err := apiClient.WebhooksApi.DeleteWebhooksSubscription(context.Background(), 2, "SUB123").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, actual.StatusCode)
}

func TestGetWebhooksSubscription(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"SUB123","sink":"https://endpoint.test","verified":true,"types":["it.fattureincloud.webhooks.cashbook.create"]}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.WebhooksApi.GetWebhooksSubscription(context.Background(), 2, "SUB123").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewWebhooksSubscription().
		SetId("SUB123").
		SetSink("https://endpoint.test").
		SetVerified(true).
		SetTypes([]model.EventType{EventTypes.CASHBOOK_CREATE})

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListWebhooksSubscriptions(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":"SUB123","sink":"https://endpoint.test","verified":true,"types":["it.fattureincloud.webhooks.cashbook.create"]},{"id":"SUB12345","sink":"https://endpoint.test","verified":true,"types":["it.fattureincloud.webhooks.cashbook.update"]}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.WebhooksApi.ListWebhooksSubscriptions(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewWebhooksSubscription().
		SetId("SUB123").
		SetSink("https://endpoint.test").
		SetVerified(true).
		SetTypes([]model.EventType{EventTypes.CASHBOOK_CREATE})

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyWebhooksSubscription(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":"SUB123","sink":"https://endpoint.test","verified":true,"types":["it.fattureincloud.webhooks.cashbook.create"]},"warnings":["error"]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.WebhooksApi.ModifyWebhooksSubscription(context.Background(), 2, "SUB123").ModifyWebhooksSubscriptionRequest(*NewModifyWebhooksSubscriptionRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewWebhooksSubscription().
		SetId("SUB123").
		SetSink("https://endpoint.test").
		SetVerified(true).
		SetTypes([]model.EventType{EventTypes.CASHBOOK_CREATE})

	expectedWarnings := []string{"error"}

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
	assert.True(t, reflect.DeepEqual(expectedWarnings, actual.Warnings))
}
