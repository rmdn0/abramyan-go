package array

import "fmt"

func Array7() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
