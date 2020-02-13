package main

import "fmt"

func main() {
	var n, s, t, w int
	fmt.Scan(&n, &s, &t, &w)

	var ans int
	if isAppropriate(w, s, t) {
		ans++
	}
	for i := 2; i <= n; i++ {
		var a int
		fmt.Scan(&a)
		if w += a; isAppropriate(w, s, t) {
			ans++
		}
	}

	fmt.Println(ans)
}

func isAppropriate(weight, low, up int) bool {
	if low <= weight && weight <= up {
		return true
	}

	return false
}
