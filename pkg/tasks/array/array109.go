package array

import "fmt"

func Array109() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	numberOfNegatives := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if ar[i] < 0 {
			numberOfNegatives++
		}
	}

	m := n + numberOfNegatives
	tempAr := make([]float64, m, m)
	k := 0
	for i := 0; i < n; i++ {
		tempAr[k] = ar[i]
		k++
		if ar[i] < 0 {
			tempAr[k] = 0
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
