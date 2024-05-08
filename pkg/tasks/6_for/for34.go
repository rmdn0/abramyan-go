package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a1 := 1.0
	a2 := 2.0
	a3 := 0.0
	fmt.Printf("%.6f\t%.6f\t", a1, a2)
	for i := 3; i <= n; i++ {
		a3 = (a1 + 2 * a2) / 3
		a1 = a2
		a2 = a3
		fmt.Printf("%.6f\t", a3)
	}
	fmt.Println()
}