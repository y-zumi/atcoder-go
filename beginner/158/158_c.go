package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1.0; i <= 10000.0; i++ {
		if int(i*0.08) == a && int(i*0.1) == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
