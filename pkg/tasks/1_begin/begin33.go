package main

import "fmt"

func main() {
	var x, y, a float32
	fmt.Print("X = ")
	fmt.Scan(&x)
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("Y = ")
	fmt.Scan(&y)
	cost := a / x
	fmt.Printf("\nЦена за 1 кг.: %.2f\n", cost)
	fmt.Printf("Цена за Y кг.: %.2f\n", cost * y)
}