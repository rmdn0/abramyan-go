package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, max1, max2, max3 float64
	fmt.Scan(&a)
	max1 = a
	max2 = a
	max3 = a
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if a > max3 || i < 3 {
			if a > max2 || i < 3 {
				if a > max1 {
					max3 = max2
					max2 = max1
					max1 = a
				} else {
					max3 = max2
					max2 = a
				}
			} else {
				max3 = a
			}
		}
	}
	fmt.Printf("\n%.2f\t%.2f\t%.2f\n", max1, max2, max3)
}
