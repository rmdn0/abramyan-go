package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Print("K = ")
	fmt.Scan(&k)
	quotient := 0
	for n >= k {
		n -= k
		quotient++
	}
	fmt.Printf("\nQuotient: %d\nRemainder: %d\n", quotient, n)
}