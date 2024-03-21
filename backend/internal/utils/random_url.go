package utils

import "math/rand"

var randomRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomURL(length int) string {
	str := make([]rune, length)

	for i := range str {
		str[i] = randomRunes[rand.Intn(len(randomRunes))]
	}

	return string(str)
}
