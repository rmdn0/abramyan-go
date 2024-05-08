package main

import "fmt"

func main() {
	var b, a, minOfGreatThenB float64
	var minOfGreatThenBIndex int
	fmt.Print("B = ")
	fmt.Scan(&b)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if a > b && a < minOfGreatThenB || a > b && minOfGreatThenB == 0 {
			minOfGreatThenB = a
			minOfGreatThenBIndex = i + 1
		}
	}
	fmt.Printf("min element: %.2f\nit's index: %d\n", minOfGreatThenB, minOfGreatThenBIndex)
}