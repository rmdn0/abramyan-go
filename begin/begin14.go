package main

import "fmt"

func main() {
	const pi float64 = 3.14
	var l float64
	fmt.Print("L = ")
	fmt.Scan(&l)
	r := l / (2 * pi)
	s := pi * r * r
	fmt.Printf("\nR = %.2f\n", r)
	fmt.Printf("S = %.2f\n", s)
}