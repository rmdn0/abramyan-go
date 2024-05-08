package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\nx = %.2f\n", -b / a)
}