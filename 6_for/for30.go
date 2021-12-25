package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	length_of_a_part := (b - a) / float64(n)
	fmt.Printf("\nH = %.5f\n", length_of_a_part)
	for i := 0; i <= n; i++ {
		power := a
		factorial := 1.0
		sin_x_value := a
		for j := 1; j <= 10; j++ {
			power *= -1 * a * a
			factorial *= float64((2 * j + 1) * (2 * j))
			sin_x_value += power / factorial
		}
		fmt.Printf("%.5f\t", 1 - sin_x_value)
		a += length_of_a_part
	}
	fmt.Println()
}