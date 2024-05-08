package main

import "fmt"

func main() {
	var x float64
	fmt.Print("X = ")
	fmt.Scan(&x)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	arctg_x_value := 0.0
	power := x
	trigger := 1.0
	for i := 1; i <= 2 * n + 1; i++ {
		if i % 2 == 1 {
			arctg_x_value += trigger * power / float64(i)
			trigger *= -1.0
		}
		power *= x
	}
	fmt.Printf("\narct(x) = %.8f\n", arctg_x_value)
}