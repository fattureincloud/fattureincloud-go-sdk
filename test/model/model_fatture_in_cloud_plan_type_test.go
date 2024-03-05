/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package model

import (
	"testing"

	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestFattureInCloudPlanType(t *testing.T) {
	assert.Equal(t, "trial", string(FattureInCloudPlanTypes.TRIAL))
	assert.Equal(t, "standard", string(FattureInCloudPlanTypes.STANDARD))
	assert.Equal(t, "premium", string(FattureInCloudPlanTypes.PREMIUM))
	assert.Equal(t, "premium_plus", string(FattureInCloudPlanTypes.PREMIUM_PLUS))
	assert.Equal(t, "complete", string(FattureInCloudPlanTypes.COMPLETE))
}
