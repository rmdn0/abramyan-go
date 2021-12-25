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
	exp_x_value := 1.0
	for i := 1; i <= n; i++ {
		factorial *= float64(i)
		power *= x
		exp_x_value += power / factorial
	}
	fmt.Printf("\nexp(x) = %.8f\n", exp_x_value)
}