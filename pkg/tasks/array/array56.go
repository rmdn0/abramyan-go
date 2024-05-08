package array

import "fmt"

func Array56() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m := n / 3
	b := make([]int, m, m)
	for i := 0; i < m; i++ {
		b[i] = a[i*3+2]
	}

	fmt.Println("\nSize of array B:", m)
	for i := 0; i < m; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}
