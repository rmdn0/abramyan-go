package main

import "fmt"

func main() {
	var a float32
	fmt.Print("Введите сторону квадрата:\na = ")
	fmt.Scan(&a)
	fmt.Println("Площадь квадрата со стороной a:\nS =", fmt.Sprintf("%.2f", a * a))
}