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

	var seriesLengths, seriesValues []int
	var count int

	for i := 0; i < n; i++ {
		seriesValues = append(seriesValues, ar[i])
		count = 1
		for i < n-1 && ar[i] == ar[i+1] {
			count++
			i++
		}
		seriesLengths = append(seriesLengths, count)
	}

	fmt.Println()
	m := len(seriesLengths)
	for i := 0; i < m; i++ {
		fmt.Print(seriesLengths[i], " ")
	}
	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Print(seriesValues[i], " ")
	}
	fmt.Println()
}
