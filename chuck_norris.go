package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func chuckIt(input string) string {
	var output string
	for i := 0; i < len(input); i++ {
		current := input[i]
		if current == '1' {
			output += "0 0"
		}
		if current == '0' {
			output += "00 0"
		}
		if i == len(input)-1 {
			return output
		}
		for current == input[i+1] {
			output += "0"
			i++
			if i == len(input)-1 {
				break
			}
		}
		if i < len(input)-1 {
			output += " "
		}
	}
	return output
}

func encodeText(s string) string {
	var text, binary string
	for i := 0; i < len(s); i++ {
		binVal := strconv.FormatInt(int64(s[i]), 2)
		for i := len(binVal); i < 7; i++ {
			binary += "0"
		}
		binary += binVal
	}
text = chuckIt(binary)
return text
}

func main() {
	var input, output string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Scan()
	input = scanner.Text()
	output = encodeText(input)
	fmt.Printf(output)

}
