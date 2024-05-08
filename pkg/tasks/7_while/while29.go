package main

import (
	"fmt"
	"math"
)

func main() {
	var e float64
	fmt.Print("e = ")
	fmt.Scan(&e)
	a1 := 1.0
	a2 := 2.0
	next_a := (a1 + 2 * a2) / 3
	k := 2
	for math.Abs(a2 - a1) >= e {
		a1 = a2
		a2 = next_a
		next_a = (a1 + 2 * a2) / 3
		k++
	}
	fmt.Printf("\nK = %d\n", k)
	fmt.Printf("Ak-1 = %.6f\n", a1)
	fmt.Printf("Ak = %.6f\n", a2)
}