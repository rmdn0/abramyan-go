package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Printf("\nAll of them are positive: ")
	fmt.Printf("%t\n", is_positive(a) && is_positive(b) && is_positive(c))
}

func is_positive(x int) bool {
	return x > 0
}