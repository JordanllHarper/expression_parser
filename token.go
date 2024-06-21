package expressionparser

import (
	"errors"
	"strconv"
	"unicode"
)

type integer_token struct{ int }
type add_operator_token struct{}
type subtract_operator_token struct{}
type token interface{}

func integer_or_error(index int, runes []rune) (token, int, error) {
	integer_accumulator := []rune{}
	for unicode.IsDigit(runes[index]) {
		integer_accumulator = append(integer_accumulator, runes[index])
		index += 1
		if index > len(runes)-1 {
			break
		}
	}
	if len(integer_accumulator) == 0 {
		return nil, index, errors.New("Not an integer")
	}

	accum_to_string := string(integer_accumulator)
	parsed_int, err := strconv.Atoi(accum_to_string)
	if err != nil {
		return nil, index, err
	}
	return integer_token{parsed_int}, index, nil

}
func operator_or_error(character rune) (token, error) {

	switch character {
	case '+':
		return add_operator_token{}, nil

	case '-':
		return subtract_operator_token{}, nil
	}

	return 0, errors.New("Not an operator")

}
