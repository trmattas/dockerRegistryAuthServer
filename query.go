package main

// Constants for all the possible query parameters for the query coming from the docker engine for AuthZ
const (
	GRANT_TYPE    = "grant_type"
	SERVICE       = "service"
	CLIENT_ID     = "client_id"
	ACCESS_TYPE   = "access_type"
	SCOPE         = "scope"
	REFRESH_TOKEN = "refresh_token"
	USERNAME      = "username"
	PASSWORD      = "password"
)

type Response struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
