package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	fmt.Printf("\nN(modified) = %d\n", c * 100 + a * 10 + b)
}