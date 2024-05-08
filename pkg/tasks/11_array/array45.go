package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var diff, minDiff float64
	var firstInd, secondInd int
	isFirstDiff := true
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			diff = math.Abs(ar[i] - ar[j])
			if isFirstDiff {
				isFirstDiff = false
				minDiff = diff
				firstInd = i + 1
				secondInd = j + 1
			} else if minDiff > diff {
				minDiff = diff
				firstInd = i + 1
				secondInd = j + 1
			}
		}
	}
	fmt.Println(firstInd, secondInd)
}
