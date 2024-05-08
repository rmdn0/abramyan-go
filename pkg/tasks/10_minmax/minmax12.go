package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, min float64
	fmt.Scan(&a)
	if a > 0 {
		min = a
	}
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a > 0 && a < min || a > 0 && min <= 0 {
			min = a
		}
	}
	fmt.Printf("minimum positive: %.2f\n", min)
}