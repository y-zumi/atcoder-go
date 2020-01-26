package main

import "fmt"

func main() {
	var h, c, sum int
	fmt.Scan(&h)
	for h > 0 {
		h = h / 2
		c++
	}
	sum = 1
	for c > 0 {
		sum *= 2
		c--
	}
	fmt.Println(sum - 1)

}
