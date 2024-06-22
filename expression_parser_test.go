package expressionparser

import (
	"testing"
)

func TestTokenizingASingleNum(t *testing.T) {
	data := "2"
	want := IntegerToken{2}
	actual := Tokenize(data)

	if _, ok := actual[0].(IntegerToken); !ok {
		t.Fatal("Token was not an integer token")
	}

	if want != actual[0] {
		t.Fatal("Wanted token != actual token")

	}
}

func TestTokenizingAMultiDigit(t *testing.T) {
	data := "222"
	want := IntegerToken{222}
	actual := Tokenize(data)

	if _, ok := actual[0].(IntegerToken); !ok {
		t.Fatal("Token was not an integer token")
	}

	if want != actual[0] {
		t.Fatal("Wanted token != actual token")

	}
}

func TestTokenizingASingleOperator(t *testing.T) {
	data := "+"
	actual := Tokenize(data)

	if _, ok := actual[0].(AddOperatorToken); !ok {
		t.Fatal("Token was not an addition token")
	}
}

func TestTokenizingBasicAddition(t *testing.T) {
	data := "2 + 2"
	expected := []Token{IntegerToken{2}, AddOperatorToken{}, IntegerToken{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}

}

func TestTokenizingBasicSubtraction(t *testing.T) {
	data := "2 - 2"
	expected := []Token{IntegerToken{2}, SubtractOperatorToken{}, IntegerToken{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}
}

func TestTokenizingBasicMultiplication(t *testing.T) {
	data := "2 * 2"
	expected := []Token{IntegerToken{2}, MultiplyOperatorToken{}, IntegerToken{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}
}

func TestTokenizingBasicDivision(t *testing.T) {
	data := "2 / 2"
	expected := []Token{IntegerToken{2}, DivideOperatorToken{}, IntegerToken{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}
}

func TestTokenizingBasicDivisionWithParenthesis(t *testing.T) {
	data := "(2 / 2)"
	expected := []Token{ParenOpeningToken{}, IntegerToken{2}, DivideOperatorToken{}, IntegerToken{2}, ParenClosingToken{}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}
}
