/*
Kenar API

Testing PostAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kenarapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/divar-ir/kenar-sdk-go"
)

func Test_kenarapi_PostAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PostAPIService PostEditPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postToken string

		resp, httpRes, err := apiClient.PostAPI.PostEditPost(context.Background(), postToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostGetImageUploadURL", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostAPI.PostGetImageUploadURL(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostGetPostStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postToken string

		resp, httpRes, err := apiClient.PostAPI.PostGetPostStats(context.Background(), postToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostSubmitEmergencyResidencePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostAPI.PostSubmitEmergencyResidencePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostAPIService PostSubmitPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostAPI.PostSubmitPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
