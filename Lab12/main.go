package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	i = 42
	f = 3.14159
	b = true
	s = "Hello, World!"

	fmt.Printf("i: %v\n", i) //%d for integer can be given
	fmt.Printf("f: %v\n", f) //%f for decimals can be given
	fmt.Printf("b: %v\n", b) //%t for boolean can be given
	fmt.Printf("s: %v\n", s) //%s for strings can be given
}
