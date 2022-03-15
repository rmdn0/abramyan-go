package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	var distinctElements []int
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
		if !isNumberPresentInSlice(distinctElements, ar[i]) {
			distinctElements = append(distinctElements, ar[i])
		}
	}

	fmt.Println(len(distinctElements))
}

func isNumberPresentInSlice(ar []int, x int) bool {
	for i := 0; i < len(ar); i++ {
		if ar[i] == x {
			return true
		}
	}
	return false
}
