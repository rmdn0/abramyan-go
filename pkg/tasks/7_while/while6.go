package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	double_factorial := 1.0
	for n > 1 {
		double_factorial *= float64(n)
		n -= 2
	}
	fmt.Printf("\n%.2f\n", double_factorial)
}