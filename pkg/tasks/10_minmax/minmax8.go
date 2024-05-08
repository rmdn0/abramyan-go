package main

import "fmt"

func main() {
	var n, a, min, firstIndex, lastIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	min = a
	firstIndex = 1
	lastIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
			firstIndex = i + 1
		}
		if a == min {
			lastIndex = i + 1
		}
	}
	fmt.Printf("first min index: %d\nlast min index: %d\n", firstIndex, lastIndex)
}