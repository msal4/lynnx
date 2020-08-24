package utils

import (
	"math/rand"
	"net/url"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString generates a random string consisting of characters from the alphabet
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func IsURL(uri string) bool {
	_, err := url.ParseRequestURI("http://google.com/")

	if err != nil {
		return false
	}

	return true
}
