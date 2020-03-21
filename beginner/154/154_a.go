package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	var a, b int
	fmt.Scan(&a, &b)
	var u string
	fmt.Scan(&u)

	if s == u {
		a--
	} else {
		b--
	}

	fmt.Println(a, b)
}
