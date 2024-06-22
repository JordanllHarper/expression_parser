package main

type Node interface {
}

type ValueNode struct{ Token }

type BranchingNode struct {
	children []Node
}
