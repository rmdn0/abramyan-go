package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	if a < b {
		a, b = b, a
	}
	fmt.Printf("\nThe greater number: %.2f\n", a)
	fmt.Printf("The lower number: %.2f\n", b)
}