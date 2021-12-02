package main

import "fmt"

func main() {
	const pi float32 = 3.14
	var alfa_degrees float32
	fmt.Print("Alfa in degrees: ")
	fmt.Scan(&alfa_degrees)
	alfa_radians := alfa_degrees * pi / 180
	fmt.Printf("\nAlfa in radians: %.2f\n", alfa_radians)
}