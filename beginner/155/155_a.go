package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	ans := make(map[int]int, 3)
	ans[a] = 1
	ans[b] = 1
	ans[c] = 1

	if len(ans) == 1 || len(ans) == 3 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")

}
