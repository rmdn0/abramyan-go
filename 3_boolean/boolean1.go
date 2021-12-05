package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Printf("\nA > 0: %t\n", a > 0)
}