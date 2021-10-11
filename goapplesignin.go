package goapplesignin

import (
	"fmt"
	"net/url"

	"github.com/BolajiOlajide/go-apple-signin/constants"
	"github.com/BolajiOlajide/go-apple-signin/models"
	"github.com/BolajiOlajide/go-apple-signin/utils"
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate = validator.New()

// GetAuthorizationURL returns an initiating auth for apple users
func GetAuthorizationURL(options models.AuthURLOptions) (string, error) {
	err := validate.Struct(&options)
	if err != nil {
		return "", err
	}

	parsedURL, err := url.Parse(constants.AppleEndpointURL)
	if err != nil {
		// ideally we would never get here since the endpoint url is a hardcoded constant
		return "", fmt.Errorf("Cannot parse Apple Base URL. %v", err)
	}

	utils.NormalizeAuthOptions(&options)

	parsedURL.Path = "/auth/authorize"
	parsedURL.ForceQuery = true

	parsedURL.RawQuery = fmt.Sprintf(
		"response_type=%s&response_mode=%s&state=%s&scope=%s&client_id=%s&redirect_uri=%s",
		options.ResponseType,
		options.ResponseMode,
		options.State,
		options.Scope,
		options.ClientID,
		options.RedirectURL,
	)

	return parsedURL.String(), nil
}
