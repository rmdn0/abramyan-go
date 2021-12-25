package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	power := 1.0
	factorial := 1.0
	cos_x_value := 0.0
	trigger := 1.0
	for i := 0; i <= 2 * n; i++ {
		if i % 2 == 0 {
			cos_x_value += trigger * power / factorial
			trigger *= -1.0
		}
		power *= x
		factorial *= float64(i + 1)
	}
	fmt.Printf("\ncos(x) = %.8f\n", cos_x_value)
}