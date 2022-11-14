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

	min := ar[0]
	max := ar[0]
	var minID, maxID int
	for i := 0; i < n; i++ {
		if min > ar[i] {
			min = ar[i]
			minID = i
		}
		if max < ar[i] {
			max = ar[i]
			maxID = i
		}
	}

	if minID < maxID {
		ar = addZeroBefore(ar, minID)
		ar = addZeroAfter(ar, maxID+1)
	} else {
		ar = addZeroAfter(ar, maxID)
		ar = addZeroBefore(ar, minID+1)
	}

	fmt.Println()
	for i := 0; i < n+2; i++ {
		fmt.Printf("%.2f ", ar[i])
	}
	fmt.Println()
}

func addZeroBefore(ar []float64, index int) []float64 {
	ar = append(ar[:index+1], ar[index:]...)
	ar[index] = 0
	return ar
}

func addZeroAfter(ar []float64, index int) []float64 {
	ar = append(ar[:index+1], ar[index:]...)
	ar[index+1] = 0
	return ar
}
