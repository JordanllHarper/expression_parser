package expressionparser

import (
	"errors"
	"strconv"
	"unicode"
)

type AddOperatorToken struct{}
type SubtractOperatorToken struct{}

type MultiplyOperatorToken struct{}
type DivideOperatorToken struct{}

type ParenOpeningToken struct{}
type ParenClosingToken struct{}

type integer_token struct{ int }
type token interface{}

func integer_or_error(index int, runes []rune) (token, int, error) {
	integer_accumulator := []rune{}
	for unicode.IsDigit(runes[index]) {
		integer_accumulator = append(integer_accumulator, runes[index])
		index += 1
		if OutOfRange(runes, index) {
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
func parenthesis_or_error(character rune) (token, error) {
	switch character {
	case '(':
		return ParenOpeningToken{}, nil

	case ')':
		return ParenClosingToken{}, nil

	}
	return nil, errors.New("Not a parenthesis")

}
func operator_or_error(character rune) (token, error) {

	switch character {
	case '+':
		return AddOperatorToken{}, nil

	case '-':
		return SubtractOperatorToken{}, nil

	case '*':
		return MultiplyOperatorToken{}, nil

	case '/':
		return DivideOperatorToken{}, nil
	}

	return 0, errors.New("Not an operator")

}
