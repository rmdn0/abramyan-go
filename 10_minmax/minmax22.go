package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, min1, min2 float64
	fmt.Scan(&a)
	min1 = a
	min2 = a
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a < min2 {
			if a < min1 {
				min2 = min1
				min1 = a
			} else if a != min1 {
				min2 = a
			}
		}
		if min1 == min2 {
			min2 = a
		}
	}
	fmt.Printf("\n%.2f\t%.2f\n", min1, min2)
}
