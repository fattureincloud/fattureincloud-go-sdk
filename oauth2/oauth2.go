package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OAuth2
type OAuth2Manager struct {
	// Client id
	clientId string
	// Base uri
	baseUri string
}

type OAuth2AuthorizationCodeManager struct {
	OAuth2Manager
	// Client secret
	clientSecret string
	// Redirect uri
	redirectUri string
}

type OAuth2DeviceCodeManager struct {
	OAuth2Manager
}

// OAuth2AuthorizationCodeParams
type OAuth2AuthorizationCodeParams struct {
	// State
	State string
	// Authorization code
	AuthorizationCode string
}

// OAuth2TokenResponse
type OAuth2TokenResponse struct {
	// Token Type
	TokenType string `json:"token_type,omitempty"`
	// Access Token
	AccessToken string `json:"access_token,omitempty"`
	// Refresh Token
	RefreshToken string `json:"refresh_token,omitempty"`
	// Expiration time
	ExpiresIn int `json:"expires_in,omitempty"`
}

// OAuth2DeviceCodeResponse
type OAuth2DeviceCodeResponse struct {
	// Device Code
	DeviceCode string `json:"device_code,omitempty"`
	// User Code
	UserCode string `json:"user_code,omitempty"`
	// Scope
	Scope map[string]string `json:"scope,omitempty"`
	// Verification URI
	VerificationUri string `json:"verification_uri,omitempty"`
	// Interval
	Interval int `json:"interval,omitempty"`
	// Expiration time
	ExpiresIn int `json:"expires_in,omitempty"`
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

// Initialize a new instance of the OAuth2DeviceCodeManager class
func NewOAuth2DeviceCodeManager(clientId string) *OAuth2DeviceCodeManager {
	this := OAuth2DeviceCodeManager{}
	this.clientId = clientId
	this.baseUri = "https://api-v2.fattureincloud.it"
	return &this
}

// Initialize a new instance of the OAuth2AuthorizationCodeParams class
func NewOAuth2AuthorizationCodeParams(code string, state string) *OAuth2AuthorizationCodeParams {
	this := OAuth2AuthorizationCodeParams{}
	this.AuthorizationCode = code
	this.State = state
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
	scopesString := getScopesString(scopes)

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
	tokenUri := fmt.Sprintf("%s/oauth/token", o.baseUri)
	data := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     o.clientId,
		"client_secret": o.clientSecret,
		"redirect_uri":  o.redirectUri,
		"code":          code,
	}

	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUri, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	res := OAuth2TokenResponse{}
	json.Unmarshal(body, &res)

	return &res, nil

}

// Refresh the access token.
func (o *OAuth2AuthorizationCodeManager) RefreshToken(refreshToken string) (*OAuth2TokenResponse, error) {
	tokenUri := fmt.Sprintf("%s/oauth/token", o.baseUri)
	data := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     o.clientId,
		"client_secret": o.clientSecret,
		"refresh_token": refreshToken,
	}

	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUri, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	res := OAuth2TokenResponse{}
	json.Unmarshal(body, &res)

	return &res, nil

}

// Sets ClientId
func (o *OAuth2DeviceCodeManager) SetClientId(v string) *OAuth2DeviceCodeManager {
	o.clientId = v
	return o
}

// Gets ClientId
func (o *OAuth2DeviceCodeManager) GetClientId() string {
	return o.clientId
}

// Sets BaseUri
func (o *OAuth2DeviceCodeManager) SetBaseUri(v string) *OAuth2DeviceCodeManager {
	o.baseUri = v
	return o
}

// Gets BaseUri
func (o *OAuth2DeviceCodeManager) GetBaseUri() string {
	return o.baseUri
}

// Retrieve the access token.
func (o *OAuth2DeviceCodeManager) GetDeviceCode(scopes []Scope) (*OAuth2DeviceCodeResponse, error) {
	scopesString := getScopesString(scopes)

	tokenUri := fmt.Sprintf("%s/oauth/device", o.baseUri)
	data := map[string]string{
		"client_id": o.clientId,
		"scope":     scopesString,
	}

	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUri, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	res := OAuth2DeviceCodeResponse{}
	json.Unmarshal(body, &res)

	return &res, nil

}

// Retrieve the access token.
func (o *OAuth2DeviceCodeManager) FetchToken(code string) (*OAuth2TokenResponse, error) {
	tokenUri := fmt.Sprintf("%s/oauth/token", o.baseUri)
	data := map[string]string{
		"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
		"client_id":   o.clientId,
		"device_code": code,
	}

	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUri, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	res := OAuth2TokenResponse{}
	json.Unmarshal(body, &res)

	return &res, nil

}

// Refresh the access token.
func (o *OAuth2DeviceCodeManager) RefreshToken(refreshToken string) (*OAuth2TokenResponse, error) {
	tokenUri := fmt.Sprintf("%s/oauth/token", o.baseUri)
	data := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     o.clientId,
		"refresh_token": refreshToken,
	}

	reqBody, _ := json.Marshal(data)
	resp, err := http.Post(tokenUri, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	res := OAuth2TokenResponse{}
	json.Unmarshal(body, &res)

	return &res, nil

}

// Build the Scope string.
func getScopesString(scopes []Scope) string {

	joinedScopes := ""
	for _, scope := range scopes {
		joinedScopes += string(scope) + " "
	}
	return strings.TrimRight(joinedScopes, " ")

}
