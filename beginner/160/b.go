package main

import "fmt"

func main() {
	var x int64
	fmt.Scan(&x)

	b500 := x / 500
	x -= b500 * 500
	b5 := x / 5

	fmt.Println(b500*1000 + b5*5)

}
