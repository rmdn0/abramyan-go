package array

import "fmt"

func Array30() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	count := 0
	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			fmt.Printf("%d ", i+1)
			count++
		}
	}
	fmt.Printf("\n%d\n", count)
}
