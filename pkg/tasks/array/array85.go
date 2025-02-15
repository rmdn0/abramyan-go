package array

import "fmt"

func Array85() {
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

	tempAr := make([]float64, k, k)
	for i := 0; i < k; i++ {
		tempAr[i] = ar[n-k+i]
	}

	for i := n - 1; i > k-1; i-- {
		ar[i] = ar[i-k]
	}

	for i := 0; i < k; i++ {
		ar[i] = tempAr[i]
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
