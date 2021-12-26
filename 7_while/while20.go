package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var digit int
	var has_2 bool
	for n > 0 {
		digit = n % 10
		if digit == 2 {
			has_2 = true
			break
		}
		n /= 10
	}
	fmt.Printf("\n%t\n", has_2)
}