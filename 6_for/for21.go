package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	factorial := 1.0
	sum := 1.0
	for i := 1; i <=n; i++ {
		factorial *= float64(i)
		sum += 1 / factorial
	}
	fmt.Printf("\nexp(1) = %.8f\n", sum)
}