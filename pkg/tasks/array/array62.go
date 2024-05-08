package array

import "fmt"

func Array62() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	var positive, negative int
	a := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] > 0 {
			positive++
		} else if a[i] < 0 {
			negative++
		}
	}

	b := make([]float64, positive, positive)
	c := make([]float64, negative, negative)
	var j, k int
	for i := 0; i < n; i++ {
		if a[i] > 0 {
			b[j] = a[i]
			j++
		} else if a[i] < 0 {
			c[k] = a[i]
			k++
		}
	}

	fmt.Println("\nSize of array B:", positive)
	for i := 0; i < positive; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println("\nSize of array C:", negative)
	for i := 0; i < negative; i++ {
		fmt.Print(c[i], " ")
	}
	fmt.Println()
}
