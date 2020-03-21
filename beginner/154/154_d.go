package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
		// 任意の数に+1して2で割るとその数の期待値になる。計算中を整数でやりたいので、+1しておいて最後に2で割る
		//(6+1)/2=3.5
		p[i]++
	}

	idx := make([]int, 0, k)
	var max, sum int
	for i := 0; i < n; i++ {
		sum += p[i]
		idx = append(idx, p[i])
		if len(idx) > k {
			sum -= idx[0]
			idx = idx[1:]
		}
		if len(idx) == k {
			if max < sum {
				max = sum
			}
		}
	}
	ans := float64(max)
	ans /= 2.0
	fmt.Printf("%.12f\n", ans)
}
