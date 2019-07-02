package practice

import "fmt"

type node struct {
	value int
	left, right *node
}

func (n *node) print() {
	fmt.Print(n.value)
}

func (n *node) traverse() {
	n.traverseFunc(func (n *node) {
		n.print()
	})
}

func (n *node) traverseFunc(f func(*node)) {
	if n == nil {
		return
	}

	n.left.traverseFunc(f)
	f(n)
	n.right.traverseFunc(f)
}


