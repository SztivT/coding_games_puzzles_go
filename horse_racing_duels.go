package main

import "fmt"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    diff := 10000000;
    power_arr := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&power_arr[i])
    }
    sort.Ints(power_arr)
    for i:= N - 1; i > 0; i-- {
        curr_diff := power_arr[i] - power_arr[i - 1]
        if curr_diff < diff {
            diff = curr_diff;
        }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(diff)// Write answer to stdout
}