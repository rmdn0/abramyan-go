package main

import "fmt"

func main() {
	var n, a, max, firstIndex, lastIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	max = a
	firstIndex = 1
	lastIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a > max {
			max = a
			firstIndex = i + 1
		}
		if a == max {
			lastIndex = i + 1
		}
	}
	fmt.Printf("first max index: %d\nlast max index: %d\n", firstIndex, lastIndex)
}