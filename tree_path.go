package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key int
	left *Node
	right *Node
}

func makeChildren(parent *Node, k int, l int, r int)  {
	if parent.key == k {
		parent.left = &Node{key: l}
		parent.right = &Node{key: r}
	} else {
		if parent.left != nil {
			makeChildren(parent.left, k, l, r)
		}
		if parent.right != nil {
			makeChildren(parent.right, k, l, r)
		}
	}
}

func printPath(p *[127]string, i *int)  {
	for *i > -1 {
		fmt.Print(p[*i])
		if *i > 0 {
			fmt.Print(" ")
		}
		*i--
	}
}

func searchTarget(t *Tree, target *int) {
	if t.root.key == *target {
		fmt.Print("Root")
		return
	}
	index := 0
	found := false
	var path[127] string
	findTarget(t.root, target, &found, &index, &path)
	index--
	printPath(&path, &index)
}

func findTarget(n *Node, t *int, found *bool, i *int, p *[127]string) {
	if n.key == *t {
		*found = true
		return
	}
	if n.right != nil && !*found {
		findTarget(n.right, t, found, i, p)
		if *found {
			p[*i] = "Right"
			*i++
		}
	}
	if n.left != nil && !*found {
		findTarget(n.left, t, found, i, p)
		if *found {
			p[*i] = "Left"
			*i++
		}
	}
}

func main() {
	var nodes int
	var target int
	var internalNodes int
	var t Tree

	_, _ = fmt.Scan(&nodes)
	_, _ = fmt.Scan(&target)
	_, _ = fmt.Scan(&internalNodes)
	for i := 0; i < internalNodes; i++ {
		var parent int
		var leftChild int
		var rightChild int
		_, _ = fmt.Scan(&parent)
		_, _ = fmt.Scan(&leftChild)
		_, _ = fmt.Scan(&rightChild)
		if t.root == nil {
			t.root = &Node{key: parent}
		}
		makeChildren(t.root, parent, leftChild, rightChild)
	}
	searchTarget(&t, &target)
}