package main

import "fmt"

func main() {
	var k, n int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var counter, a int
	const desired_number = 2
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			if a == desired_number {
				counter++
			}
		}
	}
	fmt.Printf("\n%d\n", counter)
}