package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	var diff_ab, diff_ac float64
	diff_ab = math.Abs(a - b)
	diff_ac = math.Abs(a - c)
	var nearest_point, distance float64
	if diff_ab < diff_ac {
		nearest_point = b
		distance = diff_ab
	} else {
		nearest_point = c
		distance = diff_ac
	}
	fmt.Printf("\nThe nearest point: %.2f\n", nearest_point)
	fmt.Printf("The distance: %.2f\n", distance)
}