package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	k := 1
	power := 3
	for power <= n {
		power *= 3
		k++
	}
	fmt.Printf("\n%d\n", k)
}