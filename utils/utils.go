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

// IsURL takes the provided string ("str") and checks if it is a URL
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
