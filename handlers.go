package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

// can remove
func Auth(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Query().Get("service"))

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode("blahblah")
	if err != nil {
		panic(err)
	}
}

// can remove
func AuthN(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.DumpRequest(r, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

// Handles the POST request asking for an Oauth2 bearer token in JWT format
func AuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Now, we create the JWT from the info in the query parameters
	fmt.Println("Path: ",r.URL.Path)
	fmt.Println("RawQuery: ",r.URL.RawQuery)


	// Extract parameters from encoded query parameter set
	values := r.URL.Query()
	if len(values) == 0 {
		panic("Query string was empty (length 0) -- unable to parse query parameters and create Response")
	}

	var grantType, service, clientId, accessType, rawScope, refreshToken, username, password string

	grantType = values.Get(GRANT_TYPE)
	service = values.Get(SERVICE)
	clientId = values.Get(CLIENT_ID)
	accessType = values.Get(ACCESS_TYPE)
	rawScope = values.Get(SCOPE)
	refreshToken = values.Get(REFRESH_TOKEN)
	username = values.Get(USERNAME)
	password = values.Get(PASSWORD)

	fmt.Println(grantType)
	fmt.Println(service)
	fmt.Println(clientId)
	fmt.Println(accessType)
	fmt.Println(rawScope)
	fmt.Println(refreshToken)
	fmt.Println(username)
	fmt.Println(password)


	//body, err := ioutil.ReadAll(io.Reader(r.Body))
	//if err != nil {
	//	panic(err)
	//}
	//defer r.Body.Close()

	//var tokenRequest *JsonWebToken
	//if err2 := json.Unmarshal(body, &tokenRequest); err2 != nil {
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.StatusUnprocessableEntity)
	//	fmt.Println(err)
	//
	//	//if err := json.NewEncoder(w).Encode(err); err != nil {
	//	//	panic(err)
	//	//}
	//	return
	//}

	//if err := json.NewEncoder(w).Encode(contentPayload); err != nil {
	//
	//}
}
