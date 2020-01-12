package main

import "fmt"

type submit struct {
	number int
	result string
}

type score struct {
	acNum int
	waNum int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	submits := make([]submit, m)
	for i, s := range submits {
		fmt.Scan(&s.number, &s.result)
		submits[i] = s
	}
	s := cal(submits)
	fmt.Println(s.acNum, s.waNum)
}

func cal(submits []submit) score {
	scores := make(map[int]*score)

	for _, submit := range submits {
		_, ok := scores[submit.number]
		if !ok {
			scores[submit.number] = &score{0, 0}
		}

		if scores[submit.number].acNum != 1 {
			if submit.result == "AC" {
				scores[submit.number].acNum++
			}
			if submit.result == "WA" {
				scores[submit.number].waNum++
			}
		}
	}

	r := score{0, 0}
	for _, s := range scores {
		r.acNum += s.acNum
		r.waNum += s.waNum
	}

	return r
}
