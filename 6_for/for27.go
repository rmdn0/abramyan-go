package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	numerator_multiplier := 1.0
	denomerator := 1.0
	power := x
	arcsin_x_value := x
	for i := 1; i <= n; i++ {
		power *= x * x
		numerator_multiplier *= float64(2 * i - 1)
		denomerator *= float64((2 * i + 1) * (2 * i))
		arcsin_x_value += (numerator_multiplier * power) / denomerator
		denomerator /= float64(2 * i + 1)
	}
	fmt.Printf("\narcsin(x) = %.8f\n", arcsin_x_value)
}