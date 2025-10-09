package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OAuth2DeviceCodeManager struct {
	OAuth2Manager
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

// Device code response wrapping
type DataDeviceCode struct {
	// Data
	Data OAuth2DeviceCodeResponse `json:"data,omitempty"`
}

// Initialize a new instance of the OAuth2DeviceCodeManager class
func NewOAuth2DeviceCodeManager(clientId string) *OAuth2DeviceCodeManager {
	this := OAuth2DeviceCodeManager{}
	this.clientId = clientId
	this.baseUri = "https://api-v2.fattureincloud.it"
	return &this
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
	scopesString := GetScopesString(scopes)

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
	res := DataDeviceCode{}
	json.Unmarshal(body, &res)

	return &res.Data, nil

}

// Retrieve the access token.
func (o *OAuth2DeviceCodeManager) FetchToken(code string) (*OAuth2TokenResponse, error) {
	data := map[string]string{
		"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
		"client_id":   o.clientId,
		"device_code": code,
	}

	return ExecuteTokenPost(o.baseUri, data)
}

// Refresh the access token.
func (o *OAuth2DeviceCodeManager) RefreshToken(refreshToken string) (*OAuth2TokenResponse, error) {
	data := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     o.clientId,
		"refresh_token": refreshToken,
	}

	return ExecuteTokenPost(o.baseUri, data)
}
