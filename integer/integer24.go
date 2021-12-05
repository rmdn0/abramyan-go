package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Printf("\nНомер дня недели: %d\n", k % 7)
}