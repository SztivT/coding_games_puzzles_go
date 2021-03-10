package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for l.Len() != 0 {
		fmt.Printf("%d", l.Front().Value.(int))
		l.Remove(l.Front())
		if l.Len() != 0 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
}

func fillNextLine(l *list.List) {
	c := list.New()
	for ; ; {
		if l.Len() == 0 {
			l.PushBackList(c)
			return
		}
		counter := 1
		value := l.Front().Value.(int)
		l.Remove(l.Front())
		for l.Len() != 0 && value == l.Front().Value.(int) {
			counter++
			l.Remove(l.Front())
		}
		c.PushBack(counter)
		c.PushBack(value)
	}
}

func main() {
	//the original sequence
	var R int
	fmt.Scan(&R)
	//the line we want to display
	var L int
	fmt.Scan(&L)
	//previous line
	line := list.New()
	//initializing the basis
	line.PushBack(R)
	for i:=1; i < L; i++ {
		fillNextLine(line)
	}
	printList(line)
}

