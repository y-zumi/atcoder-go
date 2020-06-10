package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	v := make([]int, len(input))
	for i, r := range input {
		v[i] = int(r - '0')
	}

	n := len(input) - 1
	op := make([]string, n)
	var ans int
	wantAns := 7

	for bit := 0; bit < (1 << uint(n)); bit++ {
		ans = v[0]
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) == (1 << uint(i)) {
				ans += v[i+1]
				op[i] = "+"
			} else {
				ans -= v[i+1]
				op[i] = "-"
			}
		}
		if ans == wantAns {
			break
		}
	}

	var ansFormula string
	for i := range op {
		ansFormula = fmt.Sprintf("%s%d%s", ansFormula, v[i], op[i])
	}
	ansFormula = fmt.Sprintf("%s%d=%d", ansFormula, v[len(v)-1], wantAns)
	fmt.Println(ansFormula)
}
