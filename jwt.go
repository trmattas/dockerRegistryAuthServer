package main

type JsonWebToken struct {
	Token    string `json:"token"`
	header   JoseHeader
	claimSet ClaimsPayload
}

type JoseHeader struct {
	Type        StringOrUri `json:"typ"`
	ContentType StringOrUri `json:"cty"`
}

type ClaimsPayload struct {
	Issuer         StringOrUri `json:"iss"`
	Subject        StringOrUri `json:"sub"`
	Audience       StringOrUri `json:"aud"`
	ExpirationTime NumericDate `json:"exp"`
	NotBefore      NumericDate `json:"nbf"`
	IssuedAt       NumericDate `json:"iat"`
	JwtId          string      `json:"jti"`
}

type StringOrUri string

type NumericDate uint64