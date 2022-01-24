package main

import "fmt"

func main() {
	var b, c, a, maxOfBCRange float64
	var maxOfBCRangeIndex int
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if a > b && a < c && a > maxOfBCRange || a > b && a < c && maxOfBCRange == 0 {
			maxOfBCRange = a
			maxOfBCRangeIndex = i + 1
		}
	}
	fmt.Printf("max element: %.2f\nit's index: %d\n", maxOfBCRange, maxOfBCRangeIndex)
}