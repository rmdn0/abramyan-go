package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var a int
	var max, minSequenceOfMaximumsLength, prevMinSequenceOfMaximumsLength int

	if n > 0 {
		fmt.Scan(&a)
		max = a
		minSequenceOfMaximumsLength = 1
		prevMinSequenceOfMaximumsLength = n
	}

	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a == max {
			minSequenceOfMaximumsLength++
		} else {
			if a > max {
				max = a
				prevMinSequenceOfMaximumsLength = n
				minSequenceOfMaximumsLength = 1
			} else if minSequenceOfMaximumsLength != 0 && minSequenceOfMaximumsLength < prevMinSequenceOfMaximumsLength {
				prevMinSequenceOfMaximumsLength = minSequenceOfMaximumsLength
				minSequenceOfMaximumsLength = 0
			} else {
				minSequenceOfMaximumsLength = 0
			}
		}
	}

	if minSequenceOfMaximumsLength > prevMinSequenceOfMaximumsLength || minSequenceOfMaximumsLength == 0 {
		minSequenceOfMaximumsLength = prevMinSequenceOfMaximumsLength
	}

	fmt.Printf("\n%d\n", minSequenceOfMaximumsLength)
}
