package array

import "fmt"

func Array105() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var k, m int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("M = ")
	fmt.Scan(&m)

	tempAr := make([]float64, n+m, n+m)
	j := 0
	for i := 0; i < n; i++ {
		tempAr[j] = ar[i]
		j++
		if i == k-1 {
			for l := 0; l < m; l++ {
				tempAr[j] = 0
				j++
			}
		}
	}
	ar = tempAr
	fmt.Println()
	for i := 0; i < n+m; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
