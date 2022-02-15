package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var a, d float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("D = ")
	fmt.Scan(&d)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		ar[i] = a + d*float64(i)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
