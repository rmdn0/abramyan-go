package main

import "fmt"

func main() {
	var a int16
	fmt.Print("A = ")
	fmt.Scan(&a)
	a1 := a / 10
	a2 := a % 10
	fmt.Printf("\nСумма: %d\n", a1 + a2)
	fmt.Printf("Произведение: %d\n", a1 * a2)
}