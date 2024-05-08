package main

import "fmt"

func main() {
	var a, b, c float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	if a < b && b < c {
		a *= 2
		b *= 2
		c *= 2
	} else {
		a *= -1
		b *= -1
		c *= -1
	}
	fmt.Printf("\nA = %.2f\n", a)
	fmt.Printf("B = %.2f\n", b)
	fmt.Printf("C = %.2f\n", c)
}