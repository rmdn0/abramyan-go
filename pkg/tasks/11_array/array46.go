package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("R = ")
	fmt.Scan(&r)

	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var diff, minDiff float64
	isFirstDiff := true
	var firstInd, secondInd int

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			diff = math.Abs(ar[i] + ar[j] - r)
			if isFirstDiff {
				isFirstDiff = false
				minDiff = diff
				firstInd = i
				secondInd = j
			} else if minDiff > diff {
				minDiff = diff
				firstInd = i
				secondInd = j
			}
		}
	}

	fmt.Println(ar[firstInd], ar[secondInd])
}
