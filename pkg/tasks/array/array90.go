package array

import "fmt"

func Array90() {
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

	j := 0
	tempAr := make([]float64, n-1, n-1)
	for i := 0; i < n; i++ {
		if i == k-1 {
			continue
		}
		tempAr[j] = ar[i]
		j++
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < n-1; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
