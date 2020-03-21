package main

import "fmt"

type bill struct {
	_10000 int
	_5000  int
	_1000  int
}

func (b bill) String() string {
	return fmt.Sprintf("%d %d %d", b._10000, b._5000, b._1000)
}

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	fmt.Println(validate(n, y).String())
}

func validate(bills int, amount int) bill {
	for i := 0; i <= bills; i++ {
		for j := 0; j <= bills-i; j++ {
			k := bills - i - j
			if amount == i*10000+j*5000+k*1000 {
				return bill{i, j, k}
			}
		}
	}
	return bill{-1, -1, -1}
}
