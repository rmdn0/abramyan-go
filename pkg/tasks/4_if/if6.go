package main

import "fmt"

func main() {
	var a, b, max float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Printf("\nMax: %.2f\n", max)
}