package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("x = ")
	fmt.Scan(&x)
	
	y := 3 * math.Pow(x, 6) - 6 * math.Pow(x, 2) - 7

	fmt.Printf("\ny = %.2f\n", y)
}