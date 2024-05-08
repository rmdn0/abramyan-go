package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, b, minS float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	minS = a * b
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		if a * b < minS {
			minS = a * b
		}
	}
	fmt.Printf("Min: %.2f\n", minS)
}