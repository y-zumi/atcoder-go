package main

import "fmt"

type input struct {
	q  int
	f1 int
	f2 string
}

func main() {
	var s string
	var q int
	fmt.Scan(&s, &q)

	queries := make([]input, q)
	for i := range queries {
		fmt.Scan(&queries[i].q)
		if queries[i].q == 2 {
			fmt.Scan(&queries[i].f1, &queries[i].f2)
		}
	}
	fmt.Println(queries)

	for _, q := range queries {
		switch q.q {
		case 1:
			s = reverse(s)
		case 2:
			switch q.f1 {
			case 1:
				s = q.f2 + s
			case 2:
				s = s + q.f2
			}
		}
		fmt.Println(s)
	}
	fmt.Println(s)

}

func reverse(str string) string {
	rs := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
