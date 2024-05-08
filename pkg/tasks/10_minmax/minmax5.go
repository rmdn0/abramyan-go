package main

import "fmt"

func main() {
	var n, maxPIndex int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var m, v, maxP float64
	fmt.Scan(&m)
	fmt.Scan(&v)
	maxP = m / v
	maxPIndex = 1
	for i := 1; i < n; i++ {
		fmt.Scan(&m)
		fmt.Scan(&v)
		if m / v > maxP {
			maxP = m / v
			maxPIndex = i + 1
		}
	}
	fmt.Printf("maxPIndex: %d\nmaxP: %.2f\n", maxPIndex, maxP)
}