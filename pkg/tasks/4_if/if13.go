package main

import "fmt"

func main() {
	var a, b, c, mean_number float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	// 1st way
	if a < b && a > c || a < c && a > b {
		mean_number = a
	} else if b < a && b > c || b < c && b > a {
		mean_number = b
	} else if c < a && c > b || c < b && c > a {
		mean_number = c
	}
	fmt.Printf("\nThe mean number: %.2f\n", mean_number)
}