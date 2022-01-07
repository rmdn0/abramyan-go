package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, power float64
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		power = a
		for j := 1; j < k; j++ {
			power *= a
		}
		fmt.Printf("%.2f ", power)
	}
	fmt.Println()
}