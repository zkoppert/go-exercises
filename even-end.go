package main

import (
	"fmt"
)

func main() {
	var i int
	for i = 1; i <= 35; i++ {
		if isEvenEnd(i) {
			fmt.Printf("%d is even-ended\n", i)
		} else {
			fmt.Printf("%d is not even-ended\n", i)
		}
	}
}

func isEvenEnd(x int) bool {
	xString := fmt.Sprintf("%d", x)

	if xString[0] == xString[len(xString)-1] {
		return true
	}
	return false
}
