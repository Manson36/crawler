package main

import (
	"fmt"
	"github.com/crawler/tree"
)

type myTreeNode struct {
	node *tree.Node
}

//后序遍历
func (mynode *myTreeNode) postOrder() {
	if mynode == nil || mynode.node == nil {
		return
	}

	left := myTreeNode{mynode.node.Left}
	right := myTreeNode{mynode.node.Right}

	left.postOrder()
	right.postOrder()
	mynode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	Myroot := myTreeNode{&root}
	Myroot.postOrder()
	fmt.Println()
}
