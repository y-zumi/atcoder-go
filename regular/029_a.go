package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	t := make([]int, N)
	for i := range t {
		fmt.Scan(&t[i])
	}

	ans := 0
	for bit := 0; bit < (1 << uint(N)); bit++ {
		m1, m2 := 0, 0
		for i := 0; i < N; i++ {
			if bit&(1<<uint(i)) == 1<<uint(i) {
				m1 += t[i]
			} else {
				m2 += t[i]
			}
		}
		time := max(m1, m2)
		if ans == 0 || time < ans {
			ans = time
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
