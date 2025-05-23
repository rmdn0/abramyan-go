package array

import "fmt"

func Array32() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	i := 0
	for i < n-1 && ar[i] > ar[i+1] || i > 0 && ar[i] > ar[i-1] {
		i++
	}
	fmt.Println(i + 1)
}
