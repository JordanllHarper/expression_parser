package expressionparser

type Node interface{}

type ValueNode struct{ Token }

type BranchingNode struct {
	token    Token
	children []Node
}
