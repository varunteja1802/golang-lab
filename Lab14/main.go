package main

import "fmt"

func main() {
	v := 42.5 // change me!
	fmt.Printf("v is of type %T\n", v)

	v = 3.14
	fmt.Printf("v is of type %T\n", v)

	var x int64 = 100
	fmt.Printf("x is of type %T\n", x)

	y := float64(x)
	fmt.Printf("y is of type %T\n", y)
}
