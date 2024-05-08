package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\nA and B are both odd: ")
	fmt.Printf("%t\n", is_odd(a) && is_odd(b))
}

func is_odd(x int) bool {
	return x % 2 == 1 || x % 2 == -1
}