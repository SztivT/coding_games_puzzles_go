package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) insert(k int) {
	if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}
	}
}

func preOrder(n *Node, e *int) {
	if n == nil {
		return
	}
	fmt.Printf("%d", n.key)
	if *e > 1 {
		fmt.Printf(" ")
	}
	*e--
	preOrder(n.left, e)
	preOrder(n.right, e)
}

func postOrder(n *Node, e *int) {
	if n == nil {
		return
	}
	postOrder(n.left, e)
	postOrder(n.right, e)
	fmt.Printf("%d", n.key)
	if *e > 1 {
		fmt.Printf(" ")
	}
	*e--
}

func inOrder(n *Node, e *int) {
	if n == nil {
		return
	}
	inOrder(n.left, e)
	fmt.Printf("%d", n.key)
	if *e > 1 {
		fmt.Printf(" ")
	}
	*e--
	inOrder(n.right, e)
}

func levelOrder(n *Node, e *int) {
	height := treeHeight(n)
	for i := 0; i < height + 1; i++ {
		printLevel(n, i, e)
	}

}

func printLevel(n *Node, h int, e *int) {
	if n == nil {
		return
	}
	if h == 1 {
		fmt.Printf("%d", n.key)
		if *e > 1 {
			fmt.Printf(" ")
		}
		*e--
	} else if h > 1 {
		printLevel(n.left, h - 1, e)
		printLevel(n.right, h - 1, e)
	}
}

func treeHeight (n *Node) int {
	if n == nil {
		return 0
	}
	hLeft := treeHeight(n.left)
	hRight := treeHeight(n.right)

	if hLeft > hRight {
		return hLeft + 1
	} else {
		return hRight + 1
	}

}

func main() {
	var t Tree
	var n int
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var vi int
		_, _ = fmt.Scanf("%d", &vi)
		if t.root == nil {
			t.root = &Node{key: vi}
		} else {
			t.root.insert(vi)
		}
	}
	elements := n
	preOrder(t.root, &elements)
	fmt.Printf("\n")
	elements = n
	inOrder(t.root, &elements)
	fmt.Printf("\n")
	elements = n
	postOrder(t.root, &elements)
	fmt.Printf("\n")
	elements = n
	levelOrder(t.root, &elements)
}
