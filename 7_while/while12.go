package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var sum, k int
	for sum <= n - k {
		k++
		sum += k
	}
	fmt.Printf("\nK = %d\nSum = %d\n", k, sum)
}