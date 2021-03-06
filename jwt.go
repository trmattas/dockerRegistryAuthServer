package main

type JoseHeader struct {
	Type  string `json:"typ"`
	Algo  string `json:"alg"`
	KeyId string `json:"kid"`
}

// The Claim Set is a JSON struct containing these standard registered claim name fields:
type ClaimSet struct {
	Issuer         StringOrUri   `json:"iss"`    // The issuer of the token, typically the fqdn of the authorization server.
	Subject        StringOrUri   `json:"sub"`    // The subject of the token; the name or id of the client which requested it. This should be empty (`""`) if the client did not authenticate.
	Audience       string        `json:"aud"`    // The intended audience of the token; the name or id of the service which will verify the token to authorize the client/subject.
	ExpirationTime uint64        `json:"exp"`    // The token should only be considered valid up to this specified date and time.
	NotBefore      uint64        `json:"nbf"`    // The token should not be considered valid before this specified date and time.
	IssuedAt       uint64        `json:"iat"`    // Specifies the date and time which the Authorization server generated this token.
	JwtId          string        `json:"jti"`    // A unique identifier for this token. Can be used by the intended audience to prevent replays of the token.
	Access         []ScopeAccess `json:"access"` // the scope -- see scope.go
	//EmptyAccess    string        `json:"access"` // only needed if access is empty -- it's just a ""
}

// Docker Registry unique array (not part of usual Claim Set)
type ScopeAccess struct {
	Type    string         `json:"type"`    // always "repository" for us, but there are actually 3 possible values here -- "registry", "repository", and "repository(plugin)"
	Name    string         `json:"name"`    // name of image repo to push to or pull from
	Actions ActionsAllowed `json:"actions"` // allowed actions on that repo
}

// "Push", "Pull", both, or neither -- a slice of size 0-2, inclusive
type ActionsAllowed struct {
	Allowed []string
}

// return this as the payload of the response to our Docker client
type EncodedToken struct {
	Token string `json:"token"`
}

type StringOrUri string
