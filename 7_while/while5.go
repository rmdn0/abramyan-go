package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	k := 0
	for n >= 2 {
		n /= 2
		k++
	}
	fmt.Printf("\n%d\n", k)
}