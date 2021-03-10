package main

import "fmt"
import "os"
import "bufio"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // width: the number of cells on the X axis
    var width int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&width)

    // height: the number of cells on the Y axis
    var height int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&height)
    var table[31][31] byte
    for i := 0; i < height; i++ {
        scanner.Scan()
        line := scanner.Text()
        _ = line // to avoid unused error // width characters, each either 0 or .
        for j := 0; j < width; j++ {
            table[i][j] = line[j]
        }
    }
    //see the field
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            print(table[i][j])
            print(" ")
        }
        print("\n")
    }
    //cherrypick
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if table[i][j] == 48 {
                //to right
                a := -1
                b := -1
                for k:= j + 1; k < width; k++ {
                    if table[i][k] == 48 {
                        a = i
                        b = k
                        break
                    }
                }
                //to bottom
                x := -1
                y := -1
                for k:= i + 1; k < height; k++ {
                    if table[k][j] == 48 {
                        x = k
                        y = j
                        break
                    }
                }
                fmt.Printf("%d %d %d %d %d %d\n", j, i, b, a, y, x)
            }
        }
    }
}