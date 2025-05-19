package begin

import "fmt"
import "math"

func Begin17() {
	var a, b, c float64
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)

	var ac = math.Abs(c - a)
	var bc = math.Abs(c - b)

	fmt.Printf("\nAC = %.2f\n", ac)
	fmt.Printf("BC = %.2f\n", bc)
	fmt.Printf("AC + BC = %.2f\n", ac+bc)
}
