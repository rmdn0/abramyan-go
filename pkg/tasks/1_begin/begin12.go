package begin

import "fmt"
import "math"

func Begin12() {
	var a, b float64
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	c := math.Sqrt(a*a + b*b)
	p := a + b + c
	fmt.Printf("\nc = %.2f\n", c)
	fmt.Printf("p = %.2f\n", p)
}
