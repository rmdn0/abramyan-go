package main

import "fmt"

func main() {
	var n, minIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, min float64
	fmt.Scan(&a)
	min = a
	minIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
			minIndex = i + 1
		}
	}
	fmt.Printf("Min index: %d\n", minIndex)
}