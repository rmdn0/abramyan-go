package array

import (
	"fmt"
	"math"
)

func Array42() {
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

	var sum, nearestSum float64
	var index int
	isFirstSum := true
	for i := 0; i < n-1; i++ {
		sum = math.Abs(ar[i] + ar[i+1] - r)
		if isFirstSum {
			nearestSum = sum
			index = i
			isFirstSum = false
		} else {
			if nearestSum > sum {
				nearestSum = sum
				index = i
			}
		}
	}

	// Printing two needed adjacent elements
	fmt.Printf("%.2f", ar[index])
	if index < n-1 {
		fmt.Printf(" %.2f\n", ar[index+1])
	}
}
