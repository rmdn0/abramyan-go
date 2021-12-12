package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)
	fmt.Print("x3 = ")
	fmt.Scan(&x3)
	fmt.Print("y3 = ")
	fmt.Scan(&y3)
	var x4, y4 int
	
	if x1 == x2 {
		x4 = x3
	} else if x2 == x3 {
		x4 = x1
	} else if x3 == x1 {
		x4 = x2
	}

	if y1 == y2 {
		y4 = y3
	} else if y2 == y3 {
		y4 = y1
	} else if y3 == y1 {
		y4 = y2
	}

	fmt.Printf("\nx4 = %d\n", x4)
	fmt.Printf("y4 = %d\n", y4)
}