package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Printf("\nf(x) = %.2f\n", f(x))
}

func f(x float64) float64 {
	if x > 0 {
		return 2 * math.Sin(x)
	} else {
		return 6 - x
	}
}