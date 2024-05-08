package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	sum := 0.0
	trigger := 1.0
	for i := 1; i <= n; i++ {
		sum += trigger * (1 + float64(i) / 10)
		trigger *= -1
	}
	fmt.Printf("\n%.2f\n", sum)
}