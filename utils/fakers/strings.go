package fakers

import (
	"math/rand"
	"time"
)

const (
	Numbers       = "0987654321"
	LowerCaseChar = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune(LowerCaseChar + UpperCaseChar + Numbers)
	result := make([]rune, length)

	for index := 0; index < length; index++ {
		result[index] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
