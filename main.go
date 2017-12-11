package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())

	router := NewAuthRouter()

	log.Fatal(http.ListenAndServeTLS(":7000", "/root/certs/domain.crt", "/root/certs/domain.key", router))
}
