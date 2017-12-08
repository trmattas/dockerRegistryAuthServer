package main

import (
	"math/rand"
	"time"
)

var alphanumerics = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") // all alphanumerics

func RandomString(length int) string {
	bytes := make([]rune, length) // make a byte slice of the desired length of our random string
	rand.Seed(time.Now().UnixNano()) // seed our rand # generator -- don't need to keep an object; seed stays in memory
	for i := range bytes {
		bytes[i] = alphanumerics[rand.Intn(len(alphanumerics))] // choose a random member of the slice
	}
	return string(bytes)
}
