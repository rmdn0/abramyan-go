package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, ineteger_part, sum float64
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		ineteger_part, _ = math.Modf(a)
		fmt.Printf("%.2f\n", ineteger_part)
		sum += ineteger_part
	}
	fmt.Printf("\nSum: %.2f\n", sum)
}