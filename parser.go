package main

import ()

func ShouldCloseBranch(token Token) bool {
	_, isClosing := token.(ParenClosingToken)
	return isClosing
}
func ShouldBranch(token Token) bool {
	_, isOpening := token.(ParenOpeningToken)
	return isOpening
}

// 2 + (3 - 4)
// start -> branch (root)
//	// 2
//  // +
//  // open paren -> branch
//  //  // (
//  //  // 3
//  //  // -
//  //  // 4
//  //  // )

func BuildNode(index int, tokens []Token, parent BranchingNode) (Node, error) {
	if index > len(tokens)-1 {
		return parent, nil
	}
	current := tokens[index]
	println(current.PrettyString())

	if ShouldBranch(current) {
		firstNodeOfBranch := ValueNode{current}
		branch := BranchingNode{children: []Node{firstNodeOfBranch}}
		result, err := BuildNode(index+1, tokens, branch)
		if err != nil {
			return nil, err
		}
		parent.children = append(parent.children, result)
		return parent, nil
	} else {
		value := ValueNode{current}
		parent.children = append(parent.children, value)
		parent, err := BuildNode(index+1, tokens, parent)
		if err != nil {
			return nil, err
		} else {
			return parent, nil
		}
	}
}

func ParseOrError(tokens []Token) (Node, error) {

	root := BranchingNode{children: []Node{}}

	node, err := BuildNode(0, tokens, root)
	if err != nil {
		return nil, err
	}

	return node, nil

}
func PrintNodes(root Node) {
	if value, ok := root.(BranchingNode); ok {
		println("")
		println("Branch:")
		for _, node := range value.children {
			if branch, ok := node.(BranchingNode); ok {
				PrintNodes(branch)
			} else {
				value, ok := node.(ValueNode)
				if ok {
					print(value.Token.PrettyString())
				}
			}
		}
	} else {
		valueNode, ok := root.(ValueNode)
		if ok {
			print(valueNode.Token.PrettyString())
		}
	}

}
