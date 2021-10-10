package goapplesignin

import "github.com/BolajiOlajide/go-apple-signin/constants"

func normalizeAuthOptions(options *AuthURLOptions) *AuthURLOptions {
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
