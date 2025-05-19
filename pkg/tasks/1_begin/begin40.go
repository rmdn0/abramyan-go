package begin

import "fmt"

func Begin40() {
	var a1, a2, b1, b2, c1, c2 float64
	fmt.Print("A1 = ")
	fmt.Scan(&a1)
	fmt.Print("B1 = ")
	fmt.Scan(&b1)
	fmt.Print("C1 = ")
	fmt.Scan(&c1)
	fmt.Print("A2 = ")
	fmt.Scan(&a2)
	fmt.Print("B2 = ")
	fmt.Scan(&b2)
	fmt.Print("C2 = ")
	fmt.Scan(&c2)
	d := a1*b2 - a2*b1
	x := (c1*b2 - c2*b1) / d
	y := (a1*c2 - a2*c1) / d
	fmt.Printf("\nx = %.2f\n", x)
	fmt.Printf("y = %.2f\n", y)
}
