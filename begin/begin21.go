package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 float64
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)

	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)

	fmt.Print("x3 = ")
	fmt.Scan(&x3)
	fmt.Print("y3 = ")
	fmt.Scan(&y3)

	a := distance(x1, y1, x2, y2)
	b := distance(x2, y2, x3, y3)
	c := distance(x3, y3, x1, y1)

	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("\nP = %.2f\n", p * 2)
	fmt.Printf("S = %.2f\n", s)
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
}