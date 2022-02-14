package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, maxEvenSequenceLength, prevMaxEvenSequenceLength int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if isEven(a) {
			maxEvenSequenceLength++
		} else {
			if maxEvenSequenceLength > prevMaxEvenSequenceLength {
				prevMaxEvenSequenceLength = maxEvenSequenceLength
			}
			maxEvenSequenceLength = 0
		}
	}
	if prevMaxEvenSequenceLength > maxEvenSequenceLength {
		maxEvenSequenceLength = prevMaxEvenSequenceLength
	}
	fmt.Printf("\n%d\n", maxEvenSequenceLength)
}

func isEven(x int) bool {
	return x%2 == 0
}
