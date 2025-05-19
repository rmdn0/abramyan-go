package begin

import "fmt"
import "math"

func Begin15() {
	const pi float64 = 3.14
	var s float64
	fmt.Print("S = ")
	fmt.Scan(&s)
	d := math.Sqrt(4 * s / pi)
	l := pi * d
	fmt.Printf("\nD = %.2f\n", d)
	fmt.Printf("L = %.2f\n", l)
}
