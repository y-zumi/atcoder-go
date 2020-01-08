package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	var cnt int
	for data = inOdd(data); data != nil; cnt++ {
		data = inOdd(devide(data))
	}
	fmt.Println(cnt)
}

func inOdd(data []int) []int {
	for _, d := range data {
		if d%2 == 1 {
			return nil
		}
	}
	return data
}

func devide(data []int) []int {
	for i, d := range data {
		data[i] = d / 2
	}
	return data
}
