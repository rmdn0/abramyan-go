package begin

import "fmt"

func Begin23() {
	var a, b, c float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)

	// d := b
	// b = a
	// a = c
	// c = d

	a, b = b, a
	a, c = c, a

	fmt.Printf("\nA = %.2f\n", a)
	fmt.Printf("B = %.2f\n", b)
	fmt.Printf("C = %.2f\n", c)
}
