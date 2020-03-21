package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)
	fmt.Println(next(c))
}

func next(alphabet string) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for i, a := range alphabets {
		if string(a) == alphabet {
			return string(alphabets[i+1])
		}
	}
	return ""
}
