package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	fmt.Printf("\nN(modified) = %d\n", a * 100 + c * 10 + b)
}