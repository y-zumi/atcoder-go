package main

import "fmt"

type coordinateTime struct {
	coordinate coordinate
	time       int
}

func (departure *coordinateTime) canMove(arrival coordinateTime) bool {
	cd := departure.coordinate.distance(arrival.coordinate)
	td := abs(arrival.time - departure.time)
	if cd <= td && cd%2 == td%2 {
		return true
	}
	return false
}

type coordinate struct {
	x int
	y int
}

func (c *coordinate) distance(c2 coordinate) int {
	xd := abs(c.x - c2.x)
	yd := abs(c.y - c2.y)
	return xd + yd
}

func main() {
	var n int
	fmt.Scan(&n)

	cts := make([]coordinateTime, n+1)
	for i, ct := range cts {
		if i == 0 {continue}
		fmt.Scan(&ct.time, &ct.coordinate.x, &ct.coordinate.y)
		cts[i] = ct
	}

	if checkRoute(cts) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func checkRoute(cts []coordinateTime) bool {
	for i := 0; i < len(cts)-1; i++ {
		if !cts[i].canMove(cts[i+1]) {
			return false
		}
	}
	return true
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}
