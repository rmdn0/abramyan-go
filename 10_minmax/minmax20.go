package main

import "fmt"

func main() {
	var n, a, min, max, minCount, maxCount, similarCount int
	fmt.Print("N = ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a < min || i == 0 {
			min = a
			minCount = 0
		}
		if a > max || i == 0 {
			max = a
			maxCount = 0;
		}
		if a == min {
			minCount++
		}
		if a == max {
			maxCount++
		}
		if max == min {
			similarCount++
		} else {
			similarCount = 0
		}
	}
	fmt.Println(minCount + maxCount - similarCount)
}