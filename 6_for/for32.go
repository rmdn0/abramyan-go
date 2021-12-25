package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	fmt.Println()
	a := 1.0
	for i := 1; i <= n; i++ {
		a = (a + 1) / float64(i)
		fmt.Printf("%.6f\t", a)
	}
	fmt.Println()
}