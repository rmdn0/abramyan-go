package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)

	ar := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	var tempAr []int
	var seriesNumber, currentSeriesLength int

	for i := 0; i < n; i++ {
		seriesNumber++
		currentSeriesLength = 0
		for i+currentSeriesLength < n-1 && ar[i+currentSeriesLength] == ar[i+currentSeriesLength+1] {
			currentSeriesLength++
		}

		if seriesNumber != k {
			tempAr = append(tempAr, ar[i:i+currentSeriesLength+1]...)
		}

		i += currentSeriesLength
	}

	ar = tempAr
	n = len(ar)

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
