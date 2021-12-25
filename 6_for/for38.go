package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	sum := 0.0
	power := 1.0
	for i := 1; i <= n; i++ {
		power = 1.0
		for j := n - i + 1; j >= 1; j-- {
			power *= float64(i)
		}
		sum += power
	}
	fmt.Printf("\n%.2f\n", sum)
}