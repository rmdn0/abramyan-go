package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Printf("\nN(inverted) = %d\n", n % 10 * 100 + n % 100 / 10 * 10 + n / 100)
}