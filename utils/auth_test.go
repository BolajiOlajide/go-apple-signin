package utils

import (
	"testing"

	"github.com/BolajiOlajide/go-apple-signin/constants"
	"github.com/BolajiOlajide/go-apple-signin/models"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeAuthOptionsWithDefaultValues(t *testing.T) {
	options := models.AuthURLOptions{}
	normalizedOpts := NormalizeAuthOptions(&options)

	assert.Equal(t, normalizedOpts.ResponseType, constants.AppleDefaultResponseType)
	assert.Equal(t, normalizedOpts.ResponseMode, constants.AppleDefaultResponseMode)
	assert.Equal(t, normalizedOpts.Scope, constants.AppleDefaultScope)
	assert.Equal(t, normalizedOpts.State, constants.AppleDefaultState)

	assert.Equal(t, normalizedOpts.ClientID, options.ClientID, "ClientID shouldn't be overwritten")
	assert.Equal(t, normalizedOpts.RedirectURL, options.RedirectURL, "RedirectURL shouldn't be overwritter")
}

func TestNormalizeAuthOptionsWithoutDefaultValues(t *testing.T) {
	options := models.AuthURLOptions{
		ClientID:     "100234",
		RedirectURL:  "http://example.com",
		Scope:        "email,name",
		State:        "randomStateValue",
		ResponseType: "refresh_token",
		ResponseMode: "json",
	}
	normalizedOpts := NormalizeAuthOptions(&options)

	assert.Equal(t, normalizedOpts.ResponseType, options.ResponseType)
	assert.Equal(t, normalizedOpts.ResponseMode, options.ResponseMode)
	assert.Equal(t, normalizedOpts.Scope, options.Scope)
	assert.Equal(t, normalizedOpts.State, options.State)

	assert.Equal(t, normalizedOpts.ClientID, options.ClientID, "ClientID shouldn't be overwritten")
	assert.Equal(t, normalizedOpts.RedirectURL, options.RedirectURL, "RedirectURL shouldn't be overwritter")
}
