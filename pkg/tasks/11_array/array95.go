package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	m := 1
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if i != 0 && ar[i] != ar[i-1] {
			m++
		}
	}

	tempAr := make([]int, m, m)
	tempAr[0] = ar[0]
	j := 1
	for i := 1; i < n; i++ {
		if ar[i] != ar[i-1] {
			tempAr[j] = ar[i]
			j++
		}
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
