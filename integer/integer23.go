package main

import "fmt"

func main() {
	var n int
	fmt.Print("N(sec) = ")
	fmt.Scan(&n)
	fmt.Printf("\nN(mins from last hr) = %d\n", n % 3600 / 60)
}