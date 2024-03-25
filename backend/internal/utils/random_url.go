package utils

import "math/rand"

var randomRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomURL(length ...int) string {
	if len(length) <= 0 || len(length) > 1 {
		length = []int{8}
	}

	str := make([]rune, length[0])

	for i := range str {
		str[i] = randomRunes[rand.Intn(len(randomRunes))]
	}

	return string(str)
}
