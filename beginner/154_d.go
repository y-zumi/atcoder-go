package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	indexes := make([]float64, k)
	max := 0
	maxIndexes := make([]float64, k)
	for i := 0; i <= n-k; i++ {
		var sum int
		for j := 0; j < k; j++ {
			sum += p[i+j]
			indexes[j] = float64(p[i+j])
		}
		// fmt.Println(max, sum, indexes)
		if max < sum {
			max = sum
			copy(maxIndexes, indexes)
		}
	}

	// fmt.Println(max, maxIndexes)

	var sum float64
	for _, v := range maxIndexes {
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
