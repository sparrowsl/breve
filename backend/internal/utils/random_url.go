package utils

import "math/rand"

var randomRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomURL(size int) string {
	str := make([]rune, size)

	for i := range str {
		str[i] = randomRunes[rand.Intn(len(randomRunes))]
	}

	return string(str)
}
