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

	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)

	for i := n - 1; i > k-1; i-- {
		ar[i] = ar[i-k]
	}
	for i := k - 1; i >= 0; i-- {
		ar[i] = 0
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
