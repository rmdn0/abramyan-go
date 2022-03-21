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
	for i := 0; i < n; i++ {
		if a[i] < 5 {
			b[i] = a[i] * 2
		} else {
			b[i] = a[i] / 2
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}