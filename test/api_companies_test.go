/*
Fatture in Cloud API v2 - API Reference

Testing CompaniesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fattureincloud

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    fattureincloud "./openapi"
)

func Test_fattureincloud_CompaniesApiService(t *testing.T) {

    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)

    t.Run("Test CompaniesApiService GetCompanyInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32

        resp, httpRes, err := apiClient.CompaniesApi.GetCompanyInfo(context.Background(), companyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
