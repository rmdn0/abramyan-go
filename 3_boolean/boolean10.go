package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\nExactly one of the A and B is odd: ")
	fmt.Printf("%t\n", is_odd(a) && !is_odd(b) || !is_odd(a) && is_odd(b))
}

func is_odd(x int) bool {
	return x % 2 == 1 || x % 2 == -1
}