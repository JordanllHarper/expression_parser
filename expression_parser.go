package expressionparser

import (
	"fmt"
)

func Tokenize(raw_expression string) []token {
	runes := []rune(raw_expression)
	tokens := []token{}

	for i := 0; i < len(raw_expression); i++ {
		token, new_index, err := integer_or_error(i, runes)
		if err == nil {
			tokens = append(tokens, token)
			i = new_index
			continue
		}

		token, err = operator_or_error(runes[i])
		if err == nil {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func main() {
	my_expression := "22 + 222 + 22222 - 3"
	tokens := Tokenize(my_expression)
	for _, token := range tokens {
		v, ok := token.(integer_token)
		if ok {
			println(v.int)
		}
		_, ok = token.(add_operator_token)
		if ok {
			fmt.Printf("%c\n", '+')
		}

		_, ok = token.(subtract_operator_token)
		if ok {
			fmt.Printf("%c\n", '-')
		}

	}
}
