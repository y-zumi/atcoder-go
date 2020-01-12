package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	points := make([]int, n-1)
	for i, p := range points {
		fmt.Scan(&p)
		points[i] = p
	}

	p := calculateLastPoint(points, m, k)
	fmt.Println(p)
}

func calculateLastPoint(points []int, objective, limit int) int {
	x := objective*(len(points)+1) - sum(points)
	if x < 0 {
		return 0
	}
	if x > limit {
		return -1
	}
	return x
}

func sum(points []int) int {
	s := 0
	for _, p := range points {
		s += p
	}
	return s
}
