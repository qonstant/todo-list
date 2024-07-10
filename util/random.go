package util

import (
	"math/rand"
	"time"
)

// RandomString generates a random string of a given length
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// RandomTime generates a random time
func RandomTime() time.Time {
	return time.Now().Add(time.Duration(rand.Intn(10000)) * time.Hour)
}

// RandomBool generates a random boolean value
func RandomBool() bool {
	return rand.Intn(2) == 1
}
