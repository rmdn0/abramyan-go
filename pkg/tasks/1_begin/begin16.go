package begin

import (
	"fmt"
	"math"
)

func Begin16() {
	var x1, x2 float64
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Printf("\n|x2 - x1| = %.2f\n", math.Abs(x2-x1))
}
