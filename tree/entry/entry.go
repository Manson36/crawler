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

	fmt.Print("In-order traversal: ")
	root.Traverse()

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxNodeValue)
}
