package array

import "fmt"

func Array112() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	fmt.Println()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ar[j] > ar[j+1] {
				swap(&ar[j], &ar[j+1])
			}
		}
		for k := 0; k < n; k++ {
			fmt.Printf("%.2f ", ar[k])
		}
		fmt.Println()
	}
}
