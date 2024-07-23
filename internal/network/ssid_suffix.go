package network

import "math/rand"

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
