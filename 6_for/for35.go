package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	a1 := 1
	a2 := 2
	a3 := 3
	a4 := 0
	fmt.Printf("\n%d\t%d\t%d\t", a1, a2, a3)
	for i := 3; i < n; i++ {
		a4 = a3 + a2 - 2 * a1
		a1 = a2
		a2 = a3
		a3 = a4
		fmt.Printf("%d\t", a4)
	}
	fmt.Println()
}