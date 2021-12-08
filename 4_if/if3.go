package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scan(&a)
	if a > 0 {
		a -= 8
	} else if a < 0 {
		a += 6
	} else {
		a = 10
	}
	fmt.Printf("\nA = %d\n", a)
}