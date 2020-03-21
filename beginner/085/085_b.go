package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]int, n)
	for i := range data {
		fmt.Scan(&data[i])
	}
	fmt.Println(len(distinct(data)))
}

func distinct(data []int) []int {
	newData := make([]int, 0, len(data))
	for _, d := range data {
		if !contain(newData, d) {
			newData = append(newData, d)
		}
	}
	return newData
}

func contain(data []int, v int) bool {
	for _, d := range data {
		if d == v {
			return true
		}
	}
	return false
}
