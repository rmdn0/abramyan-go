package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA is even and two-digit: ")
	fmt.Printf("%t\n", is_even(a) && is_two_digit(a))
}

func is_even(x int) bool {
	return x % 2 == 0
}

func is_two_digit(x int) bool {
	return x >= 10 && x <= 99 || x <= -10 && x >= -99
}