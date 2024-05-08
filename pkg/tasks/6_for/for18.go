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
	sum := power
	for i := 0; i < n; i++ {
		power *= -a
		sum += power
	}
	fmt.Printf("\n%.2f\n", sum)
}