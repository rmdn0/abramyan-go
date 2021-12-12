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
	var min, max float32
	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}
	if min > c {
		min = c
	}
	if max < c {
		max = c
	}
	var mean float32
	if max > a && a > min {
		mean = a
	} else if max > b && b > min {
		mean = b
	} else {
		mean = c
	}
	fmt.Printf("\n%.2f\t%.2f\t%.2f\t", min, mean, max)
}