package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	// (1, 1) is black
	fmt.Print("\nThe field (x, y) is white: ")
	first_condition := !is_odd(x) && is_odd(y) // x is even, y is odd - white
	second_condition := is_odd(x) && !is_odd(y) // x is odd, y is even - white
	result := first_condition || second_condition
	fmt.Printf("%t\n", result)
}

func is_odd(x int) bool {
	return x % 2 == 1
}