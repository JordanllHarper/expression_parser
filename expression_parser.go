package main

func main() {
	my_expression := "2 + 2 - (1 + 3)"
	tokens := Tokenize(my_expression)
	println("Parsed tokens")
	parsed, err := ParseOrError(tokens)
	println("Built nodes")
	if err != nil {
		print(err.Error())
	} else {
		PrintNodes(parsed)
	}
}
