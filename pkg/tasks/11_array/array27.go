package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	breakingElement := 0
	for i := 1; i < n; i++ {
		if isPositive(ar[i]) == isPositive(ar[i-1]) {
			breakingElement = i + 1
			break
		}
	}

	fmt.Println(breakingElement)
}

func isPositive(x int) bool {
	return x > 0
}
