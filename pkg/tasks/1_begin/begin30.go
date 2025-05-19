package begin

import "fmt"

func Begin30() {
	const pi float32 = 3.14
	var alfa_radians float32
	fmt.Print("Alfa in radians: ")
	fmt.Scan(&alfa_radians)
	alfa_degrees := alfa_radians * 180 / pi
	fmt.Printf("\nAlfa in degrees: %.2f\n", alfa_degrees)
}
