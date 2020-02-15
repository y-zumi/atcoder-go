package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	var ans int
	for i := len(p) - 1; i >= 0; i-- {
		if min(p[:i+1]) == p[i] {
			ans++
		}
	}
	fmt.Println(ans)
}

func min(array []int) int {
	var m int
	for i, v := range array {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}
