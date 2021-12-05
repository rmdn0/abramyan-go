package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	c_quantity := (a / c) * (b / c)
	free_space := a * b - c_quantity * c * c
	fmt.Printf("\nКоличество квадратов: %d\n", c_quantity)
	fmt.Printf("Площадь неазанятой части: %d\n", free_space)
}