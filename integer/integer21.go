package main

import "fmt"

func main() {
	var n int
	fmt.Print("N(sec) = ")
	fmt.Scan(&n)
	fmt.Printf("\nN(sec from last min) = %d\n", n % 60)
}