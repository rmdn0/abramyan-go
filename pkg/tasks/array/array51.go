package array

import "fmt"

func Array51() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	a := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	for i := 0; i < n; i++ {
		a[i], b[i] = b[i], a[i]
	}

	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}
