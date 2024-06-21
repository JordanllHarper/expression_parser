package expressionparser

import (
	"testing"
)

func TestTokenizingASingleNum(t *testing.T) {
	data := "2"
	want := integer_token{2}
	actual := Tokenize(data)

	if _, ok := actual[0].(integer_token); !ok {
		t.Fatal("Token was not an integer token")
	}

	if want != actual[0] {
		t.Fatal("Wanted token != actual token")

	}
}

func TestTokenizingAMultiDigit(t *testing.T) {
	data := "222"
	want := integer_token{222}
	actual := Tokenize(data)

	if _, ok := actual[0].(integer_token); !ok {
		t.Fatal("Token was not an integer token")
	}

	if want != actual[0] {
		t.Fatal("Wanted token != actual token")

	}
}

func TestTokenizingASingleOperator(t *testing.T) {
	data := "+"
	actual := Tokenize(data)

	if _, ok := actual[0].(add_operator_token); !ok {
		t.Fatal("Token was not an addition token")
	}
}

func TestTokenizingBasicAddition(t *testing.T) {
	data := "2 + 2"
	expected := []token{integer_token{2}, add_operator_token{}, integer_token{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}

}

func TestTokenizingBasicSubtraction(t *testing.T) {
	data := "2 - 2"
	expected := []token{integer_token{2}, subtract_operator_token{}, integer_token{2}}
	actual := Tokenize(data)

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("Tokens did not match")
		}
	}
}
