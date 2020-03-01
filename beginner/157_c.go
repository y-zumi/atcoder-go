package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]int, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&s[i], &c[i])
	}

	ans := make(map[int]int, n)
	for i := 0; i < m; i++ {
		v, ok := ans[s[i]]
		if ok {
			if v != c[i] {
				fmt.Println(-1)
				return
			}
		}
		if s[i] == 1 && c[i] == 0 {
			fmt.Println(-1)
			return
		}
		ans[s[i]] = c[i]
	}

	var num int
	v, ok := ans[1]
	if ok {
		num += v * 100
	}
	v, ok = ans[2]
	if ok {
		num += v * 10
	}
	v, ok = ans[3]
	if ok {
		num += v
	}

	fmt.Println(num)
}

// 1~3 keta
// バッティングしたらアウト
// １桁目が０であればアウト
// それ以外は１桁目なら１で、２桁目以降は０で埋める
