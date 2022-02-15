package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	fibArr := make([]int, n, n)
	fibArr[0] = 1
	fibArr[1] = 1

	for i := 2; i < n; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibArr[i])
	}
	fmt.Println()
}