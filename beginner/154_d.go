package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	p := make([]float64, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	idx := make([]float64, 0, k)
	maxIdx := make([]float64, k)
	var max, sum float64
	var head int
	for i := 0; i < n; i++ {
		sum += p[i]
		idx = append(idx, p[i])
		if i >= k-1 {
			// fmt.Println(sum, idx)
			if max < sum {
				max = sum
				copy(maxIdx, idx)
			}
			sum -= p[head]
			idx = idx[1:]
			head++
		}
	}

	// fmt.Println(max, maxIdx)

	// indexes := make([]float64, k)
	// max := 0
	// maxIndexes := make([]float64, k)
	// for i := 0; i <= n-k; i++ {
	// 	var sum int
	// 	for j := 0; j < k; j++ {
	// 		sum += p[i+j]
	// 		indexes[j] = float64(p[i+j])
	// 	}
	// 	// fmt.Println(max, sum, indexes)
	// 	if max < sum {
	// 		max = sum
	// 		copy(maxIndexes, indexes)
	// 	}
	// }

	// // fmt.Println(max, maxIndexes)

	sum = 0.0
	for _, v := range maxIdx {
		sum += expectation(v)
	}
	fmt.Printf("%.12f\n", sum)

}

func expectation(p float64) float64 {
	ratio := 1.0 / float64(p)
	var sum float64
	for i := 1.0; i <= p; i++ {
		sum += i * ratio
	}
	return sum
}
