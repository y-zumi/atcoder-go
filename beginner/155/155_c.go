package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	poll := make(map[string]int, n)
	max := 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		poll[s]++
		if max < poll[s] {
			max = poll[s]
		}
	}

	ans := make([]string, 0, n)
	for k, v := range poll {
		if v == max {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)

	for _, v := range ans {
		fmt.Println(v)
	}
}
