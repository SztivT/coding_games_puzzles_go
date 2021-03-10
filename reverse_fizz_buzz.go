package main

import "fmt"
import "os"
import "bufio"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var n int
    var fizzindex int
    var buzzindex int
    var fizz int
    var buzz int
    var offset int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    line_arr := make([]string, n);

    offset = 1;
    for i := 0; i < n; i++ {
        scanner.Scan()
        line := scanner.Text()
        _ = line // to avoid unused error
        if fizzindex == 0 && (line == "Fizz" || line == "FizzBuzz") {
            fizzindex = i;
        } else if (line == "Fizz" || line == "FizzBuzz") && fizzindex != -1  {
            fizz = i - fizzindex;
            fizzindex = -1;
        }
        if buzzindex == 0 && (line == "Buzz" || line == "FizzBuzz") {
            buzzindex = i;
        } else if (line == "Buzz" || line == "FizzBuzz") && buzzindex != -1 {
            buzz = i - buzzindex;
            buzzindex = -1;
        } else {
            num, err := strconv.Atoi(line)
            if err != nil {
                err = nil
            }
            if num - i > 0 && offset == 1 {
                offset = num - i
            }
        }
        line_arr[i] = line;

    }
    if fizzindex != -1 {
        fizz = fizzindex + offset;
    }
    if buzzindex != -1 {
        buzz = buzzindex + offset;
    }

   
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Printf("%d %d", fizz, buzz)
}