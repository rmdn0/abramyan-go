package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	summ := 0
	for i := 0; i <= n; i++ {
		summ += int(math.Pow(float64(n + i), 2))
	}
	fmt.Printf("\n%d\n", summ)
}