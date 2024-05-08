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
	fmt.Print("\nThere exists a triangle with sides a, b, c: ")
	fmt.Printf("%t\n", a + b > c && a + c > b && b + c > a)
}