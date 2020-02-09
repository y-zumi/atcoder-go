package main

import "fmt"

func main() {
	var N, T int
	fmt.Scan(&N, &T)

	ts := make([]int, N)
	for i := range ts {
		fmt.Scan(&ts[i])
	}

	var sum int
	for i := 0; i < N-1; i++ {
		sum += f(ts[i], ts[i+1], T)
	}
	fmt.Println(sum + T)
}

func f(t1, t2, T int) int {
	diff := abs(t1, t2)
	if diff >= T {
		return T
	}
	return diff
}

func abs(x, y int) int {
	a := x - y
	if a < 0 {
		return a * -1
	}
	return a
}

// t1 - t2 >= T
// 通常加算
// t1 - t2 < T
// t1に関してdiffだけを追加する
// 重複部分を取り除く
// t1=3, t2=5, diff=2, T=10, R=12
