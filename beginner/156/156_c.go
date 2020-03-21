package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}

	var sum int
	for i := range x {
		sum += x[i]
	}
	avg := float64(sum) / float64(len(x))
	avgInt, _ := strconv.Atoi(fmt.Sprintf("%.0f", avg))
	var min int
	for i := range x {
		min += (x[i] - avgInt) * (x[i] - avgInt)
	}
	fmt.Println(min)
}
