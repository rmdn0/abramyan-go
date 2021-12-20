package main

import (
	"fmt"
	"math"
)

func main() {
	var a, r1, r2, s float64
	var parameter int
	fmt.Print("Command: ")
	fmt.Scan(&parameter)
	switch parameter {
	case 1:
		fmt.Print("a = ")
		fmt.Scan(&a)
		r1 = a * math.Sqrt(3) / 6
		r2 = 2 * r1
		s = math.Pow(a, 2) * math.Sqrt(3) / 4
		fmt.Printf("\nR1 = %.2f\n", r1)
		fmt.Printf("R2 = %.2f\n", r2)
		fmt.Printf("S = %.2f\n", s)
	case 2:
		fmt.Print("R1 = ")
		fmt.Scan(&r1)
		a = r1 * 6 / math.Sqrt(3)
		r2 = 2 * r1
		s = math.Pow(a, 2) * math.Sqrt(3) / 4
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("R2 = %.2f\n", r2)
		fmt.Printf("S = %.2f\n", s)
	case 3:
		fmt.Print("R2 = ")
		fmt.Scan(&r2)
		r1 = r2 / 2
		a = r1 * 6 / math.Sqrt(3)
		s = math.Pow(a, 2) * math.Sqrt(3) / 4
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("R1 = %.2f\n", r1)
		fmt.Printf("S = %.2f\n", s)
	case 4:
		fmt.Print("S = ")
		fmt.Scan(&s)
		a = math.Sqrt(s * 4 / math.Sqrt(3))
		r1 = a * math.Sqrt(3) / 6
		r2 = 2 * r1
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("R1 = %.2f\n", r1)
		fmt.Printf("R2 = %.2f\n", r2)
	}
}