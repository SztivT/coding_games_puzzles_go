package main

import "fmt"

func stringConvert(input string, output *string, binary *[128]string, characters *[128]int) int {
	if len(input) == 0 {
		return 0
	}
	binLen := 0
	for i := 0; characters[i] != 255; i++ {
		binLen = len(binary[i])
		if binLen > len(input) {
			continue
		}
		if input[0:binLen] == binary[i] {
			*output += string(byte(characters[i]))
			return stringConvert(input[binLen:], output, binary, characters)
		}
	}
	return len(input)
}

func main() {
	var n int
	var b[128] string
	var c[128] int
	var s string
	var o string
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &b[i])
		_, _ = fmt.Scanf("%d", &c[i])
	}
	c[n] = 255
	_, _ = fmt.Scanf("%s\n", &s)
	converted := stringConvert(s, &o, &b, &c)
	if converted > 0 {
		fmt.Printf("DECODE FAIL AT INDEX %d", len(s) - converted)
	} else {
		fmt.Printf("%s", o)
	}
}