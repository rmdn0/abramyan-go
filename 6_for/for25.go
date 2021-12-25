package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ln_1_x_value := 0.0
	power := x
	trigger := 1.0
	for i := 1; i <= n; i++ {
		ln_1_x_value += trigger * power / float64(i)
		trigger *= -1.0
		power *= x
	}
	fmt.Printf("\nln(1 + x) = %.8f\n", ln_1_x_value)
}