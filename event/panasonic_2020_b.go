package main

import "fmt"

func main() {
	var h, w int64
	fmt.Scan(&h, &w)

	s := h * w
	if h == 1 || w == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(s/2 + s%2)
	}

}
