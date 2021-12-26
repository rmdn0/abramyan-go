package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	reversed_n := 0
	for n > 0 {
		reversed_n *= 10
		reversed_n += n % 10
		n /= 10
	}
	fmt.Printf("\nReversed N: %d\n", reversed_n)
}