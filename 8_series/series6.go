package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, fractional_part float64
	product := 1.0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		_, fractional_part = math.Modf(a)
		fmt.Printf("%.2f\n", fractional_part)
		product *= fractional_part
	}
	fmt.Printf("\nProduct: %.6f\n", product)
}