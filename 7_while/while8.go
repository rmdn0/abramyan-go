package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	k := n
	for k * k > n {
		k--
	}
	fmt.Printf("\n%d\n", k)
}