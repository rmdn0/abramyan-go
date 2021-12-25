package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	power := x
	factorial := 1.0
	sin_x_value := x
	for i := 1; i <= n; i++ {
		power *= -1 * x * x
		factorial *= float64((2 * i + 1) * (2 * i))
		sin_x_value += power / factorial
	}
	fmt.Printf("\nsin(x) = %.8f\n", sin_x_value)
}