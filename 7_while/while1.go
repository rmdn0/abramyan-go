package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	for a > b {
		a -= b
	}
	fmt.Printf("\n%.2f\n", a)
}