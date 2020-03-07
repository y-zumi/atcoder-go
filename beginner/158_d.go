package main

import (
	"container/list"
	"fmt"
	"strings"
)

type input struct {
	q  int
	f1 int
	f2 string
}

func main() {
	var s string
	var qn int
	fmt.Scan(&s, &qn)

	var q, f1 int
	var f2 string
	rev := false
	// top := make([]byte, 0, qn)
	// bottom := make([]byte, 0, qn)
	dq := list.New()
	dq.PushFront(s)

	for i := 0; i < qn; i++ {
		fmt.Scan(&q)
		if q == 1 {
			rev = !rev
		} else {
			fmt.Scan(&f1, &f2)
			if (!rev && f1 == 1) || (rev && f1 == 2) {
				dq.PushFront(f2)
			} else {
				dq.PushBack(f2)
			}
		}
	}

	ss := make([]string, dq.Len())
	i := 0
	for e := dq.Front(); e != nil; e = e.Next() {
		ss[i] = e.Value.(string)
		i++
	}
	ans := strings.Join(ss, "")
	if rev {
		ans = reverse(ans)
	}
	fmt.Println(ans)
}

func reverse(str string) string {
	rs := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
