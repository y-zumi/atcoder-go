package main

import "fmt"

func main() {
	a := make([]int, 3)
	c := make([][]int, 3)
	for i := range c {
		c[i] = make([]int, 3)
		fmt.Scan(&c[i][0], &c[i][1], &c[i][2])
	}

	for a[0] = 0; a[0] <= 100; a[0]++ {
		for a[1] = 0; a[1] <= 100; a[1]++ {
			for a[2] = 0; a[2] <= 100; a[2]++ {
				if c[0][0]-a[0] == c[0][1]-a[1] && c[0][0]-a[0] == c[0][2]-a[2] &&
					c[1][0]-a[0] == c[1][1]-a[1] && c[1][0]-a[0] == c[1][2]-a[2] &&
					c[2][0]-a[0] == c[2][1]-a[1] && c[2][0]-a[0] == c[2][2]-a[2] {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
