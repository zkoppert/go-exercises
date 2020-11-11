package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 21; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizz buzz\n")
		} else if i%3 == 0 {
			fmt.Printf("fizz\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz\n")
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}
