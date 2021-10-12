package goapplesignin

import (
	"errors"
	"testing"

	"github.com/BolajiOlajide/go-apple-signin/mocks"
	"github.com/BolajiOlajide/go-apple-signin/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationURLWithoutClientID(t *testing.T) {
	options := models.AuthURLOptions{}
	authURL, err := GetAuthorizationURL(options)

	assert.Equal(t, "", authURL, "The URL should be an empty string")
	assert.Error(t, err, "An error should be returned because the required fields are empty.")
	assert.EqualError(t, err, "Key: 'AuthURLOptions.ClientID' Error:Field validation for 'ClientID' failed on the 'required' tag\nKey: 'AuthURLOptions.RedirectURL' Error:Field validation for 'RedirectURL' failed on the 'required' tag")
}

func TestGetAuthorizationURLWithoutRedirectURL(t *testing.T) {
	options := models.AuthURLOptions{
		ClientID: "randomClientID",
	}
	authURL, err := GetAuthorizationURL(options)

	assert.Equal(t, "", authURL, "The URL should be an empty string")
	assert.Error(t, err, "An error should be returned because the required fields are empty.")
	assert.EqualError(t, err, "Key: 'AuthURLOptions.RedirectURL' Error:Field validation for 'RedirectURL' failed on the 'required' tag")
}

func TestGetAuthorizationURL(t *testing.T) {
	options := models.AuthURLOptions{
		ClientID:    "randomClientID",
		RedirectURL: "https://example.com",
	}
	authURL, err := GetAuthorizationURL(options)

	assert.NoError(t, err)
	assert.Equal(t, "https://appleid.apple.com/auth/authorize?response_type=code&response_mode=form_post&state=state&scope=email&client_id=randomClientID&redirect_uri=https://example.com", authURL, "The URL should be a valid URL")
}

func TestGetAuthorizationToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authTokenOption := models.AuthTokenOption{
			ClientSecret: "randomSecret",
		}
		options := models.AuthURLOptions{
			ClientID:    "randomClientID",
			RedirectURL: "https://example.com",
		}
		code := "randomCode"

		mockAuthToken := new(mocks.AuthorizationTokenService)
		mockAuthToken.On("GetAuthorizationToken", code, options, authTokenOption).Return("newToken", nil)

		authorizationTokenService := AuthorizationTokenService(mockAuthToken)
		token, err := authorizationTokenService.GetAuthorizationToken(code, options, authTokenOption)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, "newToken", token)

		mockAuthToken.AssertNumberOfCalls(t, "GetAuthorizationToken", 1)
		mockAuthToken.AssertExpectations(t)
	})
	t.Run("without client id", func(t *testing.T) {
		authTokenOption := models.AuthTokenOption{
			ClientSecret: "randomSecret",
		}
		options := models.AuthURLOptions{
			//	RedirectURL: "https://example.com",
		}
		code := "randomCode"
		mockAuthToken := new(mocks.AuthorizationTokenService)
		expectedError := errors.New("client id and redirect url are required")

		mockAuthToken.On("GetAuthorizationToken", code, options, authTokenOption).Return("", expectedError)
		authorizationTokenService := AuthorizationTokenService(mockAuthToken)
		token, err := authorizationTokenService.GetAuthorizationToken(code, options, authTokenOption)
		assert.Equal(t, "", token, "the token should be empty")
		assert.Error(t, err, "An error should be returned because the required fields are empty.")
		assert.EqualError(t, err, expectedError.Error())
		mockAuthToken.AssertNumberOfCalls(t, "GetAuthorizationToken", 1)
		mockAuthToken.AssertExpectations(t)
	})
	t.Run("without client secret", func(t *testing.T) {
		authTokenOption := models.AuthTokenOption{}
		options := models.AuthURLOptions{
			ClientID:    "randomClientID",
			RedirectURL: "https://example.com",
		}
		code := "randomCode"

		mockAuthToken := new(mocks.AuthorizationTokenService)
		expectedError := errors.New("client secret is required")

		mockAuthToken.On("GetAuthorizationToken", code, options, authTokenOption).Return("", expectedError)
		authorizationTokenService := AuthorizationTokenService(mockAuthToken)
		token, err := authorizationTokenService.GetAuthorizationToken(code, options, authTokenOption)

		assert.Equal(t, "", token, "the token should be empty")
		assert.Error(t, err, "An error should be returned because the required fields are empty.")
		assert.EqualError(t, err, expectedError.Error())

		mockAuthToken.AssertNumberOfCalls(t, "GetAuthorizationToken", 1)
		mockAuthToken.AssertExpectations(t)
	})
}
