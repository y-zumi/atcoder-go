package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	var m, ans int
	for i, v := range p {
		if i == 0 || v < m {
			m = v
		}
		if p[i] == m {
			ans++
		}
	}
	fmt.Println(ans)
}
