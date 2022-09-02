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

func TestCreateArchiveDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"date":"2021-08-20","description":"spesa 2","category":"Altri documenti","attachment_token":"jwfbaiuwbfoiewfoa8weohafw7gefa9we"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ArchiveApi.CreateArchiveDocument(context.Background(), 2).CreateArchiveDocumentRequest(*NewCreateArchiveDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewArchiveDocument().
		SetId(12345).
		SetDate("2021-08-20").
		SetDescription("spesa 2").
		SetCategory("Altri documenti").
		SetAttachmentToken("jwfbaiuwbfoiewfoa8weohafw7gefa9we")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestDeleteArchiveDocument(t *testing.T) {
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

	actual, err := apiClient.ArchiveApi.DeleteArchiveDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, actual.StatusCode)
}

func TestGetArchiveDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"date":"2021-08-20","description":"spesa 2","category":"Altri documenti","attachment_token":"jwfbaiuwbfoiewfoa8weohafw7gefa9we"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ArchiveApi.GetArchiveDocument(context.Background(), 2, 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewArchiveDocument().
		SetId(12345).
		SetDate("2021-08-20").
		SetDescription("spesa 2").
		SetCategory("Altri documenti").
		SetAttachmentToken("jwfbaiuwbfoiewfoa8weohafw7gefa9we")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestListArchiveDocuments(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":12345,"date":"2021-08-20","description":"spesa 2","category":"Altri documenti","attachment_token":"jwfbaiuwbfoiewfoa8weohafw7gefa9we"}]}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ArchiveApi.ListArchiveDocuments(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewArchiveDocument().
		SetId(12345).
		SetDate("2021-08-20").
		SetDescription("spesa 2").
		SetCategory("Altri documenti").
		SetAttachmentToken("jwfbaiuwbfoiewfoa8weohafw7gefa9we")

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[0]))
}

func TestModifyArchiveDocument(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12345,"date":"2021-08-20","description":"spesa 2","category":"Altri documenti","attachment_token":"jwfbaiuwbfoiewfoa8weohafw7gefa9we"}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.ArchiveApi.ModifyArchiveDocument(context.Background(), 2, 2).ModifyArchiveDocumentRequest(*NewModifyArchiveDocumentRequest()).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewArchiveDocument().
		SetId(12345).
		SetDate("2021-08-20").
		SetDescription("spesa 2").
		SetCategory("Altri documenti").
		SetAttachmentToken("jwfbaiuwbfoiewfoa8weohafw7gefa9we")

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}

func TestUploadArchiveDocumentAttachment(t *testing.T) {
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
	_, resp, err := apiClient.ArchiveApi.UploadArchiveDocumentAttachment(context.Background(), 2).Filename("filename").Execute()
	assert.NoError(t, err, "errore in chiamata api")

	assert.Equal(t, 200, resp.StatusCode)
}
