package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var a int
	var sequenceOfOnesStart, maxSequenceOfOnesLength int
	var prevSequenceOfOnesStart, prevMaxSequenceOfOnesLength int

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 1 {
			maxSequenceOfOnesLength++
			if i == 0 || sequenceOfOnesStart == 0 {
				sequenceOfOnesStart = i + 1
			}
		} else {
			if maxSequenceOfOnesLength > prevMaxSequenceOfOnesLength {
				prevMaxSequenceOfOnesLength = maxSequenceOfOnesLength
				prevSequenceOfOnesStart = sequenceOfOnesStart
			}
			maxSequenceOfOnesLength = 0
			sequenceOfOnesStart = 0
		}
	}

	if maxSequenceOfOnesLength < prevMaxSequenceOfOnesLength {
		maxSequenceOfOnesLength = prevMaxSequenceOfOnesLength
		sequenceOfOnesStart = prevSequenceOfOnesStart
	}

	fmt.Printf("\n%d %d\n", sequenceOfOnesStart, maxSequenceOfOnesLength)
}
