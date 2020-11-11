package main

import (
	"fmt"
)

func main() {
	var i, j int
	count := 0
	start := 1000
	end := 9999

	for i = start; i <= end; i++ {
		for j = i; j <= end; j++ {
			if isEvenEnd(i * j) {
				//fmt.Printf("%d is even-ended\n", i)
				count++
			} else {
				//fmt.Printf("%d is not even-ended\n", i)
			}
		}
	}
	fmt.Printf("There are %d even-ended numbers from %d to %d\n", count, start, end)
}

func isEvenEnd(x int) bool {
	xString := fmt.Sprintf("%d", x)

	if xString[0] == xString[len(xString)-1] {
		return true
	}
	return false
}
