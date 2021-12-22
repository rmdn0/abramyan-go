package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	product := 1
	for i := a; i <= b; i++ {
		product *= i
	}
	fmt.Printf("\n%d\n", product)
}