package goapplesignin

import (
	"fmt"
	"net/url"

	"github.com/BolajiOlajide/go-apple-signin/constants"
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// GetAuthorizationURL returns an initiating auth for apple users
func GetAuthorizationURL(options AuthURLOptions) (string, error) {
	err := validate.Struct(options)
	if err != nil {
		return "", fmt.Errorf("One of the options doesn't meet the requirement. - %v", err)
	}

	parsedURL, err := url.Parse(constants.AppleEndpointURL)
	if err != nil {
		// ideally we would never get here since the endpoint url is a hardcoded constant
		return "", fmt.Errorf("Cannot parse Apple Base URL. %v", err)
	}

	normalizeAuthOptions(&options)

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