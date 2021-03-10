package main

import "fmt"
import "os"
import "bufio"

func displayMsg(message [50]int)  {
	for i:=0;i<len(message);i++ {
		if message[i] == 0 {
			break
		}
		fmt.Printf("%c", message[i])
	}
	fmt.Printf("\n")
}
func encodeOffset(msg *[50]int, offset int) {
	for i := 0; i < len(*msg); i++ {
		if msg[i] == 0 {
			break
		}
		msg[i] += offset
		if msg[i] > 90 {
			shift := msg[i]
			for shift > 90 {
				shift -= 26
			}
			msg[i] = shift
		}
		offset++
	}
}
func decodeOffset(msg *[50]int, offset int) {
	for i := 0; i < len(*msg); i++ {
		if msg[i] == 0 {
			break
		}
		msg[i] -= offset
		if msg[i] < 65 {
			shift := msg[i]
			for shift < 65 {
				shift += 26
			}
			msg[i] = shift
		}
		offset++
	}
}
func rotorEncode(msg *[50]int, rotor [26]int) {
	for i := 0; i < len(msg); i++ {
		if msg[i] == 0 {
			break
		}
		msg[i] = rotor[msg[i] - 65]
	}
}
func rotorDecode(msg *[50]int, rotor [26]int) {
	for i := 0; i < len(msg); i++ {
		if msg[i] == 0 {
			break
		}
		for j := 0; j < len(rotor); j++ {
			if rotor[j] == msg[i] {
				msg[i] = j + 65
				break
			}
		}
	}
}
func encode (msg *[50]int, rotors[3][26] int, offset int) {
	encodeOffset(msg, offset)
	for i := 0; i < 3; i++ {
		rotorEncode(msg, rotors[i])
	}
}
func decode (msg *[50]int, rotors[3][26] int, offset int) {
	for i := 2; i > -1; i-- {
		rotorDecode(msg, rotors[i])
	}
	decodeOffset(msg, offset)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var offset int
	var rotors[3][26] int
	var msg[50] int

	scanner.Buffer(make([]byte, 1000000), 1000000)
	//getting operation
	scanner.Scan()
	operation := scanner.Text()
	_ = operation // to avoid unused error
	//getting base offset
	scanner.Scan()
	fmt.Sscan(scanner.Text(),&offset)
	//getting rotors
	for i := 0; i < 3; i++ {
		scanner.Scan()
		rotor := scanner.Text()
		for j := 0; j < 26; j++ {
			rotors[i][j] = int(rotor[j])
		}
	}
	scanner.Scan()
	message := scanner.Text()
	_ = message // to avoid unused error
	for i := 0; i < len(message); i++ {
		msg[i] = int(message[i])
	}
	msg[len(message)] = 0
	if operation == "ENCODE"{
		encode(&msg, rotors, offset)
	}
	if operation == "DECODE" {
		decode(&msg, rotors, offset)
	}
	displayMsg(msg)
}