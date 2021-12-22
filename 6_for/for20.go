package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	factorial := 1.0
	sum := factorial
	for i := 2; i <= n; i++ {
		factorial *= float64(i)
		sum += factorial
	}
	fmt.Printf("\n%.2f\n", sum)
}