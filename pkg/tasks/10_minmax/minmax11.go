package main

import "fmt"

func main() {
	var n, a, min, max, minIndex, maxIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	min = a
	max = a
	minIndex = 1
	maxIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a <= min {
			min = a
			minIndex = i + 1
		}
		if a >= max {
			max = a
			maxIndex = i + 1
		}
	}
	extremumIndex := minIndex
	if maxIndex > minIndex {
		extremumIndex = maxIndex
	}
	fmt.Printf("first extremum index: %d\n", extremumIndex)
}