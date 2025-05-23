package begin

import (
	"fmt"
	"math"
)

func Begin11() {
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\n|A| + |B| = %.2f\n", math.Abs(a)+math.Abs(b))
	fmt.Printf("|A| - |B| = %.2f\n", math.Abs(a)-math.Abs(b))
	fmt.Printf("|A| * |B| = %.2f\n", math.Abs(a)*math.Abs(b))
	fmt.Printf("|A| / |B| = %.2f\n", math.Abs(a)/math.Abs(b))
}
