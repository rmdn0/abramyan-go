package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var k, l int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("L = ")
	fmt.Scan(&l)

	m := n - (l - k + 1)
	tempAr := make([]float64, m, m)
	j := 0
	for i := 0; i < k-1; i++ {
		tempAr[j] = ar[i]
		j++
	}
	for i := l; i < n; i++ {
		tempAr[j] = ar[i]
		j++
	}

	ar = tempAr

	fmt.Println()
	fmt.Println("New size:", m)
	for i := 0; i < m; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}
