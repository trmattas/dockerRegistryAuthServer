package main

//import (
//	"fmt"
//	"net/url"
//)

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
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}


//func ParseQuery(values url.Values) map[string]string {
//	if len(values) == 0 {
//		panic("Query string was empty (length 0) -- unable to parse query parameters and create Response")
//	}
//
//	data := make(map[string]string, 0)
//	data[GRANT_TYPE] = values.Get(GRANT_TYPE)
//	data[]
//
//
//}