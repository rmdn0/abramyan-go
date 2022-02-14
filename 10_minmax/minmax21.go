package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, min, max, sum float64
	fmt.Scan(&a)
	min = a
	max = a
	sum = a
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		sum += a
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	fmt.Printf("\n%.2f\n", (sum - min - max) / float64(n - 2))
}