package main

import "fmt"

type submit struct {
	number int
	result string
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	submits := make([]submit, m)
	for i, s := range submits {
		fmt.Scan(&s.number, &s.result)
		s.number--
		submits[i] = s
	}
	ac, wa := cal(submits, n)
	fmt.Println(ac, wa)
}

func cal(submits []submit, n int) (int, int) {
	accepts := make([]bool, n)
	warongs := make([]int, n)
	cntAC, cntWA := 0, 0

	for _, submit := range submits {
		if accepts[submit.number] != true {
			if submit.result == "AC" {
				accepts[submit.number] = true
				cntAC++
				cntWA += warongs[submit.number]
			} else if submit.result == "WA" {
				warongs[submit.number]++
			}
		}
	}

	return cntAC, cntWA
}
