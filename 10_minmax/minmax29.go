package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var a int
	var min, maxSequenceOfMinsLength, prevMaxSequenceOfMinsLength int

	if n > 0 {
		fmt.Scan(&a)
		min = a
		maxSequenceOfMinsLength = 1
	}

	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a == min {
			maxSequenceOfMinsLength++
		} else {
			if a < min {
				min = a
				prevMaxSequenceOfMinsLength = 0
				maxSequenceOfMinsLength = 1
			} else if maxSequenceOfMinsLength > prevMaxSequenceOfMinsLength {
				prevMaxSequenceOfMinsLength = maxSequenceOfMinsLength
				maxSequenceOfMinsLength = 0
			} else {
				maxSequenceOfMinsLength = 0
			}
		}
	}

	if maxSequenceOfMinsLength < prevMaxSequenceOfMinsLength {
		maxSequenceOfMinsLength = prevMaxSequenceOfMinsLength
	}

	fmt.Printf("\n%d\n", maxSequenceOfMinsLength)
}
