package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var curr, prev, product, minProduct float64
	var minIndex, maxIndex int
	fmt.Scan(&prev)
	for i := 1; i < n; i++ {
		fmt.Scan(&curr)
		product = curr * prev
		if minProduct > product || i == 1 {
			minProduct = product
			minIndex = i
			maxIndex = i + 1
		}
		prev = curr
	}
	fmt.Printf("\n%d %d\n", minIndex, maxIndex)
}
