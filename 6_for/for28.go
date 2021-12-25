package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	sqrt_1_x_value := 1.0
	power := 1.0
	multiplier := 1.0
	denomerator := 1.0
	for i := 1; i <= n; i++ {
		power *= -1 * x
		multiplier *= float64(2 * i - 3)
		denomerator *= float64(2 * i)
		sqrt_1_x_value += multiplier * power / denomerator
	}
	fmt.Printf("\narcsin(x) = %.8f\n", sqrt_1_x_value)
}