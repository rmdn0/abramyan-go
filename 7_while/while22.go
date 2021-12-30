package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	divider := 2
	is_prime := true
	for divider < n {
		if n % divider == 0 {
			is_prime = false
			break
		}
		divider++
	}
	fmt.Printf("\n%t\n", is_prime)
}