package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	// create map without overlapping alphabets
	m := make(map[string]int, len(str))
	for _, s := range str {
		// mapを使わなくてもrune型から'a'を減算すれば、アルファベットの順番が取得できる。a: 0, b: 1, ...
		// fmt.Println(s - 'a')
		m[string(s)] = 1
	}

	// logic
	for _, s := range "abcdefghijklmnopqrstuvwxyz" {
		_, ok := m[string(s)]
		if !ok {
			fmt.Println(string(s))
			return
		}
	}
	fmt.Println("None")

	return
}
