package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width int
	var height int
	var text string
	var abc [30]string

	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &width)
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &height)
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &text)

	for i := 0; i < height; i++ {
		scanner.Scan()
		abc[i] = scanner.Text()
	}
	text = strings.ToUpper(text)

	for h := 0; h < height; h++ {
		for i := 0; i < len(text); i++ {
			offset := (int(text[i]) - 65) * width
			if text[i] < 'A' || text[i] > 'Z' {
				offset = 26 * width
			}
			fmt.Printf("%s", abc[h][offset:offset+width])
		}
		fmt.Printf("\n")
	}
}
