package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	square_of_n := 0
	fmt.Println()
	for i := 1; i < 2 * n; i += 2 {
		square_of_n += i
		fmt.Printf("%d\t", square_of_n)
	}
	fmt.Println()
}