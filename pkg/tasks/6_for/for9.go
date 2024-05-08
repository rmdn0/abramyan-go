package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	sum_of_squares := 0
	for i := a; i <= b; i++ {
		sum_of_squares += i * i
	}
	fmt.Printf("\n%d\n", sum_of_squares)
}