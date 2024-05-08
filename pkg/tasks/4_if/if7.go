package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	index_of_min_number := 1
	if b < a {
		index_of_min_number = 2
	}
	fmt.Printf("\nThe index of min number: %d\n", index_of_min_number)
}