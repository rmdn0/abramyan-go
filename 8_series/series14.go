package main

import "fmt"

func main() {
	var k, a, count int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Scan(&a)
	for a != 0 {
		if a < k {
			count++
		}
		fmt.Scan(&a)
	}
	fmt.Printf("\n%d\n", count)
}