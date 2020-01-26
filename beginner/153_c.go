package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	sort.Ints(h)

	if len(h) >= k {
		fmt.Println(sum(h[:len(h)-k]))
	} else {
		fmt.Println(0)
	}

}

func sum(data []int) int {
	sum := 0
	for _, d := range data {
		sum += d
	}
	return sum
}
