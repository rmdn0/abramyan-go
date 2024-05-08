package array

import "fmt"

func Array53() {
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

	c := make([]float64, n, n)
	for i := 0; i < n; i++ {
		c[i] = maxOfTwo(a[i], b[i])
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(c[i], " ")
	}
	fmt.Println()
}

func maxOfTwo(x, y float64) float64 {
	if x > y {
		return x
	} else {
		return y
	}
}
