package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Printf("Номер дня недели: %d\n", (k + n - 2) % 7 + 1)
}