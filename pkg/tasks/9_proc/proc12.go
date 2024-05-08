package main

import "fmt"

func main() {
	var a, b, c float64
	for i := 0; i < 2; i++ {
		fmt.Print("A = ")
		fmt.Scan(&a)
		fmt.Print("B = ")
		fmt.Scan(&b)
		fmt.Print("C = ")
		fmt.Scan(&c)
		sortInc3(&a, &b, &c)
		fmt.Printf("\nA = %.2f\n", a)
		fmt.Printf("B = %.2f\n", b)
		fmt.Printf("C = %.2f\n", c)
	}
}

func sortInc3(a, b, c *float64) {
	minmax(a, b)
	minmax(b, c)
	minmax(a, b)
}

func minmax(x, y *float64) {
	if *x > *y {
		*x, *y = *y, *x
	}
}