package oauth

import (
	"testing"

	oauth "github.com/fattureincloud/fattureincloud-go-sdk/v2/oauth2"
	"github.com/stretchr/testify/assert"
)

var auth = oauth.NewOAuth2AuthorizationCodeManager("CLIENT_ID", "CLIENT_SECRET", "http://localhost:8000/redirect")

func TestOAuth2AuthorizationCodeManager(t *testing.T) {
	assert.Equal(t, auth.GetClientId(), "CLIENT_ID")
	assert.Equal(t, auth.GetClientSecret(), "CLIENT_SECRET")
	assert.Equal(t, auth.GetRedirectUri(), "http://localhost:8000/redirect")
}

func TestOAuth2AuthorizationCodeParams(t *testing.T) {
	params := oauth.NewOAuth2AuthorizationCodeParams("EXAMPLE_CODE", "EXAMPLE_STATE")
	assert.Equal(t, params.AuthorizationCode, "EXAMPLE_CODE")
	assert.Equal(t, params.State, "EXAMPLE_STATE")
}

func TestGetAuthorizationUrl(t *testing.T) {
	scopes := []oauth.Scope{oauth.Scopes.ENTITY_CLIENTS_ALL, oauth.Scopes.ENTITY_CLIENTS_READ}
	url := auth.GetAuthorizationUrl(scopes, "EXAMPLE_STATE")
	assert.Equal(t, "https://api-v2.fattureincloud.it/oauth/authorize?client_id=CLIENT_ID&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2Fredirect&response_type=code&scope=entity.clients%3Aa+entity.clients%3Ar&state=EXAMPLE_STATE", url)
}

func TestGetParamsFromUrl(t *testing.T) {
	url := "http://localhost:3000/redirect?state=EXAMPLE_STATE&code=c%2FEXAMPLE_CODE"
	parameters, _ := auth.GetParamsFromUrl(url)
	assert.Equal(t, "c/EXAMPLE_CODE", parameters.AuthorizationCode)
	assert.Equal(t, "EXAMPLE_STATE", parameters.State)
}
