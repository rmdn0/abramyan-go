package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", k)
	}
	fmt.Println()
}