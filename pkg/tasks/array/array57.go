package array

import "fmt"

func Array57() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, n, n)
	j := 0
	for i := 1; i < n; i += 2 {
		b[j] = a[i]
		j++
	}
	for i := 0; i < n; i += 2 {
		b[j] = a[i]
		j++
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println()
}
