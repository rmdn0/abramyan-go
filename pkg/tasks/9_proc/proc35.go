package main

import "fmt"

func main() {
	var n int
	for i := 0; i < 5; i++ {
		fmt.Print("N = ")
		fmt.Scan(&n)
		fmt.Printf("%e\n\n", fact2(n))
	}
}

func fact2(n int) float64 {
	result := 1.0
	for n > 1 {
		result *= float64(n)
		n -= 2
	}
	return result
}