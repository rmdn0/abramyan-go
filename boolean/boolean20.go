package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	x := a / 100
	y := a % 100 / 10
	z := a % 10
	fmt.Printf("\nAll of the digits are distinct: ")
	fmt.Printf("%t\n", x != y && y != z && z != x)
}