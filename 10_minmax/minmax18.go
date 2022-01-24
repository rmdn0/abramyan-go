package main

import "fmt"

func main() {
	var n, a, firstMax, lastMax, firstIndex, lastIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Scan(&a)
	firstMax = a
	firstIndex = 1
	lastMax = a
	lastIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a > firstMax {
			firstMax = a
			firstIndex = i + 1
		}
		if a >= lastMax {
			lastMax = a
			lastIndex = i + 1
		}
	}
	if firstIndex == lastIndex {
		firstIndex--
	}
	fmt.Println(lastIndex - firstIndex - 1)
}