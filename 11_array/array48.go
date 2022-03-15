package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	var distinctElements []int
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if !isNumberPresentInSlice(distinctElements, ar[i]) {
			distinctElements = append(distinctElements, ar[i])
		}
	}

	var amountOfDuplicates, maxAmountOfDuplicates int
	for i := 0; i < len(distinctElements); i++ {
		amountOfDuplicates = 0
		for j := 0; j < n; j++ {
			if distinctElements[i] == ar[j] {
				amountOfDuplicates++
			}
		}
		if maxAmountOfDuplicates < amountOfDuplicates {
			maxAmountOfDuplicates = amountOfDuplicates
		}
	}

	fmt.Println(maxAmountOfDuplicates)
}

func isNumberPresentInSlice(ar []int, x int) bool {
	for i := 0; i < len(ar); i++ {
		if ar[i] == x {
			return true
		}
	}
	return false
}
