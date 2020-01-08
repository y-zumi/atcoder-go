package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	data := make([]int, n)
	for i := range data {
		fmt.Scan(&data[i])
	}
	data = bubbleSort(data)

	var diff int
	for i, d := range data {
		if i%2 == 0 {
			diff += d
		} else {
			diff -= d
		}
	}
	fmt.Println(diff)
}

func bubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-(i+1); j++ {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
