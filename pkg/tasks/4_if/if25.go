package main

import "fmt"

func main() {
	var x float64
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Printf("\nf(x) = %.2f\n", f(x))
}

func f(x float64) float64 {
	if x < -2 || x > 2 {
		return 2 * x
	} else {
		return -3 * x
	}
}