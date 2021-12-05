package main

import "fmt"

func main() {
	var l int
	fmt.Print("L(cm) = ")
	fmt.Scan(&l)
	fmt.Printf("\nL(m) = %d\n", l / 100)
}