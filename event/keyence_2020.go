package main

import "fmt"

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)

	var m int
	if h > w {
		m = h
	} else {
		m = w
	}

	var ans int
	for n > 0 {
		n -= m
		ans++
	}
	fmt.Println(ans)

}
