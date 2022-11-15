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

	fmt.Println()
	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			swap(&ar[i], &ar[i+1])
		}
		for j := i; j > 0; j-- {
			if ar[j] < ar[j-1] {
				swap(&ar[j], &ar[j-1])
			}
		}

		for k := 0; k < n; k++ {
			fmt.Printf("%.2f ", ar[k])
		}
		fmt.Println()
	}
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}
