package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputs := make([]int, 5)
	for i := range inputs {
		fmt.Scan(&inputs[i])
	}

	indexes := getCombineIndexes(len(inputs))
	set := make(map[int]struct{}, len(indexes))
	for _, index := range indexes {
		sum := 0
		for _, v := range index {
			sum += inputs[v]
		}
		set[sum] = struct{}{}
	}

	result := make([]int, 0, len(set))
	for k := range set {
		result = append(result, k)
	}

	sort.Ints(result)

	fmt.Println(result[len(result)-3])

}

func getCombineIndexes(bit int) [][]int {
	bits := generate(bit)
	filteredBits := filter(bits)
	indexes := findIndex(filteredBits, '1')
	return indexes
}

func generate(bit int) []string {
	num := 1
	for i := 0; i < bit; i++ {
		num *= 2
	}
	bitsArray := make([]string, 0, num)
	for i := 0; i < num; i++ {
		bitsArray = append(bitsArray, fmt.Sprintf("%05b\n", i))
	}
	return bitsArray
}

func filter(bitsArray []string) []string {
	result := make([]string, 0, len(bitsArray))
	for _, v := range bitsArray {
		if strings.Count(v, "1") == 3 {
			result = append(result, v)
		}
	}
	return result
}

func findIndex(array []string, subrune rune) [][]int {
	indexes := make([][]int, len(array))
	for i, str := range array {
		for j, r := range str {
			if r == subrune {
				indexes[i] = append(indexes[i], j)
			}
		}
	}
	return indexes
}
