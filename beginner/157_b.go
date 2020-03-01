package main

import "fmt"

func main() {
	a := make([]int, 9)
	ans := make([]bool, 9)
	for i := 0; i < 9; i++ {
		fmt.Scan(&a[i])
	}
	var n int
	fmt.Scan(&n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	for _, v := range b {
		for i, w := range a {
			if v == w {
				ans[i] = true
			}
		}
	}

	if bingo(ans) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func bingo(a []bool) bool {
	w := 3
	for x := 0; x < w; x++ {
		if a[x*w+0] && a[x*w+1] && a[x*w+2] {
			return true
		}
		if a[x+w*0] && a[x+w*1] && a[x+w*2] {
			return true
		}
	}
	if a[0] && a[1*w+1] && a[2*w+2] {
		return true
	}
	if a[2] && a[1*w+1] && a[2*w+0] {
		return true
	}
	return false
}
