package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	nodedata int
	left     *BinaryNode
	right    *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(element int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{nodedata: element, left: nil, right: nil}
	} else {
		t.root.insert(element)
	}
	return t
}

func (n *BinaryNode) insert(element int) {
	if n == nil {
		return
	} else if element <= n.nodedata {
		if n.left == nil {
			n.left = &BinaryNode{nodedata: element, left: nil, right: nil}
		} else {
			n.left.insert(element)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{nodedata: element, left: nil, right: nil}
		} else {
			n.right.insert(element)
		}
	}

}
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.nodedata)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(1).
		insert(-2).
		insert(3).
		insert(-4).
		insert(5).
		insert(-6).
		insert(7)
	print(os.Stdout, tree.root, 0, 'M')
}
