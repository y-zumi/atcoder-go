package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)
	fmt.Println(n)

	if n%2 == 1 {
		fmt.Println(0)
		return
	}

	var sum int
	for ; n >= 2; n = n - 2 {
		fmt.Println(n)
		if n%5 == 0 {
			sum++
		}
	}
	fmt.Println(sum)

}

func f(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * f(n-2)
}
