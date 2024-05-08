package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	product := 1.0
	for i := 1; i <= n; i++ {
		product *= 1 + float64(i) / 10
	}
	fmt.Printf("\n%f\n", product)
}