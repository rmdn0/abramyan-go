package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Printf("Номер дня недели: %d\n", (k + 4) % 7 + 1)
}