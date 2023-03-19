package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"Varun", "VarunTeja", "KalakotiVarunTeja", "Teja"}
	substr := "Teja"

	for i := 1; i < len(arr); i++ {
		if strings.IndexAny(arr[i], substr) >= 0 {
			fmt.Printf("%s is a substring of %s\n", substr, arr[i])
		} else {
			fmt.Printf("%s is not a substring of %s\n", substr, arr[i])
		}
	}
}
