package main

import "fmt"

func main() {
	var k int
	fmt.Print("K = ")
	fmt.Scan(&k)
	var a, numbers_in_set, total_numbers int
	for i := 0; i < k; i++ {
		fmt.Scan(&a)
		numbers_in_set = 0
		for a != 0 {
			numbers_in_set++
			fmt.Scan(&a)
		}
		fmt.Printf("%d\n", numbers_in_set)
		total_numbers += numbers_in_set
	}
	fmt.Printf("Total numbers: %d\n", total_numbers)
}