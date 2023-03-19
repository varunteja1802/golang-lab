package main

import (
	"fmt"
	"math"
)

func main() {
	sqrt := math.Sqrt(7)
	fmt.Printf("Now you have %g problems.\n", sqrt)

	if sqrt == math.Floor(sqrt) {
		fmt.Println("Looks like you have a perfect square!")
	} else {
		fmt.Println("Looks like you have a non-perfect square.")
	}
}
