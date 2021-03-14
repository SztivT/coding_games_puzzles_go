package main

import "fmt"

func memset(arr *[1000001]int, n int) {
	for i := 0; i < n; i++ {
		arr[i] = -1
	}
}

func main() {
	var ve[1000001] int
	var li[1000001] int
	var N int
	_, _ = fmt.Scan(&ve[0])
	_, _ = fmt.Scan(&N)
	memset(&li, N)
	for i:= 0; i < N; i++ {
		if li[ve[i]] == -1 {
			ve[i + 1] = 0
		} else {
			ve[i + 1] = i - li[ve[i]]
		}
		li[ve[i]] = i
	}
	fmt.Println(ve[N - 1])
}