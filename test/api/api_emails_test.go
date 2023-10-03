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
	"time"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestListEmails(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"current_page":1,"data":[{"id":22,"status":"sent","sent_date":"2022-07-17T13:53:12Z","errors_count":0,"error_log":"","from_email":"test@mail.it","from_name":"Test mail","to_email":"mail@test.it","to_name":"Mario","subject":"Test","content":"Test send email","copy_to":"","recipient_status":"unknown","recipient_date":"2022-07-18T13:53:12Z","kind":"Fatture","attachments":[]},{"id":1,"status":"sent","sent_date":"2022-07-18T13:53:12Z","errors_count":0,"error_log":"","from_email":"test@mail.it","from_name":"Test mail","to_email":"mail@test.it","to_name":"Maria","subject":"Test","content":"Test send email","copy_to":"","recipient_status":"unknown","recipient_date":"2022-07-18T13:53:12Z","kind":"Fatture","attachments":[]}],"first_page_url":"emails?page=1","next_page_url":"emails?page=1","from":1,"last_page":1,"last_page_url":"emails?page=1","path":"emails","per_page":50,"prev_page_url":"emails?page=1","to":2,"total":2}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.EmailsAPI.ListEmails(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewEmail().
		SetId(1).
		SetStatus(EmailStatuses.SENT).
		SetSentDate(time.Date(2022, 7, 18, 13, 53, 12, 0, time.UTC)).
		SetErrorsCount(0).
		SetErrorLog("").
		SetFromEmail("test@mail.it").
		SetFromName("Test mail").
		SetToEmail("mail@test.it").
		SetToName("Maria").
		SetSubject("Test").
		SetContent("Test send email").
		SetCopyTo("").
		SetRecipientStatus(EmailRecipientStatuses.UNKNOWN).
		SetRecipientDate(time.Date(2022, 7, 18, 13, 53, 12, 0, time.UTC)).
		SetKind("Fatture").
		SetAttachments([]EmailAttachment{})

	assert.True(t, reflect.DeepEqual(expected, &actual.Data[1]))
}
