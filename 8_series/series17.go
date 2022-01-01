package main

import "fmt"

func main() {
	var b, a float64
	fmt.Print("B = ")
	fmt.Scan(&b)
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	not_printed_b := true
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if b < a && not_printed_b {
			fmt.Printf("%.2f\t", b)
			not_printed_b = false
		}
		fmt.Printf("%.2f\t", a)
	}
	if not_printed_b {
		fmt.Printf("%.2f\t", b)
	}
}