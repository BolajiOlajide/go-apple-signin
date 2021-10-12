package models

// AuthURLOptions argument used to generate auth url for apple signin
type AuthURLOptions struct {
	ClientID     string `validate:"required"`
	RedirectURL  string `validate:"required"`
	Scope        string
	ResponseType string
	ResponseMode string
	State        string
}

type AuthTokenOption struct {
	ClientID     string `validate:"required"`
	RedirectURL  string `validate:"required"`
	ClientSecret string `validate:"required"`
	Scope        string
	ResponseType string
	ResponseMode string
	State        string
}
