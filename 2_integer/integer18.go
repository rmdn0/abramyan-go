package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 1000 % 10
	fmt.Printf("\nТысячи: %d\n", a)
}