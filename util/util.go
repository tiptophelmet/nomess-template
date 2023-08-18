package util

import (
	"fmt"
	"math/rand"
)

var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func IsEmpty(i interface{}) bool {
	return i == nil || (i != nil && fmt.Sprintf("%v", i) == "")
}
