package array

import "fmt"

func Array107() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	m := n*2 + n%2
	tempAr := make([]float64, m, m)
	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= (i+1)%2*2; j++ {
			tempAr[k] = ar[i]
			k++
		}
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
