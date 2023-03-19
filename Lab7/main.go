package main

import "fmt"

func split(sum float64) (x, y, z float64) {
	x = sum * 4 / 9
	y = sum - x
	z = x + y
	return
}

func main() {
	fmt.Println(split(9))
}
