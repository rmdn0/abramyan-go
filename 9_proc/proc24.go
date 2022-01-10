package main

import "fmt"

func main() {
	var k, even_numbers_counter int
	for i := 0; i < 10; i++ {
		fmt.Scan(&k)
		if even(k) {
			even_numbers_counter++
		}
	}
	fmt.Printf("\nEven numbers count: %d\n", even_numbers_counter)
}

func even(k int) bool {
	return k % 2 == 0
}