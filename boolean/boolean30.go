package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)
	fmt.Printf("\nДанный треугольник равносторонний: ")
	fmt.Printf("%t\n", a == b && b == c)
}