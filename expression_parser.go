package expressionparser

import (
	"fmt"
)

func main() {
	my_expression := "22 + 222 + 22222 - 3"
	tokens := Tokenize(my_expression)
	for _, token := range tokens {
		v, ok := token.(IntegerToken)
		if ok {
			println(v.int)
		}
		_, ok = token.(AddOperatorToken)
		if ok {
			fmt.Printf("%c\n", '+')
		}

		_, ok = token.(SubtractOperatorToken)
		if ok {
			fmt.Printf("%c\n", '-')
		}

	}
}
