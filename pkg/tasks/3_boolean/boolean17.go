package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA is odd and three-digit: ")
	fmt.Printf("%t\n", !is_even(a) && is_three_digit(a))
}

func is_even(x int) bool {
	return x % 2 == 0
}

func is_three_digit(x int) bool {
	return x >= 100 && x <= 999 || x <= -100 && x >= -999
}