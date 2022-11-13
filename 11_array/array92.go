package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	m := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if isEven(ar[i]) {
			m++
		}
	}

	tempAr := make([]int, m, m)
	j := 0
	for i := 0; i < n; i++ {
		if isEven(ar[i]) {
			tempAr[j] = ar[i]
			j++
		}
	}

	ar = tempAr

	fmt.Println()
	fmt.Println("New size:", m)
	for i := 0; i < m; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}

func isEven(x int) bool {
	return x%2 == 0
}
