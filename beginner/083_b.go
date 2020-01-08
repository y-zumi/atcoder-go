package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var sum int
	for i := 1; i <= n; i++ {
		if s := sumOfDigit(i); a <= s && s <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumOfDigit(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
