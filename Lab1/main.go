package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	if n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}
}
