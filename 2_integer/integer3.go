package main

import "fmt"

func main() {
	var size int
	fmt.Print("S(b) = ")
	fmt.Scan(&size)
	fmt.Printf("\nS(Kb) = %d\n", size / 1024)
}