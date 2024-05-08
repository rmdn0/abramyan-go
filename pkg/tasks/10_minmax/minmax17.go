package main

import "fmt"

func main() {
	var n, a, max, maxIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	max = a
	maxIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a >= max {
			max = a
			maxIndex = i + 1
		}
	}
	fmt.Println(n - maxIndex)
}