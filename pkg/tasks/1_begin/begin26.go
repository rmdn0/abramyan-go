package begin

import (
	"fmt"
	"math"
)

func Begin26() {
	var x float64
	fmt.Print("x = ")
	fmt.Scan(&x)

	y := 4*math.Pow(x-3, 6) - 7*math.Pow(x-3, 3) + 2

	fmt.Printf("\ny = %.2f\n", y)
}
