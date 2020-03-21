package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var ans string
	for range s {
		ans += "x"
	}
	fmt.Println(ans)
}
