package expressionparser

import "errors"

func IsBranchingNode(token Token) bool {
	if _, ok := token.(IntegerToken); ok {
		return false
	}
	return true

}

func BuildAppropriateNode(index int, tokens []Token) (Node, error) {
	if index > len(tokens)-1 {
		return nil, errors.New("Somehow, we got out of range")
	}
	current := tokens[index]
	if IsBranchingNode(tokens[index]) {
		node := BuildAppropriateNode(index+1, tokens)

	}
	return ValueNode{token}, nil

}

func ParseOrError(tokens []Token) (Node, error) {

	node, err := BuildAppropriateNode(0, tokens)
	if err != nil {
		return nil, err
	}

	return node, nil

}
