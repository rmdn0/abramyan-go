package main

import "fmt"

func main() {
	var k, square_numbers_counter int
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if isSquare(k) {
			square_numbers_counter++
		}
	}
	fmt.Printf("\nSquare numbers count: %d\n", square_numbers_counter)
}

func isSquare(k int) bool {
	number := 1
	for number * number < k {
		number++
	}
	return number * number == k
}
