package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())

	router := NewAuthRouter()

	log.Fatal(http.ListenAndServe(":7000", router))
}
