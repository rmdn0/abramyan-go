package array

import "fmt"

func Array55() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m := n/2 + n%2
	b := make([]int, m)
	j := 0
	for i := 0; i < n; i += 2 {
		b[j] = a[i]
		j++
	}

	fmt.Println("\nSize of array B:", m)
	for i := 0; i < m; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}
