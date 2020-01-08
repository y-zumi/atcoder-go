package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	count := strings.Count(s, "1")
	fmt.Println(count)
}
