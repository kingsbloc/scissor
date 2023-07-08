package utils

import (
	"errors"
	"math"
	"strings"
)

const (
	alphabets     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	biggestuint64 = 18446744073709551615
)

func EncodeBase62(number uint64) string {
	length := len(alphabets)
	var encBuilder strings.Builder
	encBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encBuilder.WriteByte(alphabets[(number % uint64(length))])
	}
	return encBuilder.String()
}

func DecodeBase62(encodedString string) (uint64, error) {
	var number uint64
	length := len(alphabets)

	for i, symbol := range encodedString {
		alphabetPosition := strings.IndexRune(alphabets, symbol)
		if alphabetPosition == -1 {
			return uint64(alphabetPosition), errors.New("cannot find symbol in alphabet")
		}
		number += uint64(alphabetPosition) * uint64(math.Pow(float64(length), float64(i)))
	}
	return number, nil
}
