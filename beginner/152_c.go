package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	var ans int
	var low bool
	for i := len(p) - 1; i >= 0; i-- {
		low = true
		for j := i; j >= 0; j-- {
			if p[i] > p[j] {
				low = false
				break
			}
		}
		if low == true {
			ans++
		}
	}
	fmt.Println(ans)

	// for j := 0; j < len(p); j++ {
	// 	for i := j; i < len(p); i++ {
	// 		if p[i] > p[j] {
	// 			// copy(p[i:], p[i+1:])
	// 			// p[len(p)-1] = 0
	// 			// p = p[:len(p)-1]
	// 			p = append(p[:i], p[i+1:]...)
	// 			i--
	// 		}
	// 	}
	// }
	// fmt.Println(len(p))
}
