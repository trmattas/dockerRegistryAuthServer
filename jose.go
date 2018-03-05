package main

import (
	"encoding/json"
	//"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

func CreateRS256Token(set *ClaimSet, header *JoseHeader, keyPath string) string {
	headers := make(map[string]interface{})
	headers["typ"] = header.Type
	headers["alg"] = header.Algo
	headers["kid"] = header.KeyId

	keyAsBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	privateKey, err2 := Rsa.ReadPrivate(keyAsBytes)
	if err2 != nil {
		panic(err2)
	}

	setAsBytes, _ := json.Marshal(set)                                                          // get the payload in string form
	token, err3 := jose.Sign(string(setAsBytes), jose.RS256, privateKey, jose.Headers(headers)) //jose.Header("typ", header.Type), jose.Header("alg", header.Algo), jose.Header("kid", header.KeyId))
	if err3 != nil {
		panic(err3)
	}

	return token
}
