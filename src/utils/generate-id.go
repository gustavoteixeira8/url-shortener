package utils

import "math/rand"

const (
	_LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_NUMBERS = "0123456789"
)

func NewID() string {
	b := make([]byte, 6)

	for i := range b {
		b[i] = _LETTERS[rand.Intn(len(_LETTERS))]
		b[rand.Intn(6)] = _NUMBERS[rand.Intn(len(_NUMBERS))]
	}

	return string(b)
}
