package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	parts_quantity := 0
	for a > b {
		a -= b
		parts_quantity++
	}
	fmt.Printf("\n%d\n", parts_quantity)
}