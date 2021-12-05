package main

import "fmt"

func main() {
	var x, y float32
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	fmt.Printf("\nThe point (x, y) lies in the second or third coordinatic quarters: ")
	fmt.Printf("%t\n", x < 0 && y > 0 || x < 0 && y < 0)
}