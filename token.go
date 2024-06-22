package main

import (
	"errors"
	"strconv"
	"unicode"
)

// Token types
type IntegerToken struct{ int }

func (t IntegerToken) PrettyString() string {
	return strconv.FormatInt(int64(t.int), 10)
}

type AddOperatorToken struct{}

func (t AddOperatorToken) PrettyString() string {
	return "+"
}

type SubtractOperatorToken struct{}

func (t SubtractOperatorToken) PrettyString() string {
	return "-"
}

type MultiplyOperatorToken struct{}

func (t MultiplyOperatorToken) PrettyString() string {
	return "*"
}

type DivideOperatorToken struct{}

func (t DivideOperatorToken) PrettyString() string {
	return "/"
}

type ParenOpeningToken struct{}

func (t ParenOpeningToken) PrettyString() string {
	return "("
}

type ParenClosingToken struct{}

func (t ParenClosingToken) PrettyString() string {
	return ")"
}

// Base Token
type Token interface {
	PrettyString() string
}

func IntegerOrError(index int, runes []rune) (Token, int, error) {
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
	return IntegerToken{parsed_int}, index, nil
}
func ParenthesisOrError(character rune) (Token, error) {
	switch character {
	case '(':
		return ParenOpeningToken{}, nil
	case ')':
		return ParenClosingToken{}, nil

	}
	return nil, errors.New("Not a parenthesis")

}
func OperatorOrError(character rune) (Token, error) {

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

	var t Token

	return t, errors.New("Not an operator")

}
