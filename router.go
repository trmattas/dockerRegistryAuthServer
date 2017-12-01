package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewAuthRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range authroutes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
