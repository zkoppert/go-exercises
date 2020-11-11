package main

import (
	"fmt"
)

func main() {
	for i := 1000; i <= 1010; i++ {
		isEvenEnd(i)
	}
}

func isEvenEnd(x int) {
	fmt.Printf("%x is even-ended")
	fmt.Printf("%x is not even-ended")

}
