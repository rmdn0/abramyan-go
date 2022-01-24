package main

import "fmt"

func main() {
	var n, a, min, minIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
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
	fmt.Println(minIndex - 1)
}