package oauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	oauth "github.com/fattureincloud/fattureincloud-go-sdk/v2/oauth2"
	"github.com/stretchr/testify/assert"
)

/**
 * Test OAuth2AuthorizationCodeManager
 */
func TestOAuth2AuthorizationCodeManager(t *testing.T) {
	var auth = oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")
	assert.Equal(t, auth.GetClientId(), "CLIENT_ID")
	assert.Equal(t, auth.GetClientSecret(), "CLIENT_SECRET")
	assert.Equal(t, auth.GetRedirectUri(), "http://localhost:8000/redirect")
	assert.Equal(t, auth.GetBaseUri(), "https://api-v2.fattureincloud.it")

	auth.SetClientId("NEW_CLIENT_ID")
	auth.SetClientSecret("NEW_CLIENT_SECRET")
	auth.SetRedirectUri("http://localhost:8000/new_redirect")
	auth.SetBaseUri("https://api.fattureincloud.it")

	assert.Equal(t, auth.GetClientId(), "NEW_CLIENT_ID")
	assert.Equal(t, auth.GetClientSecret(), "NEW_CLIENT_SECRET")
	assert.Equal(t, auth.GetRedirectUri(), "http://localhost:8000/new_redirect")
	assert.Equal(t, auth.GetBaseUri(), "https://api.fattureincloud.it")
}

/**
 * Test OAuth2AuthorizationCodeParams
 */
func TestOAuth2AuthorizationCodeParams(t *testing.T) {
	params := oauth.NewOAuth2AuthorizationCodeParams("a/EXAMPLE_CODE", "EXAMPLE_STATE")
	assert.Equal(t, params.GetAuthorizationCode(), "a/EXAMPLE_CODE")
	assert.Equal(t, params.GetState(), "EXAMPLE_STATE")

	params.SetAuthorizationCode("a/NEW_EXAMPLE_CODE")
	params.SetState("NEW_EXAMPLE_STATE")
	assert.Equal(t, params.GetAuthorizationCode(), "a/NEW_EXAMPLE_CODE")
	assert.Equal(t, params.GetState(), "NEW_EXAMPLE_STATE")
}

func TestGetAuthorizationUrl(t *testing.T) {
	var auth = oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")
	scopes := []oauth.Scope{oauth.Scopes.ENTITY_CLIENTS_ALL, oauth.Scopes.ENTITY_CLIENTS_READ}
	url := auth.GetAuthorizationUrl(scopes, "EXAMPLE_STATE")
	assert.Equal(t, "https://api-v2.fattureincloud.it/oauth/authorize?client_id=CLIENT_ID&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2Fredirect&response_type=code&scope=entity.clients%3Aa+entity.clients%3Ar&state=EXAMPLE_STATE", url)
}

func TestAuthFetchToken(t *testing.T) {
	stream := `{"token_type": "bearer", "access_token": "a/ACCESS_TOKEN", "refresh_token": "r/REFRESH_TOKEN", "expires_in": 86400}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(stream))
	}))
	defer server.Close()

	authTest := oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")
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

func TestAuthRefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token_type": "bearer", "access_token": "a/ACCESS_TOKEN", "refresh_token": "r/REFRESH_TOKEN", "expires_in": 86400}`))
	}))
	defer server.Close()

	authTest := oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")
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

func TestGetParamsFromUrl(t *testing.T) {
	url := "http://localhost:3000/redirect?state=EXAMPLE_STATE&code=c%2FEXAMPLE_CODE"
	var auth = oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")
	parameters, _ := auth.GetParamsFromUrl(url)
	assert.Equal(t, "c/EXAMPLE_CODE", parameters.GetAuthorizationCode())
	assert.Equal(t, "EXAMPLE_STATE", parameters.GetState())
}
