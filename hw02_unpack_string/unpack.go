package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(in string) (string, error) {
	if len(in) == 0 {
		return "", nil
	}

	var result strings.Builder
	i := 0
	step := 2

	for i < len(in)-1 {
		_, err := strconv.Atoi(string(in[i]))
		if err == nil {
			// first symbol is digit
			return "", ErrInvalidString
		}
		// symbol is literal
		digit, err := strconv.Atoi(string(in[i+1]))
		if err != nil {
			digit = 1
			step = 1
		}
		result.WriteString(strings.Repeat(string(in[i]), digit))
		i += step
		step = 2
	}
	if i < len(in) {
		result.WriteString(strings.Repeat(string(in[len(in)-1]), 1))
	}

	return result.String(), nil
}
