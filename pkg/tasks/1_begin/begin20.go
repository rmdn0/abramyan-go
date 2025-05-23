package begin

import (
	"fmt"
	"math"
)

func Begin20() {
	var x1, y1, x2, y2 float64
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)

	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	fmt.Printf("\nd = %.2f\n", d)
}
