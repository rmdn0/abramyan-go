package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a := n / 100 % 10
	fmt.Printf("\nСотни: %d\n", a)
}