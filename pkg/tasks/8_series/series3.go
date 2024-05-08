package main

import "fmt"

func main() {
	var a, sum float64
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		sum += a
	}
	fmt.Printf("\n%.2f\n", sum / 10)
}