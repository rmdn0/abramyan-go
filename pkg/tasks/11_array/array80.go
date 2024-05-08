package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := 0; i < n-1; i++ {
		ar[i] = ar[i+1]
	}
	ar[n-1] = 0

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
