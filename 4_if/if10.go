package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	if a != b {
		sum := a + b
		a = sum
		b = sum
	} else {
		a = 0
		b = 0
	}
	fmt.Printf("\nA = %d\n", a)
	fmt.Printf("B = %d\n", b)
}