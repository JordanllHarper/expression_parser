package expressionparser

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