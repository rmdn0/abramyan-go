package main

import "fmt"

func main() {
	var a float32
	fmt.Print("Введите сторону квадрата:\na = ")
	fmt.Scan(&a)
	fmt.Println("Периметр квадрата со стороной a:\nP =", fmt.Sprintf("%.2f", 4 * a))
}
