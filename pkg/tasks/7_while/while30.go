package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	squares_on_a_side := 0
	squares_on_b_side := 0
	for a > c {
		a -= c
		squares_on_a_side++
	}
	for b > c {
		b -= c
		squares_on_b_side++
	}
	fmt.Printf("\n%d\n", squares_on_a_side * squares_on_b_side)
}