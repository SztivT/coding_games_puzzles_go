package main

import "fmt"

func move(movementRange int, tower int) int {
	if tower > 1 {
		tower = 0
	}
	if movementRange  == 1 {
		return movementRange
	}
	if movementRange < 25 {
		return movementRange / 2 + (movementRange / 4) * (tower * 2)
	}
	return movementRange / 2
}

func round(w0 *int, h0 *int, width *int, height *int, posW *int, posH *int, direction string) {
	if len(direction) == 0 {
		return
	}
	switch direction[0] {
	case 'U':
		*height = *posH
		*posH -= move(*height - *h0, *width - *w0)
	case 'D':
		*h0 = *posH
		*posH += move(*height - *h0, *width - *w0)
	case 'L':
		*width = *posW
		*posW -= move(*width - *w0, *height - *h0)
	case 'R':
		*w0 = *posW
		*posW += move(*width - *w0, *height - *h0)
	}
	round(w0, h0, width, height, posW, posH, direction[1:])
}

func main() {
	var turnLeft int
	var width int
	var height int
	var posW int
	var posH int
	var direction string
	w0 := 0
	h0 := 0
	_, _ = fmt.Scanf("%d", &width)
	_, _ = fmt.Scanf("%d", &height)
	_, _ = fmt.Scanf("%d", &turnLeft)
	_, _ = fmt.Scanf("%d", &posW)
	_, _ = fmt.Scanf("%d", &posH)
	width--
	height--
	for turnLeft > 0 {
		_, _ = fmt.Scanf("%s", &direction)
		round(&w0, &h0, &width, &height, &posW, &posH, direction)
		fmt.Printf("%d %d\n", posW, posH)
		turnLeft--
	}
}