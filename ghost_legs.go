package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var W, H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&W, &H)
    var field[100][100] byte
    for i := 0; i < H; i++ {
        scanner.Scan()
        line := scanner.Text()
        _ = line // to avoid unused error
        //Getting the field
        for j := 0; j < W; j++ {
            field[i][j] = line[j];
        }
    }
    for col := 0; col < W; col++ {
        if field[0][col] == 32 {
            continue;
        }
        curr_col := col;
        for row := 1; row < H - 1; row++ {
            if field[row][curr_col + 1] == '-' {
                curr_col++;
                for field[row][curr_col] != '|' {
                    curr_col++;
                }
                continue;
            } else if curr_col != 0 {
                if field[row][curr_col - 1] == '-' {
                    curr_col--;
                    for field[row][curr_col] != '|' {
                        curr_col--;
                    }
                }
            }
        }      
        fmt.Printf("%c%c\n", field[0][col], field[H - 1][curr_col]);
    }
}