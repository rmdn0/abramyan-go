package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	fmt.Printf("\nСумма: %d\n", a + b + c)
	fmt.Printf("Произведение: %d\n", a * b * c)
}