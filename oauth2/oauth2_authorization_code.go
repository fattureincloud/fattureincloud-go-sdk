package oauth

import (
	"fmt"
	"net/url"
)

// TODO: riordiorganizzare i metodi

type OAuth2AuthorizationCodeManager struct {
	OAuth2Manager
	// Client secret
	clientSecret string
	// Redirect uri
	redirectUri string
}

// Initialize a new instance of the OAuth2AuthorizationCodeManager class
func NewOAuth2AuthorizationCodeManager(clientId string, clientSecret string, redirecUri string) *OAuth2AuthorizationCodeManager {
	this := OAuth2AuthorizationCodeManager{}
	this.clientId = clientId
	this.clientSecret = clientSecret
	this.redirectUri = redirecUri
	this.baseUri = "https://api-v2.fattureincloud.it"
	return &this
}

// Sets ClientId
func (o *OAuth2AuthorizationCodeManager) SetClientId(v string) *OAuth2AuthorizationCodeManager {
	o.clientId = v
	return o
}

// Gets ClientId
func (o *OAuth2AuthorizationCodeManager) GetClientId() string {
	return o.clientId
}

// Sets ClientSecret
func (o *OAuth2AuthorizationCodeManager) SetClientSecret(v string) *OAuth2AuthorizationCodeManager {
	o.clientSecret = v
	return o
}

// Sets ClientSecret
func (o *OAuth2AuthorizationCodeManager) GetClientSecret() string {
	return o.clientSecret
}

// Sets BaseUri
func (o *OAuth2AuthorizationCodeManager) SetBaseUri(v string) *OAuth2AuthorizationCodeManager {
	o.baseUri = v
	return o
}

// Gets BaseUri
func (o *OAuth2AuthorizationCodeManager) GetBaseUri() string {
	return o.baseUri
}

// Sets RedirectUri
func (o *OAuth2AuthorizationCodeManager) SetRedirectUri(v string) *OAuth2AuthorizationCodeManager {
	o.redirectUri = v
	return o
}

// Gets RedirectUri
func (o *OAuth2AuthorizationCodeManager) GetRedirectUri() string {
	return o.redirectUri
}

// Get the Authorization url
func (o *OAuth2AuthorizationCodeManager) GetAuthorizationUrl(scopes []Scope, state string) string {
	authorizationUri := fmt.Sprintf("%s/oauth/authorize", o.baseUri)
	scopesString := GetScopesString(scopes)

	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("client_id", o.clientId)
	values.Add("redirect_uri", o.redirectUri)
	values.Add("scope", scopesString)
	values.Add("state", state)
	query := values.Encode()

	return authorizationUri + "?" + query
}

// Get code and state parameters from url
func (o *OAuth2AuthorizationCodeManager) GetParamsFromUrl(uri string) (*OAuth2AuthorizationCodeParams, error) {
	urlObj, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	queryParams := urlObj.Query()

	state := queryParams.Get("state")
	code := queryParams.Get("code")

	return NewOAuth2AuthorizationCodeParams(code, state), nil
}

// Retrieve the access token.
func (o *OAuth2AuthorizationCodeManager) FetchToken(code string) (*OAuth2TokenResponse, error) {
	data := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     o.clientId,
		"client_secret": o.clientSecret,
		"redirect_uri":  o.redirectUri,
		"code":          code,
	}

	return ExecuteTokenPost(o.baseUri, data)
}

// Refresh the access token.
func (o *OAuth2AuthorizationCodeManager) RefreshToken(refreshToken string) (*OAuth2TokenResponse, error) {
	data := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     o.clientId,
		"client_secret": o.clientSecret,
		"refresh_token": refreshToken,
	}

	return ExecuteTokenPost(o.baseUri, data)
}

// OAuth2AuthorizationCodeParams
type OAuth2AuthorizationCodeParams struct {
	// State
	state string
	// Authorization code
	authorizationCode string
}

// Initialize a new instance of the OAuth2AuthorizationCodeParams class
func NewOAuth2AuthorizationCodeParams(code string, state string) *OAuth2AuthorizationCodeParams {
	this := OAuth2AuthorizationCodeParams{}
	this.authorizationCode = code
	this.state = state
	return &this
}

// Sets AuthorizationCode
func (o *OAuth2AuthorizationCodeParams) SetAuthorizationCode(v string) *OAuth2AuthorizationCodeParams {
	o.authorizationCode = v
	return o
}

// Sets State
func (o *OAuth2AuthorizationCodeParams) SetState(v string) *OAuth2AuthorizationCodeParams {
	o.state = v
	return o
}

// Gets AuthorizationCode
func (o *OAuth2AuthorizationCodeParams) GetAuthorizationCode() string {
	return o.authorizationCode
}

// Gets State
func (o *OAuth2AuthorizationCodeParams) GetState() string {
	return o.state
}
