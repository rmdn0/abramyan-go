package main

import "fmt"

func main() {
	var a, b, c float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	var index int
	if b == c {
		index = 1 // a
	} else if c == a {
		index = 2 // b
	} else {
		index = 3 // c
	}
	fmt.Printf("\nThe index of different number: %d\n", index)
}