package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 7.0
	circumference := 2 * math.Pi * radius
	fmt.Printf("The circumference of a circle is %v", circumference)
}
