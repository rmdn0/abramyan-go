package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	tempAr := make([]int, 0, n)
	uniqueElements := make(map[int]struct{})
	for i := n - 1; i >= 0; i-- {
		if _, ok := uniqueElements[ar[i]]; !ok {
			uniqueElements[ar[i]] = struct{}{}
			tempAr = append([]int{ar[i]}, tempAr...)
		}
	}

	ar = tempAr

	fmt.Println()
	for i := 0; i < len(ar); i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
