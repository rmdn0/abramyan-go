package array

import "fmt"

func Array89() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := 0; i < n-1; i++ {
		if ar[i] < ar[i+1] {
			swap(&ar[i], &ar[i+1])
		}
	}

	for i := n - 1; i > 0; i-- {
		if ar[i] > ar[i-1] {
			swap(&ar[i], &ar[i-1])
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
