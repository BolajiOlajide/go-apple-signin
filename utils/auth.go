package utils

import (
	"github.com/BolajiOlajide/go-apple-signin/constants"
	"github.com/BolajiOlajide/go-apple-signin/models"
)

// NormalizeAuthOptions fills empty struct values with corresponding default value
func NormalizeAuthOptions(options *models.AuthURLOptions) *models.AuthURLOptions {
	if options.ResponseType == "" {
		options.ResponseType = constants.AppleDefaultResponseType
	}

	if options.ResponseMode == "" {
		options.ResponseMode = constants.AppleDefaultResponseMode
	}

	if options.Scope == "" {
		options.Scope = constants.AppleDefaultScope
	}

	if options.State == "" {
		options.State = constants.AppleDefaultState
	}

	return options
}
func NormalizeAuthTokenOptions(options *models.AuthTokenOption) *models.AuthTokenOption {
	if options.ResponseType == "" {
		options.ResponseType = constants.AppleDefaultResponseType
	}

	if options.ResponseMode == "" {
		options.ResponseMode = constants.AppleDefaultResponseMode
	}

	if options.Scope == "" {
		options.Scope = constants.AppleDefaultScope
	}

	if options.State == "" {
		options.State = constants.AppleDefaultState
	}

	return options
}
