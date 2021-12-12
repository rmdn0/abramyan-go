package main

import "fmt"

func main() {
	var x, y float32
	fmt.Print("x = ")
	fmt.Scan(&x)
	fmt.Print("y = ")
	fmt.Scan(&y)
	var coordinate_quarter int
	if x > 0 && y > 0 {
		coordinate_quarter = 1
	} else if x < 0 && y > 0 {
		coordinate_quarter = 2
	} else if x < 0 && y < 0 {
		coordinate_quarter = 3
	} else {
		coordinate_quarter = 4
	}
	fmt.Printf("\nThe quarter of the point: %d\n", coordinate_quarter)
}