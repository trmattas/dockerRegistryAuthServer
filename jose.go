package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
)


func CreatePlaintextToken() {

	payload := `{"hello": "world"}`

	token, err := jose.Sign(payload, jose.NONE, nil)

	if (err != nil) {
		panic(err)
	}

	fmt.Printf("\nPlaintext = %v\n", token)
}