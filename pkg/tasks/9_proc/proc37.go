package main

import (
	"fmt"
	"math"
)

func main() {
	var p, a, b, c float64
	fmt.Print("P = ")
	fmt.Scan(&p)
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Printf("\n%.2f\n", power1(a, p))
	fmt.Printf("%.2f\n", power1(b, p))
	fmt.Printf("%.2f\n", power1(c, p))
}

func power1(a, b float64) float64 {
	if a <= 0 {
		return 0
	}
	return math.Exp(b * math.Log(a))
}