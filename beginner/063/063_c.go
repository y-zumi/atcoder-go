package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	scores := make([]int, n)
	for i := range scores {
		fmt.Scan(&scores[i])
	}

	ans := availableScore(scores)
	fmt.Println(ans)
}

func availableScore(scores []int) int {
	max := sum(scores)
	if bugCal(max) != 0 {
		return max
	}

	sort.Ints(scores)
	for _, s := range scores {
		available := max - s
		if bugCal(available) != 0 {
			return available
		}
	}

	return 0
}

func sum(scores []int) int {
	var sum int
	for _, s := range scores {
		sum += s
	}
	return sum
}

func bugCal(score int) int {
	if score%10 == 0 {
		return 0
	}
	return score
}
