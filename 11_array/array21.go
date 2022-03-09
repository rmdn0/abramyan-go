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

	var k, l int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("L = ")
	fmt.Scan(&l)

	var sum float64
	for i := k - 1; i < l; i++ {
		sum += ar[i]
	}

	fmt.Printf("\n%.2f\n", sum / float64(l - k + 1))
}
