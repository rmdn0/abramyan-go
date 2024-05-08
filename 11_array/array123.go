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
	var firstSeriesLength int

	for i := 0; i < n; i++ {
		seriesNumber++
		currentSeriesLength = 0
		for i+currentSeriesLength < n-1 && ar[i+currentSeriesLength] == ar[i+currentSeriesLength+1] {
			currentSeriesLength++
		}

		if i == 0 {
			firstSeriesLength = currentSeriesLength
		}

		fmt.Println("seq", ar[i:i+currentSeriesLength+1], "temp", tempAr, "ar", ar)
		if seriesNumber == k {
			changed := append(ar[i:i+currentSeriesLength+1], tempAr[firstSeriesLength+1:]...)
			fmt.Println(changed)
			changed = append(changed, tempAr[:firstSeriesLength+1]...)
			fmt.Println(changed)
			tempAr = changed
		} else {
			tempAr = append(tempAr, ar[i:i+currentSeriesLength+1]...)
		}

		i += currentSeriesLength
	}
	fmt.Println("ar", ar)

	ar = tempAr
	n = len(ar)

	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Print(ar[i], " ")
	}
	fmt.Println()
}
