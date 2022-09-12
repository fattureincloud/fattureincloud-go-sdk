/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestPermissionLevelsResponse(t *testing.T) {
	assert.Equal(t, "none", string(PermissionLevels.NONE))
	assert.Equal(t, "read", string(PermissionLevels.READ))
	assert.Equal(t, "write", string(PermissionLevels.WRITE))
	assert.Equal(t, "detailed", string(PermissionLevels.DETAILED))
}