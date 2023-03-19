package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Varun"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	var radius float32 = 5.0
	area := Pi * radius * radius
	fmt.Printf("The area of a circle with radius %v is %v\n", radius, area)

	var i int
	fmt.Println("The value of i is", i)

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)
}
