package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	is_there_odd_number := false
	for n > 0 {
		if n % 2 == 1 {
			is_there_odd_number = true
			break
		}
		n /= 10
	}
	fmt.Printf("\n%t\n", is_there_odd_number)
}