package array

import "fmt"

func Array84() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	first := ar[0]
	for i := 0; i < n-1; i++ {
		ar[i] = ar[i+1]
	}
	ar[n-1] = first

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
