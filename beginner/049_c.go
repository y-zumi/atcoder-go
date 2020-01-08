package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if check(s) == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func check(s string) bool {
	for len(s) > 0 {
		prev := s
		s = trimSuffix(s, "dream")
		s = trimSuffix(s, "dreamer")
		s = trimSuffix(s, "erase")
		s = trimSuffix(s, "eraser")
		if len(s) == len(prev) {
			return false
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}

func trimSuffix(s, suffix string) string {
	if len(s) < len(suffix) {
		return s
	}
	if s[len(s)-len(suffix):] == suffix {
		return s[:len(s)-len(suffix)]
	}
	return s
}
