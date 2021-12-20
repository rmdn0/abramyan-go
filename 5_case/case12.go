package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.14
	var r, d, l, s float64
	var parameter int
	fmt.Print("Command: ")
	fmt.Scan(&parameter)
	switch parameter {
	case 1:
		fmt.Print("R = ")
		fmt.Scan(&r)
		d = r * 2
		l = 2 * pi * r
		s = pi * math.Pow(r, 2)
		fmt.Printf("\nD = %.2f\n", d)
		fmt.Printf("L = %.2f\n", l)
		fmt.Printf("S = %.2f\n", s)
	case 2:
		fmt.Print("D = ")
		fmt.Scan(&d)
		r = d / 2
		l = 2 * pi * r
		s = pi * math.Pow(r, 2)
		fmt.Printf("\nR = %.2f\n", r)
		fmt.Printf("L = %.2f\n", l)
		fmt.Printf("S = %.2f\n", s)
	case 3:
		fmt.Print("L = ")
		fmt.Scan(&l)
		r = l / 2 / pi
		d = r * 2
		s = pi * math.Pow(r, 2)
		fmt.Printf("\nR = %.2f\n", r)
		fmt.Printf("D = %.2f\n", d)
		fmt.Printf("S = %.2f\n", s)
	case 4:
		fmt.Print("S = ")
		fmt.Scan(&s)
		r = math.Sqrt(s / pi)
		d = r * 2
		l = 2 * pi * r
		fmt.Printf("\nR = %.2f\n", r)
		fmt.Printf("D = %.2f\n", d)
		fmt.Printf("L = %.2f\n", l)
	}
}