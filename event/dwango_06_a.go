package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)
	st := make([]int, n)
	for i := range s {
		var title string
		var time int
		fmt.Scan(&title, &time)
		s[i] = title
		if i > 0 {
			st[i] = time + st[i-1]
		} else {
			st[i] = time
		}
	}
	var x string
	fmt.Scan(&x)

	var ans int
	for i := range s {
		if s[i] == x {
			ans = st[len(st)-1] - st[i]
			break
		}
	}

	fmt.Println(ans)
}
