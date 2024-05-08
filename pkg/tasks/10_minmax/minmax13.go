package main

import "fmt"

func main() {
	var n, a, max, maxIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	if isOdd(a) {
		max = a
		maxIndex = 1
	}
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if isOdd(a) && a > max || isOdd(a) && max == 0 {
			max = a
			maxIndex = i + 1
		}
	}
	fmt.Printf("maximum odd index: %d\n", maxIndex)
}

func isOdd(x int) bool {
	return x % 2 != 0
}