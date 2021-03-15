package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	dividend  int
	divisor   int
	direction byte
	left      *Node
	right     *Node
	parent    *Node
}

func getParent(n *Node, direction byte, dividend *int, divisor *int) {
	if n == nil {
		return
	} else if n.dividend == 1 && direction == 'R' {
		*divisor++
		return
	} else if n.divisor == 1 && direction == 'L' {
		*dividend++
		return
	} else if n.direction != direction {
		getParent(n.parent, direction, dividend, divisor)
	} else {
		*dividend += n.parent.dividend
		*divisor += n.parent.divisor
	}
}

func (n *Node) insert(direction byte) {
	var dividend int
	var divisor int
	dividend = 0
	divisor = 0
	if direction == 'R' {
		getParent(n, 'L', &dividend, &divisor)
		n.right = &Node{
			dividend:  dividend + n.dividend,
			divisor:   divisor + n.divisor,
			direction: direction,
			parent:    n}
	} else if direction == 'L' {
		getParent(n, 'R', &dividend, &divisor)
		n.left = &Node{
			dividend:  dividend + n.dividend,
			divisor:   divisor + n.divisor,
			direction: direction,
			parent:    n}
	}
}

func displayNode(n *Node, path string) {
	if path == "" {
		fmt.Printf("%d/%d\n", n.dividend, n.divisor)
		return
	}
	if n.left == nil {
		n.insert('L')
	}
	if n.right == nil {
		n.insert('R')
	}
	if path[0] == 'L' {
		displayNode(n.left, path[1:])
	} else if path[0] == 'R' {
		displayNode(n.right, path[1:])
	}
}

func medianWalk(n *Node, dividend *int, divisor *int, path string) {
	var target float64
	var node float64
	target = float64(*dividend) / float64(*divisor)
	node = float64(n.dividend) / float64(n.divisor)
	if target == node {
		fmt.Printf("%s\n", path)
		return
	}
	if target > node {
		path += "R"
		if n.right == nil {
			n.insert('R')
		}
		medianWalk(n.right, dividend, divisor, path)
	} else {
		path += "L"
		if n.left == nil {
			n.insert('L')
		}
		medianWalk(n.left, dividend, divisor, path)
	}
}

func main() {
	var n int
	var instructions [1000]string
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &instructions[i])
	}
	for i := 0; i < n; i++ {
		root := &Node{
			dividend: 1,
			divisor:  1}
		lineArr := strings.Split(instructions[i], "/")
		if len(lineArr) == 2 {
			dividend, _ := strconv.Atoi(lineArr[0])
			divisor, _ := strconv.Atoi(lineArr[1])
			medianWalk(root, &dividend, &divisor, "")
		} else {
			displayNode(root, lineArr[0])
		}
	}
}