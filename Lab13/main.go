package main

import (
	"fmt"
)

func main() {
	var x int = 3
	var y float64 = 9.6
	var sum = x + int(y)
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(sum)
}
