package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, b, p, maxP float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	maxP = 2 * (a + b)
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		p = 2 * (a + b)
		if p > maxP {
			maxP = p
		}
	}
	fmt.Printf("Max: %.2f\n", maxP)
}