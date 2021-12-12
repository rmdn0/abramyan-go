package main

import "fmt"

func main() {
	var a, b, c, min_number float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	// 1st way
	// if a < b {
	// 	if a < c {
	// 		min_number = a
	// 	} else {
	// 		min_number = c
	// 	}
	// } else {
	// 	if b < c {
	// 		min_number = b
	// 	} else {
	// 		min_number = c
	// 	}
	// }
	// 2nd way
	if a < b && a < c {
		min_number = a
	} else if b < a && b < c {
		min_number = b
	} else if c < a && c < b {
		min_number = c
	}
	fmt.Printf("\nThe lowest number: %.2f\n", min_number)
}