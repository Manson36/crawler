package tree

import "fmt"

func (node *Node) Traverse() {
	node.Traversefunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) Traversefunc(f func( *Node)) {
	if node == nil {
		return
	}

	node.Left.Traversefunc(f)
	f(node)
	node.Right.Traversefunc(f)
}