package constants

// AppleEndpointURL base URL for apple authentication
const AppleEndpointURL = "https://appleid.apple.com"

// TokenIssuer issuer identifier for apple authentication
const TokenIssuer = "https://appleid.apple.com"

// default variables

// AppleDefaultScope the default scope for apple authentication, it should be comma separated if > 1
const AppleDefaultScope = "email"

// AppleDefaultState the default state variable for apple authentication
const AppleDefaultState = "state"

// AppleDefaultResponseType the default response type for apple authentication
const AppleDefaultResponseType = "code"

// AppleDefaultResponseMode the default response mode for apple authentication
const AppleDefaultResponseMode = "form_post"
