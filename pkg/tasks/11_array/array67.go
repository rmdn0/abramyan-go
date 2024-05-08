package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	odd := 0
	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if !isEven(ar[i]) {
			odd = ar[i]
		}
	}

	if odd != 0 {
		for i := 0; i < n; i++ {
			if !isEven(ar[i]) {
				ar[i] += odd
			}
		}
	}

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}

func isEven(x int) bool {
	return x%2 == 0
}
