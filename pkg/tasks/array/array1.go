package array

import "fmt"

func Array1() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i*2 + 1
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}
