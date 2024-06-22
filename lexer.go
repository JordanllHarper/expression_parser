package expressionparser

func Tokenize(raw_expression string) []Token {
	runes := []rune(raw_expression)
	tokens := []Token{}

	for i := 0; i < len(runes); {
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

		token, err = parenthesis_or_error(runes[i])
		if err == nil {
			tokens = append(tokens, token)
		}
		i++
	}
	return tokens
}
