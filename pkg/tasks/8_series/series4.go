package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, sum float64
	product := 1.0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
		product *= a
	}
	fmt.Printf("\nSum: %.2f\tProduct: %.2f\n", sum, product)
}