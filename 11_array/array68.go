package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	minIndex := 0
	maxIndex := 0

	for i := 1; i < n; i++ {
		if ar[i] > ar[maxIndex] {
			maxIndex = i
		}
		if ar[i] < ar[minIndex] {
			minIndex = i
		}
	}

	ar[minIndex], ar[maxIndex] = ar[maxIndex], ar[minIndex]

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}