package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a float64
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		fmt.Printf("%d\n", int(math.Round(a)))
		sum += int(math.Round(a))
	}
	fmt.Printf("\nSum: %d\n", sum)
}