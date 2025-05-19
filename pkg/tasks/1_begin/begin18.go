package begin

import (
	"fmt"
	"math"
)

func Begin18() {
	var a, b, c float64
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)

	ac := math.Abs(c - a)
	bc := math.Abs(c - b)

	fmt.Printf("\nAC * BC = %.2f\n", ac*bc)
}
