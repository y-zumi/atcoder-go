package main

import (
	"fmt"
	"sort"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	array := make([]int, n)
	for i := range array {
		fmt.Scan(&array[i])
		array = append(array, array[i]+k)
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = array[len(ans)-1+i] - array[i]

	}
	sort.Ints(ans)
	fmt.Println(ans[0])

}
