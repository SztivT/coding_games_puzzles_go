package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)
    var field[26][26] int
    for i := 0; i < N; i++ {
        scanner.Scan()
        inputs := strings.Split(scanner.Text()," ")
        for j := 0; j < N; j++ {
            if inputs[j] == "C" {
                field[i][j] = L
            } else {
                field[i][j] = 0
            }
        }
    }
    for L > 0 {
        for i:=0;i<N;i++ {
            for j:=0;j<N;j++ {
                if field[i][j] != L {
                    continue
                }
                for a:=-1;a < 2;a++ {
                    if i - a < 0 || i - a > N - 1 {
                        continue
                    }
                    for b :=-1;b<2;b++ {
                        if j -b < 0 || j - b >N -1 {
                            continue
                        }
                        if field[i-a][j-b] > L - 1 {
                            continue
                        }
                        field[i-a][j-b] = L - 1
                    }
                }
            }
        }
        L--
    }
    counter := 0
    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if field[i][j] == 0 {
                counter++
            }
        }
    }
    fmt.Printf("%d", counter)
}