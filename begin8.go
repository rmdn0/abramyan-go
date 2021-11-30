package main

import "fmt"

func main() {
	var a, b float32
	fmt.Println("Введите два числа:")
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Println("Среднее арифметическое введенных чисел:")
	fmt.Printf("%.2f\n", (a + b) / 2)
}