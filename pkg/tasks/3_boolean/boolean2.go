package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA is odd: %t\n", a % 2 == 1 || a % 2 == -1)
}