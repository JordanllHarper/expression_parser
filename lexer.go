package main

func Tokenize(raw_expression string) []Token {
	runes := []rune(raw_expression)
	tokens := []Token{}

	for i := 0; i < len(runes); {
		token, new_index, err := IntegerOrError(i, runes)
		if err == nil {
			tokens = append(tokens, token)
			i = new_index
			continue
		}

		token, err = OperatorOrError(runes[i])
		if err == nil {
			tokens = append(tokens, token)
		}

		token, err = ParenthesisOrError(runes[i])
		if err == nil {
			tokens = append(tokens, token)
		}
		i++
	}
	return tokens
}
