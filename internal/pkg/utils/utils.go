package utils

import (
	"math/rand"
	"time"
)

const characters = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func GenerateRandomChars(length int) (res string) {
	maxln := len(characters) - 1
	for range length {
		index := RandomInt(0, int64(maxln))
		res += string(characters[index])
	}
	return res
}

func GenerateRandomPassword() string {
	return GenerateRandomChars(15)
}
