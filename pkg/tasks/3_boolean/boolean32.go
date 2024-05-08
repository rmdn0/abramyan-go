package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Print("\nДанный треугольник прямоугольный: ")
	fmt.Printf("%t\n", float64(a) == hypotenuse(b, c) || float64(b) == hypotenuse(a, c) || float64(c) == hypotenuse(a, b))
}

func hypotenuse(a, b int) float64 {
	return math.Sqrt(float64(a * a + b * b))
}