package main

import "fmt"

func main() {
	var n, a, min, minCount int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	min = a
	minCount = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a == min {
			minCount++
		}
		if a < min {
			min = a
			minCount = 1
		}
	}
	fmt.Println(minCount)
}