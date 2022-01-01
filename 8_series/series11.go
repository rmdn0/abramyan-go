package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a int
	var less_then_k bool
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a < k {
			less_then_k = true
		}
	}
	fmt.Printf("\n%t\n", less_then_k)
}