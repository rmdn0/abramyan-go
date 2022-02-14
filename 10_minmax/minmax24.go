package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var curr, prev, sum, maxSum float64
	fmt.Scan(&prev)
	for i := 1; i < n; i++ {
		fmt.Scan(&curr)
		sum = curr + prev
		if maxSum < sum || i == 1 {
			maxSum = sum
		}
		prev = curr
	}
	fmt.Printf("\n%.2f\n", maxSum)
}
