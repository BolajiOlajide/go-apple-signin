package goapplesignin

import (
	"testing"

	"github.com/BolajiOlajide/go-apple-signin/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationURLWithoutClientID(t *testing.T) {
	options := models.AuthURLOptions{}
	url, err := GetAuthorizationURL(options)

	assert.Equal(t, "", url, "The URL should be an empty string")
	assert.Error(t, err)
}
