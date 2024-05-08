package main

import "fmt"

func main() {
	const pi float32 = 3.14
	var d float32
	fmt.Print("Введите диаметр окружности:\nd = ")
	fmt.Scan(&d)
	var l float32 = pi * d
	fmt.Println("Длина окружности с диаметром d:\nL =", fmt.Sprintf("%.2f", l))
}