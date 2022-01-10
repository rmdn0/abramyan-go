package main

import "fmt"

func main() {
	var a, b float64
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		powerA3(a, &b)
		fmt.Printf("%.2f\n", b)
	}
}

func powerA3(a float64, b *float64) {
	*b = a * a * a
}