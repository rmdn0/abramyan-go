package array

import "fmt"

func Array13() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	for i := n - 1; i >= 0; i -= 2 {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
