package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	k := 1
	for k * k <= n {
		k++
	}
	fmt.Printf("\n%d\n", k)
}