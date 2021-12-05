package main

import "fmt"

func main() {
	var x, y, x1, y1, x2, y2 float32
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)
	fmt.Printf("\nThe point (x, y) lies in the rectangle (x1, y1, x2, y2): ")
	fmt.Printf("%t\n", x > x1 && y < y1 && x < x2 && y > y2)
}