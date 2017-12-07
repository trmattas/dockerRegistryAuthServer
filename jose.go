package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"encoding/json"
	"github.com/dvsekhvalnov/jose2go/keys/ecc"
	"io/ioutil"
)

func CreatePlaintextToken() {

	claimsPayload := &ClaimSet{
		Issuer: "localhost:5000", // or 172.20.80.221:5000
		Subject: "",
		Audience: "",
		ExpirationTime: 1512170034,
		NotBefore: 1512170016,
		IssuedAt: 1512170016,
		JwtId: "asdasd",
	}
	claimsPayloadByteSlice, _ := json.Marshal(claimsPayload)

	fmt.Println(string(claimsPayloadByteSlice))
	token, err := jose.Sign(string(claimsPayloadByteSlice), jose.NONE, nil, jose.Header("typ", "JWT"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nPlaintext = %v\n", token)
}

func CreateES256Token(set *ClaimSet, header *JoseHeader, keyPath string) string {
	headers := make(map[string]interface{})
	headers["typ"] = header.Type
	headers["alg"] = header.Algo
	headers["kid"] = header.KeyId

	keyAsBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(keyAsBytes))

	privateKey, err2 := ecc.ReadPrivate(keyAsBytes)
	if err2 != nil {
		panic(err2)
	}

	//// temporary solution from the github, while not running on system with PKCS8 key in a file
	//privateKey := ecc.NewPrivate([]byte{4, 114, 29, 223, 58, 3, 191, 170, 67, 128, 229, 33, 242, 178, 157, 150, 133, 25, 209, 139, 166, 69, 55, 26, 84, 48, 169, 165, 67, 232, 98, 9},
	//	[]byte{131, 116, 8, 14, 22, 150, 18, 75, 24, 181, 159, 78, 90, 51, 71, 159, 214, 186, 250, 47, 207, 246, 142, 127, 54, 183, 72, 72, 253, 21, 88, 53},
	//	[]byte{ 42, 148, 231, 48, 225, 196, 166, 201, 23, 190, 229, 199, 20, 39, 226, 70, 209, 148, 29, 70, 125, 14, 174, 66, 9, 198, 80, 251, 95, 107, 98, 206 })

	setAsBytes, _ := json.Marshal(set) // get the payload in string form
	token, err3 := jose.Sign(string(setAsBytes), jose.ES256, privateKey, jose.Headers(headers)) //jose.Header("typ", header.Type), jose.Header("alg", header.Algo), jose.Header("kid", header.KeyId))
	if err3 != nil {
		panic(err3)
	}

	return token
}

func CreateRS256Token(set *ClaimSet, header *JoseHeader, keyPath string) string {
	headers := make(map[string]interface{})
	headers["typ"] = header.Type
	headers["alg"] = header.Algo
	headers["kid"] = header.KeyId

	keyAsBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(keyAsBytes))

	privateKey, err2 := ecc.ReadPrivate(keyAsBytes)
	if err2 != nil {
		panic(err2)
	}

	//// temporary solution from the github, while not running on system with PKCS8 key in a file
	//privateKey := ecc.NewPrivate([]byte{4, 114, 29, 223, 58, 3, 191, 170, 67, 128, 229, 33, 242, 178, 157, 150, 133, 25, 209, 139, 166, 69, 55, 26, 84, 48, 169, 165, 67, 232, 98, 9},
	//	[]byte{131, 116, 8, 14, 22, 150, 18, 75, 24, 181, 159, 78, 90, 51, 71, 159, 214, 186, 250, 47, 207, 246, 142, 127, 54, 183, 72, 72, 253, 21, 88, 53},
	//	[]byte{ 42, 148, 231, 48, 225, 196, 166, 201, 23, 190, 229, 199, 20, 39, 226, 70, 209, 148, 29, 70, 125, 14, 174, 66, 9, 198, 80, 251, 95, 107, 98, 206 })

	setAsBytes, _ := json.Marshal(set) // get the payload in string form
	token, err3 := jose.Sign(string(setAsBytes), jose.RS256, privateKey, jose.Headers(headers)) //jose.Header("typ", header.Type), jose.Header("alg", header.Algo), jose.Header("kid", header.KeyId))
	if err3 != nil {
		panic(err3)
	}

	return token
}