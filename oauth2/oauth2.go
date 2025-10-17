package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OAuth2
type OAuth2Manager struct {
	// Client id
	clientId string
	// Base uri
	baseUri string
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

// Build the Scope string.
func GetScopesString(scopes []Scope) string {

	joinedScopes := ""
	for _, scope := range scopes {
		joinedScopes += string(scope) + " "
	}
	return strings.TrimRight(joinedScopes, " ")

}

func ExecuteTokenPost(baseUri string, data map[string]string) (*OAuth2TokenResponse, error) {
	tokenUri := fmt.Sprintf("%s/oauth/token", baseUri)

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
