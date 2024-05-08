package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Println()
	a := 2.0
	for i := 1; i <= n; i++ {
		a = 2 + 1 / a
		fmt.Printf("%.6f\t", a)
	}
	fmt.Println()
}