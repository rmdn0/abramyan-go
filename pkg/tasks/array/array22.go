package array

import "fmt"

func Array22() {
	var n, k, l int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("L = ")
	fmt.Scan(&l)

	var sum float64

	for i := 0; i < k-1; i++ {
		sum += ar[i]
	}

	for i := l; i < n; i++ {
		sum += ar[i]
	}

	fmt.Printf("\n%.2f\n", sum)
}
