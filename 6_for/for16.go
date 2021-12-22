package main

import "fmt"

func main() {
	var a float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	power := 1.0
	fmt.Println()
	for i := 0; i < n; i++ {
		power *= a
		fmt.Printf("%.2f\t", power)
	}
	fmt.Println()
}