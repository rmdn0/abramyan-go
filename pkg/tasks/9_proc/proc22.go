package main

import "fmt"

func main() {
	var a, b float64
	var n int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	for i := 0; i < 3; i++ {
		fmt.Printf("N%d = ", i + 1)
		fmt.Scan(&n)
		fmt.Printf("\nCalc(%.2f, %.2f, %d) = %.2f\n", a, b, i + 1, calc(a, b, n))
	}
}

func calc(x, y float64, operation int) float64 {
	var result float64
	switch operation {
	case 1:
		result = x - y
	case 2:
		result = x * y
	case 3:
		result = x / y
	default:
		result = x + y
	}
	return result
}