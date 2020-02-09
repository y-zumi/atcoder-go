package main

import "fmt"

func main() {
	var N, T int
	fmt.Scan(&N, &T)

	// TODO: bufferで入出力すれば1000msぐらい早くなる
	ts := make([]int, N)
	for i := range ts {
		fmt.Scan(&ts[i])
	}

	var ans int
	for i := 0; i < N-1; i++ {
		diff := ts[i+1] - ts[i]

		if diff >= T {
			ans += T
		} else {
			ans += diff
		}
	}
	fmt.Println(ans + T)
}
