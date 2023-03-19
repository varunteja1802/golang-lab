package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var name string = "John"
	fmt.Printf("Type: %T Value: %v\n", name, name)

	var ratio float64 = 3.14
	fmt.Printf("Type: %T Value: %v\n", ratio, ratio)

	var sum int = int(MaxInt) + 1
	fmt.Printf("Type: %T Value: %v\n", sum, sum)
}
