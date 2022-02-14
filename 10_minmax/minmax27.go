package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var curr, prev int
	var equalSequenceStart, maxEqualSequenceLength int
	var prevEqualSequenceStart, prevMaxEqualSequenceLength int

	if n > 0 {
		fmt.Scan(&prev)
		equalSequenceStart = 1
		maxEqualSequenceLength = 1
	}

	for i := 1; i < n; i++ {
		fmt.Scan(&curr)
		if curr == prev {
			maxEqualSequenceLength++
		} else {
			if maxEqualSequenceLength > prevMaxEqualSequenceLength {
				prevMaxEqualSequenceLength = maxEqualSequenceLength
				prevEqualSequenceStart = equalSequenceStart
			}
			maxEqualSequenceLength = 1
			equalSequenceStart = i + 1
		}
		prev = curr
	}

	if maxEqualSequenceLength <= prevMaxEqualSequenceLength {
		maxEqualSequenceLength = prevMaxEqualSequenceLength
		equalSequenceStart = prevEqualSequenceStart
	}

	fmt.Printf("\n%d %d\n", equalSequenceStart, maxEqualSequenceLength)
}
