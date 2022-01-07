package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, power float64
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		power = a
		for j := 0; j < i; j++ {
			power *= a
		}
		fmt.Printf("%.2f ", power)
	}
	fmt.Println()
}