package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, min, max float64
	fmt.Scan(&a)
	min = a
	max = a
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	fmt.Printf("Min: %.2f\nMax: %.2f\n", min, max)
}