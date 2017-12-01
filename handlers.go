package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Query().Get("service"))

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode("blahblah")
	if err != nil {
		panic(err)
	}
}

func AuthN(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.DumpRequest(r, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func RequestAuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
