package main

import "net/http"

// Custom struct for the Route, allowing us to specify METHOD
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Slice of "Route"(s) are represented as "Routes"
type Routes []Route

// Some Routes as an example
// These routes are for the authorization server
var authroutes = Routes{
	Route{
		"Index",
		"GET",
		"/index",
		Index,
	},
	Route{
		"GetToken",
		"GET",
		"/token",
		AuthToken,
	},/*
	Route{
		"Authn",
		"GET",
		"/authn",
		AuthN,
	},*/
	Route{
		"PostToken",
		"POST",
		"/token",
		AuthToken,
	},
}
