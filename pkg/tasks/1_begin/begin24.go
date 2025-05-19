package begin

import "fmt"

func Begin24() {
	var a, b, c float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)

	// d := c
	// c = a
	// a = b
	// b = d

	a, c = c, a
	a, b = b, a

	fmt.Printf("\nA = %.2f\n", a)
	fmt.Printf("B = %.2f\n", b)
	fmt.Printf("C = %.2f\n", c)
}
