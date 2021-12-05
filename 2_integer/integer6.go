package main

import "fmt"

func main() {
	var a int16
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA(1) = %d\n", a / 10)
	fmt.Printf("A(2) = %d\n", a % 10)
}