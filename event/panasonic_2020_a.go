package main

import "fmt"

func main() {
	d := []int{1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51}
	var k int
	fmt.Scan(&k)
	fmt.Println(d[k-1])
}
