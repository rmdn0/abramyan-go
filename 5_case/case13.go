package main

import (
	"fmt"
	"math"
)

func main() {
	var a, c, h, s float64
	var parameter int
	fmt.Print("Command: ")
	fmt.Scan(&parameter)
	switch parameter {
	case 1:
		fmt.Print("a = ")
		fmt.Scan(&a)
		c = a * math.Sqrt(2)
		h = c / 2
		s = c * h / 2
		fmt.Printf("\nc = %.2f\n", c)
		fmt.Printf("h = %.2f\n", h)
		fmt.Printf("S = %.2f\n", s)
	case 2:
		fmt.Print("c = ")
		fmt.Scan(&c)
		a = c / math.Sqrt(2)
		h = c / 2
		s = c * h / 2
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("h = %.2f\n", h)
		fmt.Printf("S = %.2f\n", s)
	case 3:
		fmt.Print("h = ")
		fmt.Scan(&h)
		c = h * 2
		a = c / math.Sqrt(2)
		s = c * h / 2
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("c = %.2f\n", c)
		fmt.Printf("S = %.2f\n", s)
	case 4:
		fmt.Print("S = ")
		fmt.Scan(&s)
		c = math.Sqrt(4 * s)
		a = c / math.Sqrt(2)
		h = c / 2
		fmt.Printf("\na = %.2f\n", a)
		fmt.Printf("c = %.2f\n", c)
		fmt.Printf("h = %.2f\n", h)
	}
}