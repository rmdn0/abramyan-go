package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a float64
	fmt.Scan(&a)
	prev_a := a
	is_increasing := true
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if prev_a >= a && is_increasing {
			is_increasing = false
		}
		prev_a = a
	}
	fmt.Printf("\n%t\n", is_increasing)
}