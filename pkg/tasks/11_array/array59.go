package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	a := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]float64, n, n)
	var sum float64
	for i := 0; i < n; i++ {
		sum = 0
		for j := 0; j <= i; j++ {
			sum += a[j]
		}
		b[i] = sum / float64(i+1)
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", b[i])
	}
	fmt.Println()
}
