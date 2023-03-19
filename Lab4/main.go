package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	radius := (add(42, 13))
	circumference := 2 * math.Pi * float64(radius)
	fmt.Printf("The circumference of a circle is %v", circumference)
}
