package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	factorial := 1.0
	for i := 2; i <= n; i++ {
		factorial *= float64(i)
	}
	fmt.Printf("\nN! = %.2f\n", factorial)
}