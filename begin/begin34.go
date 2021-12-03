package main

import "fmt"

func main() {
	var x, y, a, b float32
	fmt.Print("X = ")
	fmt.Scan(&x)
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("Y = ")
	fmt.Scan(&y)
	fmt.Print("B = ")
	fmt.Scan(&b)
	cost_a := a / x
	cost_b := b / y
	fmt.Printf("\nЦена шоколадных конфет: %.2f\n", cost_a)
	fmt.Printf("Цена ирисок: %.2f\n", cost_b)
	fmt.Printf("\nОтношение шоколад/ирис: %.2f\n", cost_a / cost_b)
}