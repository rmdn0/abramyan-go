package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var max float64
	var maxID int
	fmt.Println()
	for i := 0; i < n-1; i++ {

		max = ar[0]
		maxID = 0
		for j := 1; j < n-i; j++ {
			if max < ar[j] {
				max = ar[j]
				maxID = j
			}
		}

		swap(&ar[maxID], &ar[n-i-1])

		for k := 0; k < n; k++ {
			fmt.Printf("%.2f ", ar[k])
		}
		fmt.Println()
	}
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}
