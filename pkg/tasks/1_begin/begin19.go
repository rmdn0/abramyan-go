package begin

import (
	"fmt"
	"math"
)

func Begin19() {
	var x1, y1, x2, y2 float64
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)

	a := math.Abs(y2 - y1)
	b := math.Abs(x2 - x1)
	p := 2 * (a + b)
	s := a * b
	fmt.Printf("\nP = %.2f\n", p)
	fmt.Printf("S = %.2f\n", s)
}
