package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1 / float64(i)
	}
	fmt.Printf("\n%f\n", sum)
}