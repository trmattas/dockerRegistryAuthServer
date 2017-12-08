package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

/*
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
*/

// Handles the POST request asking for an Oauth2 bearer token in JWT format
func AuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Now, we create the JWT from the info in the query parameters
	//fmt.Println("Path: ", r.URL.Path)
	//fmt.Println("RawQuery: ", r.URL.RawQuery)

	// Extract parameters from encoded query parameter set
	values := r.URL.Query()
	if len(values) == 0 {
		panic("Query string was empty (length 0) -- unable to parse query parameters and create Response")
	}

	var grantType, service, clientId, accessType, rawScope, refreshToken /*, username, password*/ string

	grantType = values.Get(GRANT_TYPE)
	service = values.Get(SERVICE)
	clientId = values.Get(CLIENT_ID)
	accessType = values.Get(ACCESS_TYPE)
	rawScope = values.Get(SCOPE)
	refreshToken = values.Get(REFRESH_TOKEN)
	//username = values.Get(USERNAME)
	//password = values.Get(PASSWORD)

	fmt.Println(grantType)
	fmt.Println(service)
	fmt.Println(clientId)
	fmt.Println(accessType)
	fmt.Println(rawScope)
	fmt.Println(refreshToken)
	//fmt.Println(username)
	//fmt.Println(password)

	fmt.Println("-----------------------------------")

	// Handle creating the claim set
	var scope []ScopeAccess
	claimSet := &ClaimSet{
		Issuer:         "localhost:7000", // or 172.20.80.221:7000 -- the auth server
		Subject:        "",
		Audience:       "localhost:5000",                                               // the docker registry address
		ExpirationTime: uint64(time.Now().Add(time.Minute * time.Duration(10)).Unix()), // always now + 10 minutes time
		NotBefore:      uint64(time.Now().Unix()),
		IssuedAt:       uint64(time.Now().Unix()),
		JwtId:          RandomString(15),
	}
	// parse the access scope, and insert it into the claim set as needed
	scope, err := ParseScope(rawScope)
	if err != nil {
		panic(err)
	}
	if len(scope) == 0 {
		// make sure it is a "" that gets encoded here, in the "access" part, since no scope was defined
		//claimSet.EmptyAccess = ""
	} else {
		claimSet.Access = scope
	}

	// Make sure our weird use of Json is working
	jsonClaimSet, jsonErr := json.Marshal(claimSet)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println(string(jsonClaimSet))

	// Create the "kid", from the DER encoded key
	derKeyAsBytes, err0 := ioutil.ReadFile("pkcs8_1024.der")
	if err0 != nil {
		panic(err0)
	}
	//fmt.Println(string(derKeyAsBytes)) //debugging
	kid := CreateKidFromDer(derKeyAsBytes)

	// Create the header, using the kid
	header := &JoseHeader{
		Type:  "JWT",
		Algo:  "RS256",
		KeyId: kid,
	}

	// Create the actual JWT, using the PEM encoded key, as well as claimset and header
	token := CreateRS256Token(claimSet, header, "pkcs8_1024.pem")

	// pack the token into the right header
	//fmt.Println(token) // debugging

	response := Response{
		AccessToken: token,
		ExpiresIn:   600, // 600 seconds = 10 minutes
		Scope:       ScopeToResponse(scope),
	}
	// return the right JSON

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
