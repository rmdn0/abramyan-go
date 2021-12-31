package main

import "fmt"

func main() {
	var a, product float64
	product = 1
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		product *= a
	}
	fmt.Printf("\n%.2f\n", product)
}