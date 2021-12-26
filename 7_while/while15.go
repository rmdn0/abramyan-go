package main

import "fmt"

func main() {
	var percent float64
	fmt.Print("P = ")
	fmt.Scan(&percent)
	deposit := 1000.0
	var duration int
	for deposit <= 1100 {
		duration++
		deposit += deposit * percent / 100
	}
	fmt.Printf("\nK = %d\nS = %.2f\n", duration, deposit)
}