package main

import "fmt"

func main() {
	var m int
	fmt.Print("M(kg) = ")
	fmt.Scan(&m)
	fmt.Printf("\nM(t) = %d\n", m / 1000)
}