package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var sum, maxSum float64
	var index int
	isFirstSum := true
	for i := 0; i < n-1; i++ {
		sum = ar[i] + ar[i+1]
		if isFirstSum {
			maxSum = sum
			index = i
			isFirstSum = false
		} else {
			if maxSum < sum {
				maxSum = sum
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
