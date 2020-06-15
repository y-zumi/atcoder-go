package main

import (
	"fmt"
	"math"
)

func main() {
	var S string
	fmt.Scan(&S)

	num := make([]int, len(S))
	for i, r := range S {
		num[i] = int(r - '0')
	}

	arr := make([]int, 0, (1 << uint(len(S))))
	for bit := 0; bit < (1 << uint(len(S)-1)); bit++ {
		s := 0
		for i := 0; i < len(S); i++ {
			if bit&(1<<uint(i)) == (1 << uint(i)) {
				p := pow(num[s : i+1])
				arr = append(arr, p)
				s = i + 1
			}
		}
		arr = append(arr, pow(num[s:len(num)]))
	}

	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	fmt.Println(sum)

}

func pow(data []int) int {
	if len(data) == 1 {
		return data[0]
	}

	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i] * int(math.Pow10(len(data)-i-1))
	}
	return sum
}
