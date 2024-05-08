package main

import "fmt"

func main() {
	var a float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	var sum float64
	var k int
	for sum <= a {
		k++
		sum += 1 / float64(k)
	}
	fmt.Printf("\nK = %d\nSum = %.5f\n", k, sum)
}