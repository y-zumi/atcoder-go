package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := range x {
		fmt.Scan(&x[i], &y[i])
	}

	var maxDist float64
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := dist(x[i], y[i], x[j], y[j])
			if d > maxDist {
				maxDist = d
			}
		}
	}

	fmt.Println(maxDist)
}

func dist(x1, y1, x2, y2 float64) float64 {
	px := math.Pow((x2 - x1), 2)
	py := math.Pow((y2 - y1), 2)
	ans := math.Sqrt(px + py)
	return ans
}
