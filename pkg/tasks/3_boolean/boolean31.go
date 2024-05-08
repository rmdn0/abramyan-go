package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	fmt.Print("\nДанный треугольник равнобедренный: ")
	fmt.Printf("%t\n", a == b || b == c || c == a)
}