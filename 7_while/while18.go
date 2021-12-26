package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	digits_count := 0
	digits_sum := 0
	for n > 0 {
		digits_sum += n % 10
		n /= 10
		digits_count++
	}
	fmt.Printf("\nDigits count: %d\nDigits sum: %d\n", digits_count, digits_sum)
}