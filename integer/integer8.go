package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	n1 := n % 10 * 10 + n / 10
	fmt.Printf("\nN(inverted) = %d\n", n1)
}