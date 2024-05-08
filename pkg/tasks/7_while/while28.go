package main

import (
	"fmt"
	"math"
)

func main() {
	var e float64
	fmt.Print("e = ")
	fmt.Scan(&e)
	prev_a := 2.0
	next_a := 2 + 1 / prev_a
	k := 2
	for math.Abs(next_a - prev_a) >= e {
		prev_a = next_a
		next_a = 2 + 1 / prev_a
		k++
	}
	fmt.Printf("\nK = %d\n", k)
	fmt.Printf("Ak-1 = %.6f\n", prev_a)
	fmt.Printf("Ak = %.6f\n", next_a)
}