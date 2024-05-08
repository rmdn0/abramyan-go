package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	var result int
	if x == 0 && y == 0 {
		result = 0
	} else if y == 0 {
		result = 1
	} else if x == 0 {
		result = 2
	} else {
		result = 3
	}
	fmt.Printf("\nResult: %d\n", result)
}