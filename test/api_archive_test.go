/*
Fatture in Cloud API v2 - API Reference

Testing ArchiveApiService

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

func Test_fattureincloud_ArchiveApiService(t *testing.T) {

    configuration := fattureincloud.NewConfiguration()
    apiClient := fattureincloud.NewAPIClient(configuration)

    t.Run("Test ArchiveApiService CreateArchiveDocument", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32

        resp, httpRes, err := apiClient.ArchiveApi.CreateArchiveDocument(context.Background(), companyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArchiveApiService DeleteArchiveDocument", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32
        var documentId int32

        resp, httpRes, err := apiClient.ArchiveApi.DeleteArchiveDocument(context.Background(), companyId, documentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArchiveApiService GetArchiveDocument", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32
        var documentId int32

        resp, httpRes, err := apiClient.ArchiveApi.GetArchiveDocument(context.Background(), companyId, documentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArchiveApiService ListArchiveDocuments", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32

        resp, httpRes, err := apiClient.ArchiveApi.ListArchiveDocuments(context.Background(), companyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArchiveApiService ModifyArchiveDocument", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32
        var documentId int32

        resp, httpRes, err := apiClient.ArchiveApi.ModifyArchiveDocument(context.Background(), companyId, documentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArchiveApiService UploadArchiveDocumentAttachment", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var companyId int32

        resp, httpRes, err := apiClient.ArchiveApi.UploadArchiveDocumentAttachment(context.Background(), companyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
