package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())

	// Read DER encoded key into memory from file
	derKeyAsBytes, err0 := ioutil.ReadFile("pkcs8.der")
	if err0 != nil {
		panic(err0)
	}
	fmt.Println(string(derKeyAsBytes)) //debugging

	kid := CreateKidFromDer(derKeyAsBytes)
	fmt.Println(kid) // debugging

	payload := &ClaimSet{
		Issuer:         "localhost:5000", // or 172.20.80.221:5000
		Subject:        "",
		Audience:       "",
		ExpirationTime: 1512170034,
		NotBefore:      1512170016,
		IssuedAt:       1512170016,
		JwtId:          "asdasd",
	}

	//header := &JoseHeader{ // debugging -- can delete
	//	Type: "JWT",
	//	Algo: "ES256",
	//	KeyId: "PYYO:TEWU:V7JH:26JV:AQTZ:LJC3:SXVJ:XGHA:34F2:2LAQ:ZRMK:Z7Q6", // from https://docs.docker.com/registry/spec/auth/jwt/#getting-a-bearer-token
	//}
	header := &JoseHeader{
		Type:  "JWT",
		Algo:  "RS256",
		KeyId: kid,
	}
	token := CreateRS256Token(payload, header, "pkcs8.pem")
	fmt.Println(token)

	router := NewAuthRouter()
	go func() {
		log.Fatal(http.ListenAndServe(":7000", router))
	}()

	for {
		a := 1
		a = a + 1
	}
}
