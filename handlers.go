package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, user. This is the auth server for a Docker registry.")
}

// Handles the POST request asking for an Oauth2 bearer token in JWT format
func AuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Remote Host: ", r.RemoteAddr)

	// Now, we create the JWT from the info in the query parameters

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
	// ToDo -- Docker Registry Auth only (for now) uses the Password Authorization Grant Type
	// TODO need to refactor
	// see https://docs.docker.com/registry/spec/auth/oauth/
	//username = values.Get(USERNAME)
	//password = values.Get(PASSWORD)


	// debugging
	fmt.Println(grantType)
	fmt.Println(service)
	fmt.Println(clientId)
	fmt.Println(accessType)
	fmt.Println(rawScope)
	fmt.Println(refreshToken)

	// TODO this is where we would be handling AuthZ (make sure this user with this action in our ACL)

	fmt.Println("-----------------------------------")

	// Handle creating the claim set
	var scope []ScopeAccess
	claimSet := &ClaimSet{
		Issuer:         "auth-server", // the auth server -- this string has to directly match what is in the config file
		Subject:        "",
		Audience:       r.RemoteAddr,                                                   // the docker registry address
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
	derKeyAsBytes, err0 := ioutil.ReadFile("/root/go/src/dockerRegistryAuthServer/pkcs8_1024.der")
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
	token := CreateRS256Token(claimSet, header, "/root/go/src/dockerRegistryAuthServer/pkcs8_1024.pem")

	// pack the token into the right header
	response := Response{
		AccessToken: token,
		ExpiresIn:   600, // 600 seconds = 10 minutes
		Scope:       ScopeToResponse(scope),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
