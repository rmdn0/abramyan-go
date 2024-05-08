package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	positive := false
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > 0 {
			positive = true
		}
	}
	fmt.Printf("\n%t\n", positive)
}