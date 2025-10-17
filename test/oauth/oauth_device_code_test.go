package oauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	oauth "github.com/fattureincloud/fattureincloud-go-sdk/v2/oauth2"
	"github.com/stretchr/testify/assert"
)

func TestOAuth2DeviceCodeManager(t *testing.T) {
	var dev = oauth.NewOAuth2DeviceCodeManager("CLIENT_ID")
	assert.Equal(t, dev.GetClientId(), "CLIENT_ID")
	assert.Equal(t, dev.GetBaseUri(), "https://api-v2.fattureincloud.it")

	dev.SetClientId("NEW_CLIENT_ID")
	dev.SetBaseUri("https://api.fattureincloud.it")

	assert.Equal(t, dev.GetClientId(), "NEW_CLIENT_ID")
	assert.Equal(t, dev.GetBaseUri(), "https://api.fattureincloud.it")
}

func TestGetDeviceCode(t *testing.T) {
	stream := `{"data": {"device_code": "EXAMPLE_DEVICE_CODE", "user_code": "EXAMPLE_USER_CODE", "scope": {"entity.clients:a": "entity.clients:a"}, "verification_uri": "http://localhost:8000/verify", "expires_in": 900, "interval": 5}}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(stream))
	}))
	defer server.Close()

	authTest := oauth.NewOAuth2DeviceCodeManager("CLIENT_ID")
	authTest.SetBaseUri(server.URL)
	params, err := authTest.GetDeviceCode([]oauth.Scope{oauth.Scopes.ENTITY_CLIENTS_ALL, oauth.Scopes.ENTITY_CLIENTS_READ})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	assert.Equal(t, "EXAMPLE_DEVICE_CODE", params.DeviceCode)
	assert.Equal(t, "EXAMPLE_USER_CODE", params.UserCode)
	assert.Equal(t, map[string]string{"entity.clients:a": "entity.clients:a"}, params.Scope)
	assert.Equal(t, "http://localhost:8000/verify", params.VerificationUri)
	assert.Equal(t, 900, params.ExpiresIn)
	assert.Equal(t, 5, params.Interval)
}

func TestDeviceFetchToken(t *testing.T) {
	stream := `{"token_type": "bearer", "access_token": "a/ACCESS_TOKEN", "refresh_token": "r/REFRESH_TOKEN", "expires_in": 86400}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(stream))
	}))
	defer server.Close()

	authTest := oauth.NewOAuth2DeviceCodeManager("CLIENT_ID")
	authTest.SetBaseUri(server.URL)

	token, err := authTest.FetchToken("a/EXAMPLE_CODE")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, "bearer", token.TokenType)
	assert.Equal(t, "a/ACCESS_TOKEN", token.AccessToken)
	assert.Equal(t, "r/REFRESH_TOKEN", token.RefreshToken)
	assert.Equal(t, int(86400), token.ExpiresIn)
}

func TestDeviceRefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token_type": "bearer", "access_token": "a/ACCESS_TOKEN", "refresh_token": "r/REFRESH_TOKEN", "expires_in": 86400}`))
	}))
	defer server.Close()

	authTest := oauth.NewOAuth2DeviceCodeManager("CLIENT_ID")
	authTest.SetBaseUri(server.URL)

	token, err := authTest.RefreshToken("r/REFRESH_TOKEN")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, "bearer", token.TokenType)
	assert.Equal(t, "a/ACCESS_TOKEN", token.AccessToken)
	assert.Equal(t, "r/REFRESH_TOKEN", token.RefreshToken)
	assert.Equal(t, int(86400), token.ExpiresIn)
}
